package justgo

type UserInterface interface {
	Serve()
}

func RunUserInterface() {
	httpServer := &HttpInterface{}
	httpServer.Serve()
}
