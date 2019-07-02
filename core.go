package justgo

var justGoHttp *HttpInterface

func Start() {
	Log.Info("starting justgo")
	Initialise()
	RunAppInterface()
}

func Initialise() {
	Config.Load()
	Log.Load()
	Instrument.Load()
	if len(appInterfaces) == 0 {
		justGoHttp = getDefaultHttpInterface()
		RegisterInterface(justGoHttp)
	}
}


