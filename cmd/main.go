package main

import (
	"fmt"
	"os"

	"github.com/reyhanmichiels/action-bump-server/semver"
)

func main() {
	currentVersion := os.Getenv("current_version")
	level := os.Getenv("level")
	prefix := os.Getenv("preid")

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
