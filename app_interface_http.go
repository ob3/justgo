package justgo

import (
	"context"
	golangLogger "log"
	"net/http"
)

type HttpInterface struct {
	Handler http.Handler
	server *http.Server
}

func (httpInterface HttpInterface) Serve() {
	logWriter := log.Writer()
	defer logWriter.Close()

	httpInterface.server = &http.Server{
		Addr:     ":"+Config.GetStringOrDefault(ConfigKey.APP_PORT, "8080"),
		ErrorLog: golangLogger.New(logWriter, "", 0),
		Handler: httpInterface.Handler,
	}

	log.Printf("listening on %s", httpInterface.server.Addr)
	listenServer(httpInterface.server)
}

func (httpInterface HttpInterface) ShutDown() {
	log.Info("API server shutting down")
	// Finish all apis being served and shutdown gracefully
	if httpInterface.server != nil {
		shutdown := httpInterface.server.Shutdown(context.Background())
		if shutdown != nil {
			log.Info(shutdown)
		}
	}
	log.Info("API server shutdown complete")
}

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal(err.Error())
	}
}

func GetDefaultHttpInterface() *HttpInterface {

	router := GetRouter()
	AddRoute(http.MethodGet, "/ping", pingHandler)
	return &HttpInterface{
		Handler: router,
	}
}

