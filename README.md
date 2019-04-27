# versionbump

versionbump is a tool to automatically bump version contained in a file.

## Usage

```bash
% versionbump --help
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

```

## Examples

**Minor Version Bump**
```bash
% cat ./VERSION
1.0.0
% versionbump minor ./VERSION 
% cat ./VERSION
1.1.0
```

**Version Embeded in a Makefile**
```
# Makefile
PROJECT = someproject
COMMITSHA = $(shell git rev-parse --short HEAD)
VERSION = 2.1.0 

release:
    git tag $(VERSION)
    git push --tags
```

Update Version
```bash
versionbump list ./Makefile
2.1.0
versionbump major ./Makefile
versionbump list ./Makefile
3.0.0
```

```
# Makefile
PROJECT = someproject
COMMITSHA = $(shell git rev-parse --short HEAD)
VERSION = 3.0.0 

release:
    git tag $(VERSION)
    git push --tags
```
