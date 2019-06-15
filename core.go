package justgo

var Http *HttpInterface

func Start() {
	Log.Info("starting justgo")
	Config.Load()
	Log.Load()

	if len(appInterfaces) == 0 {
		Http = &HttpInterface{}
		RegisterInterface(Http)
	}
	RunAppInterface()
}
