package justgo

func Start() {
	Config.Load()
	Log.Load()
	Log.Info("starting")
	RunUserInterface()
}
