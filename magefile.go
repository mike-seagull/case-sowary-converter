//go:build mage
// +build mage

package main

import (
	// "encoding/binary"
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var BIN_DIR string = "bin"
var BIN_NAME string = "csc"
var MAIN_DIR string = "cmd"
var MAIN_FULLPATH string = fmt.Sprintf("./%s/%s", MAIN_DIR, BIN_NAME)

var Default = All // Default target

type BuildEnv struct {
	goos        string
	goarch      string
	binary_name string
}

var MAC = BuildEnv{"darwin", "amd64", fmt.Sprintf("./%s/%s-mac", BIN_DIR, BIN_NAME)}
var LINUX = BuildEnv{"linux", "amd64", fmt.Sprintf("./%s/%s", BIN_DIR, BIN_NAME)}
var WINDOWS = BuildEnv{"windows", "amd64", fmt.Sprintf("./%s/%s-win.exe", BIN_DIR, BIN_NAME)}
var ENVS = []BuildEnv{MAC, LINUX, WINDOWS}

func Build() {
	for _, env := range ENVS {
		env_vars := map[string]string{
			"GOOS":   env.goos,
			"GOARCH": env.goarch,
		}
		sh.RunWith(env_vars, "go", "build", "-o", env.binary_name, MAIN_FULLPATH)
	}
}
func Test() { // NOTE: output is only displayed if mage is run with "-v" option
	sh.Run("go", "test", "-v", "pkg/case_converter")
}
func Clean() {
	sh.Run("go", "clean")
	for _, env := range ENVS {
		sh.Rm(env.binary_name)
	}
}
func All() {
	mg.SerialDeps(Clean, Test, Build)
}
