package justgo

import (
	"github.com/spf13/viper"
	"os"
	"runtime/debug"
	"strconv"
)

var Config *config

var DEFAULT_CONFIGS = map[string]string{
	"APP_PORT": "8080",
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

func (c *config) GetString(key string) string {
	return c.fatalGetString(key)
}

func (c *config) Load() {
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

func (c *config)fatalGetString(key string) string {
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