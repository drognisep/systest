package main

import (
	"github.com/saylorsolutions/x/cli"
	"github.com/spf13/pflag"
	"os"
	"path/filepath"
)

var (
	configLocation string
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("unable to determine user home directory")
	}
	configLocation = filepath.Join(home, "systest.yaml")
}

func main() {
	tlc := cli.NewCommandSet("systest")
	setupRunCmd(tlc)
	setupRunDirCmd(tlc)

	if tlc.RespondUsage(baseUsage) {
		return
	}
	if err := tlc.Exec(os.Args[1:]); err != nil {
		os.Exit(-1)
	}
}

var baseUsage = `systest is an automated, approachable application for integration testing, using Lua scripts.`

// These setup commands can get pretty long, so it's probably best to separate these out to their own files at some point.

func setupRunCmd(tlc *cli.CommandSet) {
	cmd := tlc.AddCommand("run", "run SCRIPT_FILE", "r")
	cmd.Flags().StringVarP(&configLocation, "config", "C", configLocation, "Sets the location of systest configuration")
	cmd.Does(func(flags *pflag.FlagSet, out *cli.Printer) error {
		// This is where the actual sub-command logic goes.
		out.Println("In 'run' command")
		return nil
	})
	cmd.Usage("SCRIPT_FILE is a single Lua file that should be evaluated as a test script")
}

func setupRunDirCmd(tlc *cli.CommandSet) {
	cmd := tlc.AddCommand("rundir", "rundir DIR", "r")
	cmd.Flags().StringVarP(&configLocation, "config", "C", configLocation, "Sets the location of systest configuration")
	cmd.Flags().BoolP("recursive", "R", false, "Specifies that directories should be searched recursively for Lua test files")
	cmd.Flags().StringP("ext", "x", "test", "Specifies an alternative file extension used to recognize Lua test files.")
	cmd.Does(func(flags *pflag.FlagSet, out *cli.Printer) error {
		// This is where the actual sub-command logic goes.
		out.Println("In 'rundir' command")
		return nil
	})
	cmd.Usage("DIR is a directory containing Lua file that should be evaluated as test scripts")
}
