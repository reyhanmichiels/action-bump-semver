package semver

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	semverFormat    = "v%v.%v.%v%v"
	LevelMajor      = "major"
	LevelMinor      = "minor"
	LevelPatch      = "patch"
	LevelPrerelease = "prerelease"
)

type SemanticVersion struct {
	Major              int
	Minor              int
	Patch              int
	Prefix             string
	Prerelease         int
	isPrerelease       bool
	isUpdatePrerelease bool
}

func Init(semverString string) SemanticVersion {
	var (
		major        int
		minor        int
		patch        int
		prefix       string
		isPrerelease bool
		prerelease   int
		err          error
	)

	semverString = semverString[1:]

	if strings.Contains(semverString, "-") {
		// check if prefix exist
		hyphenIndex := strings.Index(semverString, "-")
		if len(semverString[hyphenIndex+1:]) <= 1 {
			log.Fatalf("current version is not valid")
		}

		lastDotIndex := strings.LastIndex(semverString, ".")
		prefix = semverString[hyphenIndex+1 : lastDotIndex]
		prerelease, err = strconv.Atoi(string(semverString[lastDotIndex+1:]))
		if err != nil {
			log.Fatalf("failed parsing prerelease: %v", err)
			return SemanticVersion{}
		}

		semverString = semverString[:hyphenIndex]
		isPrerelease = true
	}

	semverSlices := strings.Split(semverString, ".")

	major, err = strconv.Atoi(semverSlices[0])
	if err != nil {
		log.Fatalf("failed parsing major: %v", err)
		return SemanticVersion{}
	}

	minor, err = strconv.Atoi(semverSlices[1])
	if err != nil {
		log.Fatalf("failed parsing minor: %v", err)
		return SemanticVersion{}
	}

	patch, err = strconv.Atoi(semverSlices[2])
	if err != nil {
		log.Fatalf("failed parsing patch: %v", err)
		return SemanticVersion{}
	}

	return SemanticVersion{
		Major:        major,
		Minor:        minor,
		Patch:        patch,
		Prefix:       prefix,
		Prerelease:   prerelease,
		isPrerelease: isPrerelease,
	}
}

func (s *SemanticVersion) UpdateMajor() {
	s.Major++
	s.Minor = 0
	s.Patch = 0
	s.Prefix = ""
	s.Prerelease = 0
}

func (s *SemanticVersion) UpdateMinor() {
	s.Minor++
	s.Patch = 0
	s.Prefix = ""
	s.Prerelease = 0
}

func (s *SemanticVersion) UpdatePatch() {
	if !s.isPrerelease {
		s.Patch++
	}
	s.Prefix = ""
	s.Prerelease = 0
}

func (s *SemanticVersion) UpdatePrerelease(prefix string) {
	s.isUpdatePrerelease = true
	s.Prefix = prefix

	if s.isPrerelease {
		s.Prerelease++
		return
	}

	s.Patch++
}

func (s *SemanticVersion) Build() string {
	var prerelease string
	if s.isUpdatePrerelease {
		prerelease = fmt.Sprintf("%v%v%v", "-", s.getPrefix(), s.Prerelease)
	}

	return fmt.Sprintf(semverFormat, s.Major, s.Minor, s.Patch, prerelease)
}

func (s *SemanticVersion) getPrefix() string {
	if s.Prefix == "" {
		return ""
	}

	return fmt.Sprint(s.Prefix + ".")
}
