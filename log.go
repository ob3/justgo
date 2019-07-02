package justgo

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *justGoLog

type justGoLog struct {
	*logrus.Logger
}

func (log *justGoLog) Load() {
	sentryEnabled := Config.GetBooleanOrDefault(ConfigKey.SENTRY_ENABLED, false)
	level, e := logrus.ParseLevel(Config.GetString(ConfigKey.LOG_LEVEL))
	if e != nil {
		Log.Fatal(e)
	}

	Log = &justGoLog{&logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}}

	if sentryEnabled {
		levels := []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel, logrus.PanicLevel, logrus.FatalLevel}
		sentry, e := NewSentryHook(Config.GetString(ConfigKey.SENTRY_DSN), levels)
		if e != nil {
			Log.WithField("error", e).Error("failed to initialise sentry. sentry reporting not enabled")
		} else {
			Log.AddHook(sentry)
		}

	}
}

func init() {
	Log = &justGoLog{&logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}}
}
