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

}

func init() {
	Log = &justGoLog{&logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}}
}
