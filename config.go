package justgo

var Config *config

type config struct {

}

func (c *config) Load() {

}

func init() {
	Config = &config{}
}