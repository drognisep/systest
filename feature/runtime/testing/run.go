package testing

import (
	"fmt"
	"github.com/drognisep/systest/feature/runtime/errors"
	lua "github.com/yuin/gopher-lua"
)

func run(tstate *State) func(l *lua.LState) int {
	return func(l *lua.LState) int {
		testName := l.ToString(1)
		if len(testName) == 0 {
			errors.LError(l, "empty test name")
			return 0
		}
		fn := l.ToFunction(2)
		p := lua.P{
			Fn:      fn,
			NRet:    0,
			Protect: true,
		}
		if err := l.CallByParam(p); err != nil {
			fmt.Printf("[FAIL] %s: %s\n", testName, err.Error())
			tstate.AddFailed(testName)
		} else {
			fmt.Printf("[PASS] %s\n", testName)
			tstate.AddPassed(testName)
		}
		return 0
	}
}
