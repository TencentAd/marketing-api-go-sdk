package main

import (
	"encoding/json"
	"fmt"

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
)

// RefreshAccessTokenExample ...
type RefreshAccessTokenExample struct {
	TAds           *ads.SDKClient
	AccessToken    string
	ClientId       int64
	ClientSecret   string
	GrantType      string
	OauthTokenOpts *api.OauthTokenOpts
}

// Init ...
func (e *RefreshAccessTokenExample) Init() {
	e.TAds = ads.Init(&config.SDKConfig{
		IsDebug: true,
	})
	// oauth/token不提供沙箱环境
	e.TAds.UseProduction()
	// YOUR CLIENT ID
	e.ClientId = int64(0)
	e.ClientSecret = "YOUR CLIENT SECRET"
	e.GrantType = "refresh_token"
	e.OauthTokenOpts = &api.OauthTokenOpts{
		RefreshToken: optional.NewString("YOUR REFRESH TOKEN"),
	}
}

// RunExample ...
func (e *RefreshAccessTokenExample) RunExample() {
	tads := e.TAds

	// change ctx as needed
	ctx := *tads.Ctx
	response, _, err := tads.Oauth().Token(ctx, e.ClientId, e.ClientSecret, e.GrantType, e.OauthTokenOpts)

	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			// When Api returns an error
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Resopnse error:", string(errStr))
		} else {
			// When validation fails or other local issues
			fmt.Println("Error:", err)
		}
	} else {
		// 从返回里获得AccessToken并设置到tads中
		tads.SetAccessToken(response.AccessToken)
	}
}

func main() {
	e := &RefreshAccessTokenExample{}
	e.Init()
	e.RunExample()
}
