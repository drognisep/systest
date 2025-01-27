package testing

import lua "github.com/yuin/gopher-lua"

const (
	ModuleName = "testing"
)

type State struct {
	Passed      []string
	PassedCount int
	Failed      []string
	FailedCount int
}

func (s *State) AddPassed(name string) {
	s.PassedCount++
	s.Passed = append(s.Passed, name)
}

func (s *State) AddFailed(name string) {
	s.FailedCount++
	s.Failed = append(s.Failed, name)
}

// Loader is a loader function for including testing module functions.
func Loader(testState *State) func(l *lua.LState) int {
	return func(l *lua.LState) int {
		mod := l.SetFuncs(l.NewTable(), map[string]lua.LGFunction{
			"assertEqual": assertEqual,
			"assertTrue":  assertTrue,
			"run":         run(testState),
		})
		l.Push(mod)
		return 1
	}
}
