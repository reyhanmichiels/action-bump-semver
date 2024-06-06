package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/reyhanmichiels/action-bump-server/semver"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err load env: %v", err)
	}

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
