package justgo

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *log

type log struct {
	*logrus.Logger
}

func (log log) Load() {

}

func init() {
	Log = &log{&logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}}
}
