package main

import (
	. "github.com/saylorsolutions/modmake"
)

const (
	version = "0.1.0"
)

func main() {
	b := NewBuild()

	systest := NewAppBuild("systest", "./cmd/systest", version)
	systest.Variant("windows", "amd64")
	systest.Variant("windows", "arm64")
	systest.Variant("linux", "amd64")
	systest.Variant("linux", "arm64")
	systest.Variant("darwin", "amd64")
	systest.Variant("darwin", "arm64")
	b.ImportApp(systest)

	b.Execute()
}
