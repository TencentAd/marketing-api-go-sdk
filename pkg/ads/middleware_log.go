package ads

import (
	"net/http"
	"net/http/httputil"
	"os"
)

type LogMiddleware struct {
	tads *SDKClient
}

func (l *LogMiddleware) Handle(req *http.Request, next func(req *http.Request) (rsp *http.Response, err error)) (rsp *http.Response, err error) {
	var debugFile *os.File
	tads := l.tads
	if tads.IsDebug() {
		if tads.GetDebugFile() == "" {
			debugFile = os.Stdout
		} else {
			debugFile, err = os.OpenFile(tads.GetDebugFile(), os.O_APPEND, 0666)
		}
		if request, err := httputil.DumpRequestOut(req, true); err == nil {
			if debugFile != nil {
				_, err = debugFile.Write(request)
				_, err = debugFile.WriteString("\n")
			}
		}
	}
	rsp, err = next(req)
	if tads.IsDebug() {
		if response, err := httputil.DumpResponse(rsp, true); err == nil {
			if debugFile != nil {
				_, err = debugFile.Write(response)
				_, err = debugFile.WriteString("\n")
			}
		}
	}

	return rsp, err
}
