package ads

import (
	"net/http"
	"net/url"
)

type DiffHostMiddleware struct {
	tads *SDKClient
}

func (d *DiffHostMiddleware) Handle(
	req *http.Request,
	next func(req *http.Request) (rsp *http.Response, err error),
) (rsp *http.Response, err error) {
	interfaceUrl := req.URL.Path[len("/"+d.tads.ApiVersion):]
	if specialUrl, ok := SpecialUrlInterfaceList[interfaceUrl]; ok {
		if specialUrl.Host != "" {
			u, _ := url.Parse(specialUrl.Host)
			req.Header.Set("Host", u.Host)
			uri := u.Host + u.Path
			req.URL.Host = uri
			req.Host = uri
		}
		if specialUrl.Path != "" {
			req.URL.Path = specialUrl.Path
		}
	}
	rsp, err = next(req)
	return rsp, err
}
