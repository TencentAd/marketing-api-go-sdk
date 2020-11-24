package ads

import (
	"net/http"
)

// Middleware ...
type Middleware interface {
	Handle(req *http.Request, next func(req *http.Request) (rsp *http.Response, err error)) (rsp *http.Response, err error)
}
