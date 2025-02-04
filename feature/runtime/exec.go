package runtime

import (
	"fmt"
	"github.com/drognisep/systest/feature/runtime/testing"
	lua "github.com/yuin/gopher-lua"
	"io/fs"
	"path/filepath"
)

func (s *State) newLuaState() *lua.LState {
	// Lua state initialization should be controlled from here.
	// https://github.com/yuin/gopher-lua?tab=readme-ov-file#opening-a-subset-of-builtin-modules
	l := lua.NewState(lua.Options{
		SkipOpenLibs: true,
	})
	// Base lib loading is done here.
	for _, pair := range []struct {
		n string
		f lua.LGFunction
	}{
		{lua.LoadLibName, lua.OpenPackage},
		{lua.BaseLibName, lua.OpenBase},
		{lua.TabLibName, lua.OpenTable},
	} {
		if err := l.CallByParam(lua.P{
			Fn:      l.NewFunction(pair.f),
			NRet:    0,
			Protect: true,
		}, lua.LString(pair.n)); err != nil {
			panic(err)
		}
	}
	// Extension modules are added here.
	l.PreloadModule(testing.ModuleName, testing.Loader(s.testState))
	return l
}

// RunFile runs a single file with a new [lua.LState].
func (s *State) RunFile(name string) error {
	l := s.newLuaState()
	defer l.Close()
	return l.DoFile(name)
}

func (s *State) RunDir(path string) error {
	return filepath.WalkDir(path, func(_path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(_path) == ".lua" {
			if err := s.RunFile(_path); err != nil {
				return fmt.Errorf("failed to run file '%s': %w", _path, err)
			}
		}
		return nil
	})
}
