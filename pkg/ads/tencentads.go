package ads

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	uuid "github.com/google/uuid"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
)

// SDKClient ...
type SDKClient struct {
	http.RoundTripper
	Config         *config.SDKConfig
	Client         *api.APIClient
	Ctx            *context.Context
	Version        string
	ApiVersion     string
	middlewareList []Middleware
}

// Init ...
func Init(cfg *config.SDKConfig) *SDKClient {
	version := "1.7.66"
	apiVersion := "v1.1"
	ctx := context.Background()
	nonce := uuid.New().String()
	apiKey := config.APIKey{
		AccessToken: cfg.AccessToken,
		Timestamp:   strconv.FormatInt(time.Now().Unix(), 10),
		Nonce:       nonce[0:8] + nonce[9:13] + nonce[14:18],
	}
	ctx = context.WithValue(ctx, config.ContextAPIKey, apiKey)
	client := api.NewAPIClient(cfg)
	sdkClient := &SDKClient{
		Config:       cfg,
		Ctx:          &ctx,
		Client:       client,
		RoundTripper: http.DefaultTransport,
		Version:      version,
		ApiVersion:   apiVersion,
	}
	sdkClient.Client.Cfg.HTTPClient.Transport = sdkClient
	sdkClient.UseSandbox()
	sdkClient.middlewareList = []Middleware{
		&AuthMiddleware{sdkClient},
		&HeaderMiddleware{sdkClient},
		&DiffHostMiddleware{sdkClient},
		&LogMiddleware{sdkClient},
	}
	return sdkClient
}

// SetIpPort set ip port in Host
func (tads *SDKClient) SetIpPort(ip string, port int, schema string) {
	ipPort := fmt.Sprintf("%s:%d", ip, port)
	tads.SetHost(ipPort, schema)
}

// SetHost ...
func (tads *SDKClient) SetHost(host string, schema string) *SDKClient {
	modified := false
	if host != "" {
		tads.Client.Cfg.Host = host
		modified = true
	}
	if schema != "" {
		tads.Client.Cfg.Scheme = schema
		modified = true
	}
	if modified {
		tads.Client.Cfg.BasePath = tads.Client.Cfg.Scheme + "://" + tads.Client.Cfg.Host + "/" + tads.ApiVersion
	}
	return tads
}

// UseSandbox ...
func (tads *SDKClient) UseSandbox() *SDKClient {
	return tads.SetHost("sandbox-api.e.qq.com", "https")
}

// UseProduction ...
func (tads *SDKClient) UseProduction() *SDKClient {
	return tads.SetHost("api.e.qq.com", "https")
}

// GetAccessToken ...
func (tads *SDKClient) GetAccessToken() string {
	return tads.Config.AccessToken
}

// SetAccessToken ...
func (tads *SDKClient) SetAccessToken(accessToken string) *SDKClient {
	tads.Config.AccessToken = accessToken
	return tads
}

// GetVersion ...
func (tads *SDKClient) GetVersion() string {
	return tads.Version
}

// IsDebug ...
func (tads *SDKClient) IsDebug() bool {
	return tads.Config.IsDebug
}

// SetDebug ...
func (tads *SDKClient) SetDebug(debug bool) *SDKClient {
	tads.Config.IsDebug = debug
	return tads
}

// GetDebugFile ...
func (tads *SDKClient) GetDebugFile() string {
	return tads.Config.DebugFile
}

// SetDebugFile ...
func (tads *SDKClient) SetDebugFile(debugFile string) *SDKClient {
	tads.Config.DebugFile = debugFile
	return tads
}

// IsMonitor ...
func (tads *SDKClient) IsMonitor() bool {
	return !tads.Config.SkipMonitor
}

// SetHeaders ...
func (tads *SDKClient) SetHeaders(header http.Header) *SDKClient {
	tads.Config.GlobalConfig.HttpOption.Header = header
	return tads
}

// SetHeader ...
func (tads *SDKClient) SetHeader(key string, value string) *SDKClient {
	if tads.Config.GlobalConfig.HttpOption.Header == nil {
		tads.Config.GlobalConfig.HttpOption.Header = http.Header{}
	}
	tads.Config.GlobalConfig.HttpOption.Header.Set(key, value)
	return tads
}

// SetMonitor ...
func (tads *SDKClient) SetMonitor(monitor bool) *SDKClient {
	tads.Config.SkipMonitor = !monitor
	return tads
}

// AppendMiddleware ...
func (tads *SDKClient) AppendMiddleware(middleware Middleware) *SDKClient {
	tads.middlewareList = append(tads.middlewareList, middleware)
	return tads
}

// RoundTrip ...
func (tads *SDKClient) RoundTrip(req *http.Request) (rsp *http.Response, err error) {
	beforeFunc := func(req *http.Request) (rsp *http.Response, err error) {
		return tads.RoundTripper.RoundTrip(req)
	}
	middlewareCount := len(tads.middlewareList)
	// 逆序遍历
	for i := middlewareCount - 1; i >= 0; i-- {
		beforeFunc = tads.GenMiddlewareHandleFunc(tads.middlewareList[i], beforeFunc)
	}
	return beforeFunc(req)
}

// GenMiddlewareHandleFunc ...
func (tads *SDKClient) GenMiddlewareHandleFunc(
	middleware Middleware,
	beforeFunc func(req *http.Request) (rsp *http.Response, err error),
) func(req *http.Request) (rsp *http.Response, err error) {
	return func(req *http.Request) (rsp *http.Response, err error) {
		return middleware.Handle(req, beforeFunc)
	}
}

// SetNameService ...
