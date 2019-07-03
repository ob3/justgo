package justgo

import "context"

var justGoHttp *HttpInterface
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
	if len(appInterfaces) == 0 {
		justGoHttp = GetDefaultHttpInterface()
		RegisterInterface(justGoHttp)
	}
}


