package justgo

var ConfigKey configKey

type configKey struct {
	APP_NAME  string
	APP_PORT  string
	LOG_LEVEL string

	NEWRELIC_ENABLED string
	NEWRELIC_LICENSE string

	STATSD_ENABLED string
	STATSD_HOST    string
	STATSD_PORT    string
	STATSD_PREFIX  string

	STATSD_FLUSH_PERIOD_IN_SECONDS string

	SENTRY_ENABLED string
	SENTRY_DSN     string
}

func init() {
	ConfigKey = configKey{
		APP_NAME:  "APP_NAME",
		LOG_LEVEL: "LOG_LEVEL",
		APP_PORT:  "APP_PORT",

		NEWRELIC_ENABLED: "NEWRELIC_ENABLED",
		NEWRELIC_LICENSE: "NEWRELIC_LICENSE",

		STATSD_HOST:                    "STATSD_HOST",
		STATSD_PORT:                    "STATSD_PORT",
		STATSD_PREFIX:                  "STATSD_PREFIX",
		STATSD_FLUSH_PERIOD_IN_SECONDS: "STATSD_FLUSH_PERIOD_IN_SECONDS",
		STATSD_ENABLED:                 "STATSD_ENABLED",

		SENTRY_ENABLED: "SENTRY_ENABLED",
		SENTRY_DSN:     "SENTRY_DSN",
	}
}
