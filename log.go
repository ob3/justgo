package justgo

import (
	"fmt"
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
	fmt.Println("init Log")
	Log = &log{&logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}}
}
