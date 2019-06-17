package justgo

import (
	"bou.ke/monkey"
	"github.com/spf13/viper"
	"gotest.tools/assert"
	"os"
	"testing"
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
		assert.Equal(t, 1, code )
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
	viper.SetDefault(key, "123")
	value := Config.fatalGetString(key)
	assert.Equal(t, "123", value)
}

func TestCheckKeyShouldExitWithCode1IfNotFound(t *testing.T) {
	fakeExit := func(code int) {
		assert.Equal(t, 1, code )
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	checkKey("DUMMY-ENV")
}


