package justgo

import (
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/_integrations/nrgorilla/v1"
	"net/http"
)

var router *mux.Router

func getRouter() http.Handler {
	if router == nil {
		InstantiateNewRouter()
	}
	return router
}

func InstantiateNewRouter() {
	if Instrument.NewRelic != nil {
		router = nrgorilla.InstrumentRoutes(mux.NewRouter(), Instrument.NewRelic)
	} else {
		router = mux.NewRouter()
	}
}


func AddRoute(method string, pattern string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	i := getRouter().(*mux.Router)

	if len(middlewares) == 0 {
		i.HandleFunc(pattern, handler).Methods(method)
		return
	}

	var chainedMiddleWares http.Handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		if i == len(middlewares) - 1 {
			chainedMiddleWares = middlewares[i](handler)
		} else {
			chainedMiddleWares = middlewares[i](chainedMiddleWares)
		}
	}
	i.HandleFunc(pattern, chainedMiddleWares.ServeHTTP).Methods(method)
}
