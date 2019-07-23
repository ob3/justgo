package justgo

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *justGoLog

type justGoLog struct {
	*logrus.Logger
}

func (log *justGoLog) Load() {
	sentryEnabled := Config.GetBooleanOrDefault(ConfigKey.SENTRY_ENABLED, false)
	level, e := logrus.ParseLevel(Config.GetStringOrDefault(ConfigKey.LOG_LEVEL, "info"))
	if e != nil {
		log.Fatal(e)
	}

	log = &justGoLog{&logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}}

	if sentryEnabled {
		levels := []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel, logrus.PanicLevel, logrus.FatalLevel}
		sentry, e := NewSentryHook(Config.GetString(ConfigKey.SENTRY_DSN), levels)
		if e != nil {
			log.WithField("error", e).Error("failed to initialise sentry. sentry reporting not enabled")
		} else {
			log.AddHook(sentry)
		}

	}
}

func GetLogger() *justGoLog {
	return log
}

func init() {
	log = &justGoLog{&logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}}
}
