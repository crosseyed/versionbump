package versionbump

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"
)

// Setup
func root(t *testing.T) string {
	p, err := filepath.Abs("../../")
	if err != nil {
		assert.Fail(t, "Error getting absolute root path", err)
	}
	return p
}

func TestNewBumpFileExists(t *testing.T) {
	p := path.Join(root(t), "tmp/testfiles/version.txt")

	_, err := NewBumpFile(p)
	assert.Nil(t, err)
}

func TestNewBumpFileDoesNotExists(t *testing.T) {
	p := path.Join(root(t), "non_such_file.txt")

	_, err := NewBumpFile(p)
	assert.Error(t, err)
}

func TestNewBumpFileIsDir(t *testing.T) {
	p := path.Join(root(t), "tmp/")

	_, err := NewBumpFile(p)
	assert.Error(t, err)
}

func TestNewBumpFileNoPerissions(t *testing.T) {
	p := path.Join(root(t), "tmp/testfiles/version_no_access.txt")

	_, err := NewBumpFile(p)
	assert.Error(t, err)
}

func TestBumpFile_Version(t *testing.T) {
	inversion := "1.3.4"
	input := fmt.Sprintf("Project: awesomproject\nVersion: %s\n", inversion)
	p := path.Join(root(t), "tmp/testfiles/TestBumpFile_Version.txt")

	f, err := os.Create(p)
	if err != nil {
		assert.Fail(t, "", err)
	}
	if _, err := f.Write([]byte(input)); err != nil {
		assert.Fail(t, "", err)
	}
	if err := f.Close(); err != nil {
		assert.Fail(t, "", err)
	}

	output, err := ioutil.ReadFile(p)
	if err != nil {
		assert.Fail(t, "", err)
	}

	bump, _ := NewBumpFile(p)
	ver := bump.Version()

	assert.Contains(t, string(output), ver)
}

func TestBumpFile_BumpVersion_major(t *testing.T) {
	inversion := "1.3.4"
	bumpversion := "2.0.0"
	input := fmt.Sprintf(`Name:           bello
Version:        %s
Release:        1%%{?dist}
Summary:        Hello World example implemented in bash script`, inversion)

	p := path.Join(root(t), "tmp/testfiles/TestBumpFile_BumpVersion_major.txt")
	f, err := os.Create(p)
	if err != nil {
		assert.Fail(t, "", err)
	}
	if _, err := f.Write([]byte(input)); err != nil {
		assert.Fail(t, "", err)
	}
	if err := f.Close(); err != nil {
		assert.Fail(t, "", err)
	}

	bump, err := NewBumpFile(p)
	if err != nil {
		assert.Fail(t, "", err)
	}
	if err := bump.BumpVersion(MAJOR); err != nil {
		assert.Fail(t, "", err)
	}
	ver := bump.Version()

	assert.Equal(t, ver, bumpversion)
}

func TestBumpFile_BumpVersion_minor(t *testing.T) {
	inversion := "2.3.4"
	bumpversion := "2.4.0"
	input := fmt.Sprintf(`{
  "project": "42"
  "version": "%s"
}`, inversion)

	p := path.Join(root(t), "tmp/testfiles/TestBumpFile_BumpVersion_minor.txt")
	f, err := os.Create(p)
	if err != nil {
		assert.Fail(t, "", err)
	}
	if _, err := f.Write([]byte(input)); err != nil {
		assert.Fail(t, "", err)
	}
	if err := f.Close(); err != nil {
		assert.Fail(t, "", err)
	}

	bump, err := NewBumpFile(p)
	if err != nil {
		assert.Fail(t, "", err)
	}
	if err := bump.BumpVersion(MINOR); err != nil {
		assert.Fail(t, "", err)
	}
	ver := bump.Version()

	assert.Equal(t, ver, bumpversion)
}

func TestBumpFile_BumpVersion_patch(t *testing.T) {
	inversion := "10.11.4"
	bumpversion := "10.11.5"
	input := fmt.Sprintf(`
_
  project: life
  version: %s
}`, inversion)

	p := path.Join(root(t), "tmp/testfiles/TestBumpFile_BumpVersion_patch.txt")
	f, err := os.Create(p)
	if err != nil {
		assert.Fail(t, "", err)
	}
	if _, err := f.Write([]byte(input)); err != nil {
		assert.Fail(t, "", err)
	}
	if err := f.Close(); err != nil {
		assert.Fail(t, "", err)
	}

	bump, err := NewBumpFile(p)
	if err != nil {
		assert.Fail(t, "", err)
	}
	if err := bump.BumpVersion(PATCH); err != nil {
		assert.Fail(t, "", err)
	}
	ver := bump.Version()

	assert.Equal(t, ver, bumpversion)
}

func TestBumpFile_uniformVersions(t *testing.T) {
	inversion := "1.1.1"
	input := fmt.Sprintf(`
- %s
- %s
- %s
- %s
`, inversion, inversion, inversion, inversion)

	p := path.Join(root(t), "tmp/testfiles/TestBumpFile_uniformVersions.txt")
	f, err := os.Create(p)
	if err != nil {
		assert.Fail(t, "", err)
	}
	if _, err := f.Write([]byte(input)); err != nil {
		assert.Fail(t, "", err)
	}
	if err := f.Close(); err != nil {
		assert.Fail(t, "", err)
	}
	if err != nil {
		assert.Fail(t, "", err)
	}

	_, err = NewBumpFile(p)
	if err != nil {
		assert.Fail(t, "", err)
	}
}

func TestBumpFile_uniformVersions_not(t *testing.T) {
	input := fmt.Sprint(`- 1.0.1
- 1.1.1
- 2.0.0
- 3.1.0
`)

	p := path.Join(root(t), "tmp/testfiles/TestBumpFile_uniformVersions.txt")
	f, err := os.Create(p)
	if err != nil {
		assert.Fail(t, "", err)
	}
	if _, err := f.Write([]byte(input)); err != nil {
		assert.Fail(t, "", err)
	}
	if err := f.Close(); err != nil {
		assert.Fail(t, "", err)
	}
	if err != nil {
		assert.Fail(t, "", err)
	}

	_, err = NewBumpFile(p)
	assert.Error(t, err)
}
