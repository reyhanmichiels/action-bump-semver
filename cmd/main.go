package main

import (
	"fmt"
	"os"

	"github.com/reyhanmichiels/action-bump-server/semver"
)

func main() {
	currentVersion := os.Getenv("current_version")
	fmt.Println(fmt.Sprintf("::debug::{%v %v}", "current_version:", currentVersion))
	level := os.Getenv("level")
	fmt.Println(fmt.Sprintf("::debug::{%v %v}", "level:", level))
	prefix := os.Getenv("preid")
	fmt.Println(fmt.Sprintf("::debug::{%v %v}", "preid:", prefix))

	sv := semver.Init(currentVersion)

	switch level {
	case semver.LevelPrerelease:
		sv.UpdatePrerelease(prefix)
	case semver.LevelPatch:
		sv.UpdatePatch()
	case semver.LevelMinor:
		sv.UpdateMinor()
	case semver.LevelMajor:
		sv.UpdateMajor()
	}

	fmt.Println(fmt.Sprintf(`::set-output name=new_version::%s`, sv.Build()))
}
