package main

import (
	"fmt"

	"github.com/reyhanmichiels/action-bump-server/action"
	"github.com/reyhanmichiels/action-bump-server/semver"
)

func main() {
	currentVersion := action.GetInput("current_version")
	level := action.GetInput("level")
	prefix := action.GetInput("preid")

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
