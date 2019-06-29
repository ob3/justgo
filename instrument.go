package justgo

import newrelic "github.com/newrelic/go-agent"

var Instrument *instrument

type instrument struct {
	NewRelic *newrelic.Application
}

func (instrument *instrument) Load() {
	appName := Config.GetStringOrDefault("APP_NAME", "Undefined App Name")
	newRelicLicense := Config.GetStringOrDefault("NEWRELIC_LICENSE", "")
	if newRelicLicense != "" {
		cfg := newrelic.NewConfig(appName, newRelicLicense)
		nrApp, err := newrelic.NewApplication(cfg)
		if err != nil {
			Log.Error(err)
		}
		Instrument.NewRelic = &nrApp
	}
}

func init() {
	Instrument = &instrument{}
}
