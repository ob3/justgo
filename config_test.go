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
