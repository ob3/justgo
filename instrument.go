package justgo

import (
	"github.com/newrelic/go-agent"
	"time"
)

var Instrument *instrument

type instrument struct {
	NewRelic newrelic.Application
	metrics  []*Metric
}

func (instrument *instrument) Load() {

	enableNewRelic := Config.GetBooleanOrDefault(ConfigKey.NEWRELIC_ENABLED, false)
	enableStatsD := Config.GetBooleanOrDefault(ConfigKey.STATSD_ENABLED, false)

	if enableNewRelic {
		appName := Config.GetStringOrDefault(ConfigKey.APP_NAME, "Undefined App Name")
		newRelicLicense := Config.GetStringOrDefault(ConfigKey.NEWRELIC_LICENSE, "")
		if newRelicLicense != "" {
			cfg := newrelic.NewConfig(appName, newRelicLicense)
			nrApp, err := newrelic.NewApplication(cfg)
			if err != nil {
				Log.Error("disabling newrelic ", err)
			}
			Instrument.NewRelic = nrApp
		}
	}

	if len(instrument.metrics) == 0 && enableStatsD {
		defaultStatsD := getDefaultStatsD()
		err := defaultStatsD.Init()
		if err == nil {
			instrument.AddMetric(defaultStatsD)
		}
	}
}

func getDefaultStatsD() *metricStatsD {
	return &metricStatsD{
		Host: Config.GetStringOrDefault(ConfigKey.STATSD_HOST, "localhost"),
		Port: Config.GetIntOrDefault(ConfigKey.STATSD_PORT, 8125),
		Prefix: Config.GetStringOrDefault(ConfigKey.STATSD_PREFIX, ""),
		FlushPeriod: time.Duration(Config.GetIntOrDefault(ConfigKey.STATSD_FLUSH_PERIOD_IN_SECONDS, 20)) * time.Second,
		AppName: Config.GetStringOrDefault(ConfigKey.APP_NAME, "Undefined App Name"),
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
