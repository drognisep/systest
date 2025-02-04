package errors

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

func LError(l *lua.LState, msg string) {
	l.Error(lua.LString(msg), 0)
}

func LErrorf(l *lua.LState, msg string, args ...any) {
	l.Error(lua.LString(fmt.Sprintf(msg, args...)), 0)
}
