package mods

import (
	"fmt"
)

var (
	versionString   = ""
	versionGitSHA   = ""
	buildTimestamp  = ""
	goVersionString = ""
)

func VersionDescription() string {
	return fmt.Sprintf("%s (%v, %v, go %v)", versionString, versionGitSHA, buildTimestamp, goVersionString)
}

func Version() string {
	return versionString
}

func GitSHA() string {
	return versionGitSHA
}
