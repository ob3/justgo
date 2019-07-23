package justgo

import (
	"fmt"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/magiconair/properties/assert"
)

func TestStart(t *testing.T) {

	configured, logged, runInterface, defaultInterface, registerInterface := false, false, false, false, false

	dummyConfigLoad := func(*config) {
		configured = true
	}
	dummyLogLoad := func(*justGoLog) {
		logged = true
	}
	dummyRunAppInterface := func() {
		runInterface = true
	}
	dummyGetDefaultInterface := func() *HttpInterface {
		defaultInterface = true
		return nil
	}
	dummyRegisterInterface := func(AppInterface) {
		registerInterface = true
		fmt.Println(registerInterface)
	}

	pConfig := monkey.PatchInstanceMethod(reflect.TypeOf(Config), "Load", dummyConfigLoad)
	pLog := monkey.PatchInstanceMethod(reflect.TypeOf(log), "Load", dummyLogLoad)
	pRunInterface := monkey.Patch(RunAppInterfaces, dummyRunAppInterface)
	pGetDefaultHttpInterface := monkey.Patch(GetDefaultHttpInterface, dummyGetDefaultInterface)
	pRegisterInterface := monkey.Patch(RegisterInterface, dummyRegisterInterface)
	defer pConfig.Unpatch()
	defer pLog.Unpatch()
	defer pRunInterface.Unpatch()
	defer pGetDefaultHttpInterface.Unpatch()
	defer pRegisterInterface.Unpatch()

	Start()

	defer func() {
		assert.Equal(t, configured, true)
		assert.Equal(t, logged, true)
		assert.Equal(t, runInterface, true)
		assert.Equal(t, defaultInterface, true)
		//assert.Equal(t, registerInterface, true)
		fmt.Println(registerInterface)
	}()
}

func TestStartShouldNotUseDefaultInterfaceIfNotEmpty(t *testing.T) {

	configured, logged, runInterface, defaultInterface, registerInterface := false, false, false, false, false

	httpInterface := GetDefaultHttpInterface()
	RegisterInterface(httpInterface)

	dummyConfigLoad := func(*config) { configured = true }
	dummyLogLoad := func(*justGoLog) { logged = true }
	dummyRunAppInterface := func() { runInterface = true }
	dummyGetDefaultInterface := func() *HttpInterface {
		defaultInterface = true
		return nil
	}
	dummyRegisterInterface := func(AppInterface) { registerInterface = true }

	pConfig := monkey.PatchInstanceMethod(reflect.TypeOf(Config), "Load", dummyConfigLoad)
	pLog := monkey.PatchInstanceMethod(reflect.TypeOf(log), "Load", dummyLogLoad)
	pRunInterface := monkey.Patch(RunAppInterfaces, dummyRunAppInterface)
	pGetDefaultHttpInterface := monkey.Patch(GetDefaultHttpInterface, dummyGetDefaultInterface)
	pRegisterInterface := monkey.Patch(RegisterInterface, dummyRegisterInterface)
	defer pConfig.Unpatch()
	defer pLog.Unpatch()
	defer pRunInterface.Unpatch()
	defer pGetDefaultHttpInterface.Unpatch()
	defer pRegisterInterface.Unpatch()

	Start()

	defer func() {
		assert.Equal(t, configured, true)
		assert.Equal(t, logged, true)
		assert.Equal(t, runInterface, true)
		assert.Equal(t, defaultInterface, false)
		assert.Equal(t, registerInterface, false)

	}()
}
