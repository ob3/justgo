package justgo

import (
	"os"
	"os/signal"
	"syscall"
)

type AppInterface interface {
	Serve()
	ShutDown()
}

var appInterfaces []*AppInterface

func RunAppInterfaces() {
	for _, appInterface := range appInterfaces {
		singleInterface := *appInterface
		go singleInterface.Serve()
	}
	//time.Sleep(2)
	WaitForShutdown()
}

func WaitForShutdown() {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-shutdownSignal
	Log.Info("shutting down")
	// Finish all apis being served and shutdown gracefully
	for _, appInterface := range appInterfaces {
		singleInterface := *appInterface
		go singleInterface.ShutDown()
	}
	Log.Info("shutdown complete")
}

func RegisterInterface(appInterface AppInterface) {
	appInterfaces = append(appInterfaces, &appInterface)
}
