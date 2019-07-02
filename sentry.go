package justgo

// credits https://gist.github.com/belak/3c205bc30a12cef954f57b542abcfb54

import (
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
)

var severityMap = map[logrus.Level]raven.Severity{
	logrus.DebugLevel: raven.DEBUG,
	logrus.InfoLevel:  raven.INFO,
	logrus.WarnLevel:  raven.WARNING,
	logrus.ErrorLevel: raven.ERROR,
	logrus.FatalLevel: raven.FATAL,
	logrus.PanicLevel: raven.FATAL,
}

type sentryHook struct {
	client *raven.Client
	levels []logrus.Level
}

func (h *sentryHook) Levels() []logrus.Level {
	return h.levels
}

func (h *sentryHook) Fire(e *logrus.Entry) error {
	// Using NewPacket is a little uglier, but it ensures all the
	// required fields are set.
	p := raven.NewPacket(e.Message)
	p.Level = severityMap[e.Level]
	p.Timestamp = raven.Timestamp(e.Time)

	// e.Data has all the variables we registered when logging
	// this, so we loop through them and make sure to grab the
	// error separately.
	var err error
	for k, v := range e.Data {
		if k == logrus.ErrorKey {
			err = v.(error)
		} else {
			p.Extra[k] = v
		}
	}

	// If there wasn't an error, we create one based on the
	// message. This needs to be done so we can have a
	// raven.Exception which will actually be logged to sentry in
	// a sane way.
	if err == nil {
		err = errors.New(e.Message)
	}

	// Create a new stack trace and exception to store in sentry
	//
	// Note that raven.NewStacktrace is currently set to ignore
	// the first frame of the Stacktrace (this function) and will
	// grab 3 lines of context.
	stack := raven.NewStacktrace(1, 3, nil)
	exc := raven.NewException(err, stack)

	// Add the exception to the sentry packet
	p.Interfaces = append(p.Interfaces, exc)

	// Send the packet we just built to sentry.
	h.client.Capture(p, nil)

	return nil
}


func NewSentryHook(dsn string, level []logrus.Level) (*sentryHook, error) {
	client, e := raven.New(dsn)
	if e != nil {
		return nil, e
	}
	return &sentryHook{levels: level, client: client}, nil
}
