package main

import (
	"fmt"
	"github.com/ob3/justgo"
	"net/http"
)

func main() {
	justgo.AddRoute(http.MethodGet, "/with-middleware", headerPrinterHandler, middleWareDummyOne, otherAuthHandler)
	justgo.AddRoute(http.MethodGet, "/no-middleware", headerPrinterHandler)
	justgo.Start()
}

func middleWareDummyOne(next http.Handler) http.Handler {
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		justgo.Log.Println("handling middleware 1")
		w.Header().Add("x-middleware-dummy-one", "true")
		next.ServeHTTP(w, r)
	}))
}

func otherAuthHandler(next http.Handler) http.Handler {
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		justgo.Log.Println("handling other auth handler")
		w.Header().Add("x-middleware-dummy-two", "any value")
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
