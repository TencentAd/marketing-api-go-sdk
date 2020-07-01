package ads

import (
	"net/http"
)

type Middleware interface {
	Handle(req *http.Request, next func(req *http.Request) (rsp *http.Response, err error)) (rsp *http.Response, err error)
}
