package versionbump

import (
	"fmt"
	"strconv"
)

// Version section struct
// The version will follow semver.
type Version struct {
	major int
	minor int
	Patch int
}

// Return Version as a String
func (v Version) Version() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.Patch)
}

// Bump patch version and create new version
func (v Version) BumpPatch() Version {
	return Version{
		major: v.major,
		minor: v.minor,
		Patch: v.Patch + 1,
	}
}

// Bump minor version and create new version
func (v Version) BumpMinor() Version {
	return Version{
		major: v.major,
		minor: v.minor + 1,
	}
}

// Bump major version and create new version
func (v Version) BumpMajor() Version {
	return Version{
		major: v.major + 1,
	}
}

// Create version struct from parsed string
func NewVersion(major, minor, patch string, msg *string) Version {
	_major, _ := strconv.Atoi(major)
	_minor, _ := strconv.Atoi(minor)
	_patch, _ := strconv.Atoi(patch)
	return Version{
		major: _major,
		minor: _minor,
		Patch: _patch,
	}
}
