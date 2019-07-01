package justgo

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpInterface struct {
	Handler http.Handler
}

func (httpInterface HttpInterface) Serve() {
	logWriter := Log.Writer()
	defer logWriter.Close()

	server := &http.Server{
		Addr:     ":"+Config.GetString(ConfigKey.APP_PORT),
		ErrorLog: log.New(logWriter, "", 0),
		Handler: httpInterface.Handler,
	}


	go listenServer(server)
	Log.Printf("listening on %s", server.Addr)
	waitForShutdown(server)
}

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		Log.Fatal(err.Error())
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig
	Log.Info("API server shutting down")
	// Finish all apis being served and shutdown gracefully
	apiServer.Shutdown(context.Background())
	Log.Info("API server shutdown complete")
}

func getDefaultHttpInterface() *HttpInterface {

	router := GetRouter()
	AddRoute(http.MethodGet, "/ping", pingHandler)
	return &HttpInterface{
		Handler: router,
	}
}

