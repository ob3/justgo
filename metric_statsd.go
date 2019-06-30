package justgo

import (
	"fmt"
	"gopkg.in/alexcesaro/statsd.v2"
	"time"
)

var statsD *statsd.Client
type metricStatsD struct {
	Host string
	Port int64
	Prefix string
	FlushPeriod time.Duration
	AppName string
}

func (mStatsD *metricStatsD) Increment(key string) {
	statsD.Increment(key)
}

func (mStatsD *metricStatsD) Init() error {

	address := statsd.Address(fmt.Sprintf("%s:%d", mStatsD.Host, mStatsD.Port))
	prefix := statsd.Prefix(mStatsD.AppName)
	flushPeriod := statsd.FlushPeriod(mStatsD.FlushPeriod)

	var err error
	statsD, err = statsd.New(address, prefix, flushPeriod)
	if err != nil {
		return fmt.Errorf("error initiating statsD %+v", err)
	}

	return nil
}



func (mStatsD *metricStatsD) CloseStatsDClient() {
	if statsD != nil {
		statsD.Close()
	}
}

