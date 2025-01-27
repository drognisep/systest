package runtime

import (
	"github.com/drognisep/systest/feature/runtime/testing"
	lua "github.com/yuin/gopher-lua"
)

// State represents the state of the runtime.
// This will be used to provide configuration and secrets at a later point.
type State struct {
	testState *testing.State
	luaState  *lua.LState
}

func New() *State {
	state := new(testing.State)
	return &State{
		testState: state,
	}
}

func (s *State) PassingTests() int {
	if s.testState == nil {
		return 0
	}
	return s.testState.PassedCount
}

func (s *State) FailingTests() int {
	if s.testState == nil {
		return 0
	}
	return s.testState.FailedCount
}

func (s *State) TotalTests() int {
	return s.PassingTests() + s.FailingTests()
}
