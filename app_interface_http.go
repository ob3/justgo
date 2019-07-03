package justgo

import (
	"context"
	"log"
	"net/http"
)

type HttpInterface struct {
	Handler http.Handler
	server *http.Server
}

func (httpInterface HttpInterface) Serve() {
	logWriter := Log.Writer()
	defer logWriter.Close()

	httpInterface.server = &http.Server{
		Addr:     ":"+Config.GetString(ConfigKey.APP_PORT),
		ErrorLog: log.New(logWriter, "", 0),
		Handler: httpInterface.Handler,
	}

	Log.Printf("listening on %s", httpInterface.server.Addr)
	listenServer(httpInterface.server)
}

func (httpInterface HttpInterface) ShutDown() {
	Log.Info("API server shutting down")
	// Finish all apis being served and shutdown gracefully
	if httpInterface.server != nil {
		shutdown := httpInterface.server.Shutdown(context.Background())
		if shutdown != nil {
			Log.Info(shutdown)
		}
	}
	Log.Info("API server shutdown complete")
}

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		Log.Fatal(err.Error())
	}
}

func GetDefaultHttpInterface() *HttpInterface {

	router := GetRouter()
	AddRoute(http.MethodGet, "/ping", pingHandler)
	return &HttpInterface{
		Handler: router,
	}
}

