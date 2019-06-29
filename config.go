package justgo

import (
	"os"
	"runtime/debug"
	"strconv"

	"github.com/spf13/viper"
)

var Config *config

var DEFAULT_CONFIGS = map[string]string{
	"APP_PORT":  "8080",
	"LOG_LEVEL": "debug",
}

type config struct {
}

func (c *config) GetInt(key string) int64 {

	configStr := c.fatalGetString(key)
	config, err := strconv.ParseInt(configStr, 10, 64)
	if err != nil {
		Log.
			WithField("key", key).
			WithField("err", err).
			Fatalf("can't parsing key %s error: %s", key, err.Error())
	}
	return config
}

func (c *config) GetStringOrDefault(key string, fallback string) string {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		Log.WithField("key", key).WithField("value", fallback).Info("config not found using default")
		return fallback
	}
	return c.GetString(key)

}
func (c *config) GetString(key string) string {
	return c.fatalGetString(key)
}

func (c *config) Load(paths ...string) {
	for key, value := range DEFAULT_CONFIGS {
		viper.SetDefault(key, value)
	}
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func init() {
	Config = &config{}
}

func (c *config) fatalGetString(key string) string {
	checkKey(key)
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func checkKey(key string) {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		debug.PrintStack()
		Log.Fatalf("%s key is not set", key)
	}
}
