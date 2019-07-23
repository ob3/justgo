package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
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

	// disable default http handler
	justgo.Config.Add(justgo.ConfigKey.DEFAULT_INTERFACE_HTTP_ENABLED, "false")

	// add your own http handler
	justgo.RegisterInterface(justgo.GetDefaultHttpInterface())

	// add cli handler via telnet
	cliInterface := &justgo.CliInterface{Address: ":12345"}

	// add command on telnet command
	cliInterface.AddCommand("test", echo)
	cliInterface.AddCommand("panic", panicCommand)
	cliInterface.AddCommand("fatal", fatalCommand)

	// register interface
	justgo.RegisterInterface(cliInterface)

	// add custom metrics
	justgo.Config.Add("METRIC_PREFIX_DUMMY", "MY_PREFIX_")
	metric := &customMetric{}
	justgo.GetInstrument().AddMetric(metric)

	// enable database
	justgo.Config.Add("DB_DRIVER", "postgres")
	justgo.Config.Add("DB_CONNECTION_STRING", "dbname=postgres user=postgres password=abcdef host=localhost sslmode=disable")

	db := justgo.GetDB()
	justgo.Log.Info(db)


	justgo.Start()

}

type customMetric struct {
}

func echo(param string) string {
	return fmt.Sprintf("param1 is: %s", param)
}

func panicCommand(param string) string {
	panic("this is paniced command")
}

func fatalCommand(param string) string {
	justgo.Log.Fatal("fatal command")
	return param
}

func (cm *customMetric) Increment(key string) {
	prefix := justgo.Config.GetString("METRIC_PREFIX_DUMMY")
	justgo.Log.Info("incrementing metric: ", prefix, key)
}

func middleWareDummyOne(next http.Handler) http.Handler {
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		justgo.Log.Println("handling middleware 1")
		w.Header().Add("x-middleware-dummy-one", "true")
		justgo.GetInstrument().Increment("instrument-key-1")
		next.ServeHTTP(w, r)
	}))
}

func otherAuthHandler(next http.Handler) http.Handler {
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		justgo.Log.Println("handling other auth handler")
		w.Header().Add("x-middleware-dummy-two", "any value")
		justgo.GetInstrument().Increment("instrument-key-2")
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
