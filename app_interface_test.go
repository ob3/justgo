package justgo

import (
	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
	"testing"
)

func TestRunAppInterfaces(t *testing.T) {
	appInterfaces = []*AppInterface{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	waitForShutDown := false
	dummyWaitForShutDown := func() {waitForShutDown = true}
	monkey.Patch(WaitForShutdown, dummyWaitForShutDown)


	mockInterface := NewMockAppInterface(ctrl)
	RegisterInterface(mockInterface)

	RunAppInterfaces()
	mockInterface.EXPECT().Serve()

	defer func() {
		assert.Equal(t, waitForShutDown, true)
	}()
}

func TestRegisterInterface(t *testing.T) {
	appInterfaces = []*AppInterface{}
	assert.Equal(t, len(appInterfaces), 0)
	RegisterInterface(NewMockAppInterface(nil))
	assert.Equal(t, len(appInterfaces), 1)
	RegisterInterface(NewMockAppInterface(nil))
	assert.Equal(t, len(appInterfaces), 2)

}