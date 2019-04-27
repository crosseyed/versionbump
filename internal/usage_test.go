package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cmd = "versionbump"

func TestUsage(t *testing.T) {
	usage := `Usage:
    ls [-lart] <file>...
`
	assert.Equal(t, *Usage(&usage), usage)
}

func TestArgs_List(t *testing.T) {
	p := "some/fake/verison.txt"
	cmdList := []string{cmd, "list", p}
	cliops, err := Args(cmdList[1:], nil, `1.1.0`)
	if err != nil {
		assert.Fail(t, "", err)
	}
	assert.False(t, cliops.Major)
	assert.False(t, cliops.Minor)
	assert.False(t, cliops.Patch)
	assert.True(t, cliops.List)

	assert.Equal(t, p, cliops.File)
}

func TestArgs_Major(t *testing.T) {
	p := "some/fake/verison.txt"
	c := []string{cmd, "major", p}
	cliops, err := Args(c[1:], nil, `1.1.0`)
	if err != nil {
		assert.Fail(t, "", err)
	}
	assert.True(t, cliops.Major)
	assert.False(t, cliops.Minor)
	assert.False(t, cliops.Patch)
	assert.False(t, cliops.List)

	assert.Equal(t, p, cliops.File)
}

func TestArgs_Minor(t *testing.T) {
	p := "some/fake/verison.txt"
	c := []string{cmd, "minor", p}
	cliops, err := Args(c[1:], nil, `1.1.0`)
	if err != nil {
		assert.Fail(t, "", err)
	}
	assert.False(t, cliops.Major)
	assert.True(t, cliops.Minor)
	assert.False(t, cliops.Patch)
	assert.False(t, cliops.List)

	assert.Equal(t, p, cliops.File)
}

func TestArgs_Patch(t *testing.T) {
	p := "some/fake/verison.txt"
	c := []string{cmd, "patch", p}
	cliops, err := Args(c[1:], nil, `1.1.0`)
	if err != nil {
		assert.Fail(t, "", err)
	}
	assert.False(t, cliops.Major)
	assert.False(t, cliops.Minor)
	assert.True(t, cliops.Patch)
	assert.False(t, cliops.List)

	assert.Equal(t, p, cliops.File)
}
