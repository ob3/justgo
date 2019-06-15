package justgo

var justGoHttp *HttpInterface

func Start() {
	Log.Info("starting justgo")
	Config.Load()
	Log.Load()

	if len(appInterfaces) == 0 {
		justGoHttp = getDefaultHttpInterface()
		RegisterInterface(justGoHttp)
	}

	RunAppInterface()
}


