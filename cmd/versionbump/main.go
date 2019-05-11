package main

import (
	"errors"
	"fmt"
	"github.com/crosseyed/versionbump/internal"
	"github.com/crosseyed/versionbump/pkg/versionbump"
	"os"
)

var App = internal.AppHelper{}

func main() {
	// Inject into internal
	internal.App = App

	opts, err := internal.Args(os.Args[1:], nil, Version)
	App.Errors(err, func(err error) {
		// Should never get here
		panic(err)
	})

	exit := run(opts)
	App.Exit(exit)
}

func run(conf *internal.CliOps) int {
	var ex int
	switch {
	case (conf.Major || conf.Minor || conf.Patch):
		ex = runBump(conf)
	case conf.List:
		ex = runList(conf)
	default:
		panic(errors.New("Unhandled option"))
	}
	return ex
}

func runBump(conf *internal.CliOps) int {
	bump, _ := versionbump.NewBumpFile(conf.File)
	var atomBump int
	switch {
	case conf.Major:
		atomBump = versionbump.MAJOR
	case conf.Minor:
		atomBump = versionbump.MINOR
	case conf.Patch:
		atomBump = versionbump.PATCH
	}
	if conf.Checktags {
		hasTags(bump.Version())
	}
	_ = bump.BumpVersion(atomBump)
	return 0
}

func hasTags(ver string) {
	path, err := os.Getwd()
	App.Errors(err, errPanic)
	git := internal.PlainOpen(path)
	for _, t := range git.ListTags() {
		if t == ver {
			return
		}
		if t == "v"+ver {
			return
		}
	}
	_, err = fmt.Fprintf(os.Stderr, "Git tag %s not found. Bump tag of \"%s\" aborted\n", ver, path) // nolint
	App.Errors(err, errPanic)
	App.Exit(-1)
}

func runList(conf *internal.CliOps) int {
	bump, _ := versionbump.NewBumpFile(conf.File)
	fmt.Println(bump.Version())
	return 0
}

// Error handlers
func errPanic(err error) {
	panic(err)
}
