package justgo

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpInterface struct {
	http.Server
}

func (httpInterface HttpInterface) Serve() {
	//server := &http.Server{Addr: portInfo, Handler: n}
	server := &HttpInterface{http.Server{}}
	go listenServer(server)
	waitForShutdown(server)
}

func listenServer(apiServer *HttpInterface) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		Log.Fatal(err.Error())
	}
}

func waitForShutdown(apiServer *HttpInterface) {
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

