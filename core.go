package justgo

import (
	"context"
)

var justGoHttp *HttpInterface
var justGoCli *CliInterface
var appContext context.Context

func Start() {
	log.Info("starting justgo")
	Initialise()
	RunAppInterfaces()
}

func Initialise() {
	appContext = context.Background()
	Config.Load()
	log.Load()
	defaultInstrument.Load()
	defaultStorage.Load()
	enableHttpInterface := Config.GetBooleanOrDefault(ConfigKey.DEFAULT_INTERFACE_HTTP_ENABLED, true)
	if enableHttpInterface {
		justGoHttp = GetDefaultHttpInterface()
		RegisterInterface(justGoHttp)
	}

}


