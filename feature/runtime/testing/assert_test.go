package testing

import (
	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
	"testing"
)

func TestAssertions(t *testing.T) {
	l := lua.NewState()
	state := new(State)
	l.PreloadModule(ModuleName, Loader(state))
	assert.NoError(t, l.DoFile("assert_test.lua"))
	assert.Equal(t, 4, state.PassedCount)
	assert.Equal(t, 2, state.FailedCount)
}
