package internal

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"log"
	"strings"
)

type Git struct {
	repo *git.Repository
}

func PlainOpen(path string) Git {
	repo, err := git.PlainOpen(path)
	App.Errors(err, errGitLogStderr_Exit)
	return Git{repo: repo}
}

func (r *Git) ListTags() []string {
	tl := []string{}
	tagrefs, err := r.repo.Tags()
	App.Errors(err, errGitLogStderr_Exit)
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		refname := string(t.Name())
		tag := strings.ReplaceAll(refname, "refs/tags/", "")

		tl = append(tl, tag)
		return nil
	})
	App.Errors(err, errGitLogStderr_Exit)
	return tl
}

// Error Handlers
func errGitLogStderr_Exit(err error) {
	log.Printf(`git: %v`, err)
	App.Exit(-1)
}
