package main

import (
	"log"

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

	err := action.SetOutput("new_version", sv.Build())
	if err != nil {
		log.Fatal("failed set output:", err)
	}
}
