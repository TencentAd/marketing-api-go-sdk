package ads

import (
	"net/http"
	"runtime"
)

// HeaderMiddleware ...
type HeaderMiddleware struct {
	tads *SDKClient
}

// Handle ...
func (h *HeaderMiddleware) Handle(
	req *http.Request,
	next func(req *http.Request) (rsp *http.Response, err error),
) (rsp *http.Response, err error) {
	tads := h.tads
	if tads.Config.GlobalConfig.HttpOption.Header != nil {
		for k, v := range tads.Config.GlobalConfig.HttpOption.Header {
			req.Header[k] = v
		}
	}
	if host := req.Header.Get("Host"); host != "" {
		req.Host = host
	}
	if tads.IsMonitor() {
		req.Header.Add("X-SDK-LANGUAGE", "GO")
		req.Header.Add("X-SDK-LANGUAGE-VERSION", runtime.Version())
		req.Header.Add("X-SDK-VERSION", tads.GetVersion())
		req.Header.Add("X-SDK-OS", runtime.GOOS)
	}
	rsp, err = next(req)
	return rsp, err
}
