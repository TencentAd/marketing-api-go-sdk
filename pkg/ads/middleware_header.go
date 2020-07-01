package ads

import (
	"net/http"
	"runtime"
)

type HeaderMiddleware struct {
	tads *SDKClient
}

func (h *HeaderMiddleware) Handle(req *http.Request, next func(req *http.Request) (rsp *http.Response, err error)) (rsp *http.Response, err error) {
	tads := h.tads
	if tads.Config.GlobalConfig.HttpOption.Header != nil {
		req.Header = tads.Config.GlobalConfig.HttpOption.Header
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
