package justgo

import (
	"context"
)

var justGoHttp *HttpInterface
var justGoCli *CliInterface
var appContext context.Context

func Start() {
	Log.Info("starting justgo")
	Initialise()
	RunAppInterfaces()
}

func Initialise() {
	appContext = context.Background()
	Config.Load()
	Log.Load()
	Instrument.Load()
	enableHttpInterface := Config.GetBooleanOrDefault(ConfigKey.DEFAULT_INTERFACE_HTTP_ENABLED, true)
	if enableHttpInterface {
		justGoHttp = GetDefaultHttpInterface()
		RegisterInterface(justGoHttp)
	}
}


