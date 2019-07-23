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
	WaitForShutdown()
}

func WaitForShutdown() {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-shutdownSignal
	log.Info("shutting down")
	// Finish all apis being served and shutdown gracefully
	for _, appInterface := range appInterfaces {
		singleInterface := *appInterface
		go singleInterface.ShutDown()
	}
	log.Info("shutdown complete")
}

func RegisterInterface(appInterface AppInterface) {
	appInterfaces = append(appInterfaces, &appInterface)
}
