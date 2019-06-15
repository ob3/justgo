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

func AddRoute(method string, pattern string, handler http.HandlerFunc) {
	router.HandleFunc(pattern, handler).Methods(method)
}
