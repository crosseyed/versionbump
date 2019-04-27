package internal

import (
	"github.com/docopt/docopt-go"
)

type CliOps struct {
	File      string
	Checktags bool
	Major     bool
	Minor     bool
	Patch     bool
	List      bool
}

func UsageDefault() *string {
	usage := `
Usage:
    versionbump (major|minor|patch) [--checktags] <file>
    versionbump list <file>

Options:
    -h --help       Show this screen.
    --version       Show version.
    -c --checktags  Avoid bumping if current version is not tagged.
    <file>          Name of file.
    major           major bump.
    minor           minor bump.
    patch           patch bump.
    list            Show version from file.
`
	return &usage
}

func Usage(doc *string) *string {
	if doc != nil {
		return doc
	}
	usage := UsageDefault()
	return usage
}

func Args(args []string, doc *string, version string) (*CliOps, error) {
	opts, err := docopt.ParseArgs(*Usage(doc), args, version) // nolint
	var conf = &CliOps{}
	err = opts.Bind(conf)
	App.Errors(err, errUsagePanic)
	return conf, nil
}

// Error Handler(s)
func errUsagePanic(err error) {
	panic(err)
}
