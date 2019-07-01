package justgo

import (
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/_integrations/nrgorilla/v1"
	"net/http"
)

var router *mux.Router
var mandatoryMiddleWares []func(handler http.Handler) http.Handler
func GetRouter() http.Handler {
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

func AlwaysUseMiddleware(middleWares ... func (handler http.Handler) http.Handler) {
	mandatoryMiddleWares = append(mandatoryMiddleWares, middleWares...)
}

func AddRoute(method string, pattern string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	i := GetRouter().(*mux.Router)

	joinedMiddleWares := append(mandatoryMiddleWares, middlewares...)
	if len(joinedMiddleWares) == 0 {
		i.HandleFunc(pattern, handler).Methods(method)
		return
	}

	var chainedMiddleWares http.Handler
	for i := len(joinedMiddleWares) - 1; i >= 0; i-- {
		if i == len(joinedMiddleWares) - 1 {
			chainedMiddleWares = joinedMiddleWares[i](handler)
		} else {
			chainedMiddleWares = joinedMiddleWares[i](chainedMiddleWares)
		}
	}
	i.HandleFunc(pattern, chainedMiddleWares.ServeHTTP).Methods(method)
}
