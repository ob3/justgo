package main

import (
	"fmt"
	"net/http"

	"github.com/ob3/justgo"
)

func main() {
	// add route with middleware
	justgo.AddRoute(http.MethodGet, "/with-middleware", headerPrinterHandler, middleWareDummyOne, otherAuthHandler)

	// add route without any middleware
	justgo.AddRoute(http.MethodGet, "/no-middleware", headerPrinterHandler)

	// config in code
	justgo.Config.Add("APP_NAME", "My Volatile Config")

	// set custom config path
	justgo.Config.ConfigFile("./sample/anything.yml")


	// add custom metrics
	justgo.Config.Add("METRIC_PREFIX_DUMMY", "MY_PREFIX_")
	metric := &customMetric{}
	justgo.Instrument.AddMetric(metric)

	justgo.Start()

}

type customMetric struct {
}

func (cm *customMetric) Increment(key string) {
	prefix := justgo.Config.GetString("METRIC_PREFIX_DUMMY")
	justgo.Log.Info("incrementing metric: ", prefix, key)
}

func middleWareDummyOne(next http.Handler) http.Handler {
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		justgo.Log.Println("handling middleware 1")
		w.Header().Add("x-middleware-dummy-one", "true")
		justgo.Instrument.Increment("instrument-key-1")
		next.ServeHTTP(w, r)
	}))
}

func otherAuthHandler(next http.Handler) http.Handler {
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		justgo.Log.Println("handling other auth handler")
		w.Header().Add("x-middleware-dummy-two", "any value")
		justgo.Instrument.Increment("instrument-key-2")
		next.ServeHTTP(w, r)
	}))
}

func headerPrinterHandler(w http.ResponseWriter, request *http.Request) {
	justgo.Log.Println("handling poooong")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	stringHeaders := fmt.Sprintf("a %s", w.Header())
	w.Write([]byte(stringHeaders))
}
