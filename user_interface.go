package justgo

type ApplicationInterface interface {
	Serve()
}

var appInterfaces []*ApplicationInterface

func RunAppInterface() {
	for _, appInterface := range appInterfaces {
		singleInterface := *appInterface
		singleInterface.Serve()
	}
}

func RegisterInterface(appInterface ApplicationInterface) {
	appInterfaces = append(appInterfaces, &appInterface)
}
