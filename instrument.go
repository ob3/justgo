package justgo

import "github.com/newrelic/go-agent"

var Instrument *instrument

type instrument struct {
	NewRelic newrelic.Application
	metrics  []*Metric
}

func (instrument *instrument) Load() {
	appName := Config.GetStringOrDefault("APP_NAME", "Undefined App Name")
	newRelicLicense := Config.GetStringOrDefault("NEWRELIC_LICENSE", "")
	if newRelicLicense != "" {
		cfg := newrelic.NewConfig(appName, newRelicLicense)
		nrApp, err := newrelic.NewApplication(cfg)
		if err != nil {
			Log.Error("disabling newrelic ", err)
		}
		Instrument.NewRelic = nrApp
	}
}

func (instrument *instrument) AddMetric(metric Metric) {
	instrument.metrics = append(instrument.metrics, &metric)
}

func (instrument *instrument) Increment(key string) {

	for _, metric := range instrument.metrics {
		metricPointer := *metric
		metricPointer.Increment(key)
	}
}

func init() {
	Instrument = &instrument{}
}
