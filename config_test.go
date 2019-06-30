package justgo

import (
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/spf13/viper"
	"gotest.tools/assert"
)

func TestGetIntShouldReturnIntIfKeyValueIsInteger(t *testing.T) {
	key := "dummy_key"
	os.Setenv(key, "8080")
	defer os.Unsetenv(key)
	Config.Load()
	value := Config.GetInt(key)
	assert.Equal(t, value, int64(8080))
}

func TestGetIntShouldReturnIntIfKeyValueIsNotInteger(t *testing.T) {
	fakeExit := func(code int) {
		assert.Equal(t, 1, code)
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	key := "dummy_key"
	os.Setenv(key, key)
	defer os.Unsetenv(key)
	Config.Load()
	getInt := Config.GetInt(key)
	assert.Equal(t, int64(0), getInt)
}

func TestGetStringShouldReturnString(t *testing.T) {
	key := "dummy_key"
	os.Setenv(key, "8080")
	defer os.Unsetenv(key)
	Config.Load()
	value := Config.GetString(key)
	assert.Equal(t, value, "8080")
}

func TestFatalGetStringShouldFallbackToViperIfNotFound(t *testing.T) {
	key := "DUMMY-KEY"
	os.Setenv(key, "")
	defer os.Unsetenv(key)
	viper.SetDefault(key, "123")
	value := Config.fatalGetString(key)
	assert.Equal(t, "123", value)
}

func TestFatalGetStringShouldExitIfNotFound(t *testing.T) {
	exited := false
	defer func() {
		assert.Equal(t, true, exited)
	}()

	fakeExit := func(code int) {
		exited = true
		assert.Equal(t, 1, code)
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	Config.fatalGetString("anything")
}

func TestCheckKeyShouldExitWithCode1IfNotFound(t *testing.T) {
	exited := false
	defer func() {
		assert.Equal(t, true, exited)
	}()

	fakeExit := func(code int) {
		exited = true
		assert.Equal(t, 1, code)
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	checkKey("DUMMY-ENV")
}

func TestGetStringOrDefaultShouldReturnDefaultIfConfigNotExist(t *testing.T) {
	value := Config.GetStringOrDefault("ANY-KEY", "123")
	assert.Equal(t, "123", value)
}

func TestGetStringOrDefaultShouldReturnConfiguredKeyIfExist(t *testing.T) {
	Config.Add("ANY-KEY", "CONFIGURED-VALUE")
	Config.Load()
	value := Config.GetStringOrDefault("ANY-KEY", "123")
	assert.Equal(t, "CONFIGURED-VALUE", value)
}

func TestGetIntOrDefaultShouldReturnDefaultIfConfigNotExist(t *testing.T) {
	value := Config.GetIntOrDefault("ANY-KEY-2", 123)
	assert.Equal(t, int64(123), value)
}

func TestGetIntOrDefaultShouldReturnConfiguredKeyIfExist(t *testing.T) {
	Config.Add("ANY-KEY", "456")
	Config.Load()
	value := Config.GetIntOrDefault("ANY-KEY", 123)
	assert.Equal(t, int64(456), value)
}

func TestGetIntOrDefaultShouldReturnDefaultIfEmptyString(t *testing.T) {
	Config.Add("ANY-KEY", "")
	Config.Load()
	value := Config.GetIntOrDefault("ANY-KEY", 123)
	assert.Equal(t, int64(123), value)
}

func TestGetIntOrDefaultShouldReturnDefaultIfNotInteger(t *testing.T) {
	Config.Add("ANY-KEY-4", "not-integer")
	Config.Load()
	value := Config.GetIntOrDefault("ANY-KEY-4", 123)
	assert.Equal(t, int64(123), value)
}

func TestGetIntOrDefaultShouldReturnDefaultIfNotBoolean(t *testing.T) {
	Config.Add("ANY-KEY-4", "not-integer")
	Config.Load()
	value := Config.GetIntOrDefault("ANY-KEY-4", 123)
	assert.Equal(t, int64(123), value)
}

func TestConfigFileShouldSetCustomConfigFile(t *testing.T) {
	Config.ConfigFile("/any/path/a.yml")
	assert.Equal(t, configFile, "/any/path/a.yml")
}

func TestConfigGetBooleanOrDefaultShouldReturnDefaultIfNotSet(t *testing.T) {
	value := Config.GetBooleanOrDefault("any-boolean-key", false)
	assert.Equal(t, false, value)
}

func TestConfigGetBooleanOrDefaultShouldReturnDfaultIfConfiguredValueNotBoolean(t *testing.T) {
	Config.Add("any-boolean-key", "not-boolean")
	Config.Load()
	value := Config.GetBooleanOrDefault("any-boolean-key", false)
	assert.Equal(t, false, value)
}

func TestConfigGetBooleanOrDefaultShouldReturnConfigIfSet(t *testing.T) {
	Config.Add("any-boolean-key", "true")
	Config.Load()
	value := Config.GetBooleanOrDefault("any-boolean-key", false)
	assert.Equal(t, true, value)
}
