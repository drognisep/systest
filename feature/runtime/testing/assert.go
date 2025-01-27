package testing

import (
	"github.com/drognisep/systest/feature/runtime/errors"
	lua "github.com/yuin/gopher-lua"
)

func assertEqual(l *lua.LState) int {
	first := l.Get(1)
	second := l.Get(2)
	if first.Type() != second.Type() {
		errors.LError(l, "values are different types")
		return 0
	}
	switch first.Type() {
	case lua.LTNil:
		return 0
	case lua.LTString:
		fallthrough
	case lua.LTNumber:
		fallthrough
	case lua.LTBool:
		if first.String() != second.String() {
			errors.LErrorf(l, "value '%s' does not equal value '%s'", first, second)
		}
		return 0
	default:
		errors.LErrorf(l, "invalid argument types for comparison: '%s'", first.Type())
		return 0
	}
}

func assertTrue(l *lua.LState) int {
	bval, ok := l.Get(1).(lua.LBool)
	if !ok {
		errors.LError(l, "value is not boolean")
		return 0
	}
	if !bval {
		errors.LError(l, "value is not true")
	}
	return 0
}
