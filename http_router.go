package justgo

import (
	"github.com/gorilla/mux"
	"net/http"
)

var router *mux.Router

func getRouter() http.Handler {
	if router == nil {
		router = mux.NewRouter()
	}

	return router
}


func AddRoute(method string, pattern string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	i := getRouter().(*mux.Router)

	if len(middlewares) == 0 {
		i.HandleFunc(pattern, handler).Methods(method)
		return
	}

	var chainedMiddleWares http.Handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		Log.Printf("middleware index %d", i)
		if i == len(middlewares) - 1 {
			chainedMiddleWares = middlewares[i](handler)
		} else {
			chainedMiddleWares = middlewares[i](chainedMiddleWares)
		}
	}
	i.HandleFunc(pattern, chainedMiddleWares.ServeHTTP).Methods(method)
}
