// Code to automatically bump versions
package versionbump

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	_ "path/filepath"
	"regexp"
	"strings"
)

const (
	MAJOR = iota // MAJOR bump flag
	MINOR        // MINOR bump flag
	PATCH        // PATCH bump flag
)

var regex = regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`) // nolint

//
// Interfaces
//

type Bump interface {
	VersionRegex(regex *string) error
	Version() string
	BumpVersion(part int) error
}

//
// Structs
//

// Bump Version contained in a file. By default
// Do not instantiate using BumpFile instead use NewBumpFile()
type BumpFile struct {
	Bump
	path     string
	contents []byte
	curver   *Version
}

// Opens a new bump file for reading. Note that the Bump interface is returned not BumpFile
func NewBumpFile(path string) (Bump, error) {
	bf := BumpFile{path: path}
	if err := bf.parse(); err != nil {
		return nil, err
	}

	return &bf, nil
}

// Get the current version
func (b *BumpFile) Version() string {
	return b.curver.Version()
}

// Atomically writes the version to file. msg *string is not currently implimented
//
// MAJOR version bump:
//   f.BumpVersion(internal.MAJOR)
//   f.BumpVersion(0)
//
// MINOR version bump:
//   f.BumpVersion(internal.MINOR)
//   f.BumpVersion(1)
//
// PATCH version bump:
//   f.BumpVersion(internal.PATCH)
//   f.BumpVersion(2)
//
func (b *BumpFile) BumpVersion(part int) error {
	curver, newver := b.bump(part)
	oldContents := string(b.contents)
	newContents := strings.Replace(oldContents, curver, newver, -1)
	tmpfile, err := ioutil.TempFile("", "bump-*")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = os.Remove(tmpfile.Name())
	}()

	if _, err := tmpfile.Write([]byte(newContents)); err != nil {
		panic(err)
	}
	if err := tmpfile.Close(); err != nil {
		panic(err)
	}
	src := tmpfile.Name()
	dst := b.path
	if err := os.Rename(src, dst); err != nil {
		panic(err)
	}
	if err := b.parse(); err != nil {
		return err
	}
	return nil
}

func (b *BumpFile) parse() error {
	if err := b.parseFile(); err != nil {
		return err
	}
	if err := b.parseVersion(); err != nil {
		return err
	}
	return nil
}

func (b *BumpFile) parseFile() error {
	path := b.path
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fi.IsDir() {
		return fmt.Errorf("b %s is a directory", path)
	}

	b.contents, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return nil
}

// Parse current version. Returns the count of number of versions found
func (b *BumpFile) parseVersion() error {
	cnt := 0
	matches := regex.FindAllString(string(b.contents), -1)
	verstrs, uniform := uniformVersions(matches)
	cnt = len(verstrs)
	switch {
	case cnt > 1:
		if !uniform {
			return fmt.Errorf(`Expected one version found multiple in %s: %v`, b.path, verstrs)
		}
	case cnt == 0:
		fmt.Printf("No version found in %s", b.path)
		return nil
	}

	verAtoms := strings.Split(verstrs[0], ".")
	verObj := NewVersion(verAtoms[0], verAtoms[1], verAtoms[2], nil)
	b.curver = &verObj

	return nil
}

func (b *BumpFile) bump(part int) (string, string) {
	curverStr := b.curver.Version()
	var newver Version
	switch part {
	case MAJOR:
		newver = b.curver.BumpMajor()
	case MINOR:
		newver = b.curver.BumpMinor()
	case PATCH:
		newver = b.curver.BumpPatch()
	}
	newverStr := newver.Version()
	return curverStr, newverStr
}

//
// Functions
//

func uniformVersions(vers []string) ([]string, bool) {
	uniform := true
	first := vers[0]
	verstrs := []string{first}
	for _, ver := range vers {
		if first != ver {
			uniform = false
			verstrs = append(verstrs, string(ver))
		}
	}
	return verstrs, uniform
}