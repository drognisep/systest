package testing

import (
	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
	"testing"
)

func TestRun(t *testing.T) {
	l := lua.NewState()
	defer l.Close()
	tstate := new(State)
	l.PreloadModule(ModuleName, Loader(tstate))
	assert.NoError(t, l.DoFile("run_test.lua"))
	assert.Equal(t, 1, tstate.PassedCount)
	assert.Equal(t, 1, tstate.FailedCount)
}
