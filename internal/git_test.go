package internal

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func root() string {
	abs, _ := filepath.Abs("..")
	return abs
}

func TestGit_ListTags(t *testing.T) {
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd) // nolint

	tmpdir := filepath.Join(root(), "tmp")
	gitsh := filepath.Join(root(), "test/fixtures/gittag.sh")
	d := "gittestdir"
	g := filepath.Join(tmpdir, d)

	cmd := exec.Command(gitsh, d, tmpdir)
	if err := cmd.Run(); err != nil {
		assert.Fail(t, "", err)
	}
	bump := PlainOpen(g)

	tags := bump.ListTags()
	assert.True(t, len(tags) > 0)
	assert.Contains(t, tags, "1.0.0")
	assert.Contains(t, tags, "1.1.0")
}

func TestGit_ListTags_no_git(t *testing.T) {
}

func TestGit_ListTags_no_repo(t *testing.T) {
}
