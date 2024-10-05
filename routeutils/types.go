package routeutils

import "net/http"

type HandlerStruct struct {
	Url         string
	HandlerName func(w http.ResponseWriter, r *http.Request)
	Methods     []string
}
