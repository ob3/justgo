package justgo

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *justGoLog

type justGoLog struct {
	*logrus.Logger
}

func (log justGoLog) Load() {
	level, e := logrus.ParseLevel(Config.GetString("LOG_LEVEL"))
	if e != nil {
		Log.Fatal("invalid log LOG_LEVEL value")
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
