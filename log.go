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
}

func init() {
	Log = &justGoLog{&logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}}
}
