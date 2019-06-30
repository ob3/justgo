package justgo

import (
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

var Config *config

var DEFAULT_CONFIGS = map[string]string{
	ConfigKey.APP_PORT:  "8080",
	ConfigKey.LOG_LEVEL: "debug",
}

var configFile = ""

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

func (c *config) Add(key string, value string) {
	DEFAULT_CONFIGS[key] = value
}

func (c *config) ConfigFile(path string) {
	configFile = path
}

func (c *config) Load() {
	for key, value := range DEFAULT_CONFIGS {
		viper.SetDefault(key, value)
	}

	viper.AutomaticEnv()

	if configFile != "" {
		Log.Info("using config file ", configFile)
		dir, file := filepath.Split(configFile)

		fileSplitted := strings.Split(file, ".")
		viper.SetConfigName(fileSplitted[0])
		viper.AddConfigPath(dir)
		viper.SetConfigType(fileSplitted[1])
	} else {
		viper.SetConfigName("application")
		viper.AddConfigPath("./")
		viper.AddConfigPath("../")
		viper.AddConfigPath("../../")
		viper.SetConfigType("yaml")
	}

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

func (c *config) GetIntOrDefault(key string, fallback int64) int64 {
	value := c.GetStringOrDefault(key, "")
	if value == "" {
		return fallback
	}

	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return fallback
	}
	return intVal

}

func (c *config) GetBooleanOrDefault(key string, fallback bool) bool {
	value := c.GetStringOrDefault(key, "")
	if value == "" {
		return fallback
	} else {
		intVal, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return intVal
	}
}

func checkKey(key string) {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		debug.PrintStack()
		Log.Fatalf("%s key is not set", key)
	}
}
