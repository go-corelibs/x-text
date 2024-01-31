# Go Text

[![Go Reference](https://pkg.go.dev/badge/github.com/go-corelibs/x-text.svg)](https://pkg.go.dev/github.com/go-corelibs/x-text)

This repository holds supplementary Go libraries for text processing, many involving Unicode.

## Go-CoreLibs Fork

The purpose of this fork is to enable multilingual support within the overall
Go-Enjin and Go-Curses projects and due to various reasons, this has proven to
require specific changes to the upstream project that may or may not ever
propagate back.

### Maintained Changes

| #  | commit message |
|----|----------------|
| 0  | `cmd/gotext,message/pipeline: ssoroka - go modules support, see: https://github.com/golang/text/commit/059a5b1dfb6e3c931765fdc5ad8a49299f7bf55b` |
| 1  | `*.{go,json}: replace instances of golang.org/x/text with github.com/go-corelibs/x-text` |
| 2  | `language: implement MarshalBinary and UnmarshalBinary methods for Tag structures` |
| 3  | `language: added func Compare(a Tag, others ...Tag) (equal bool)` |
| 4  | `cmd/gotext: document rewrite sub-command -w option` |
| 5  | `cmd/gotext: add IMPORTANT notices to docstrings and --help output` |
| 6  | `message/catalog: implement Include catalog Option and Builder support` |
| 7  | `message/pipeline: do not clear message .Key during State.Merge process` |
| 8  | `cmd/gotext,message/pipeline: document -srclang global option; implement -declare-var and -go-build global options` |
| 9  | `.gitignore: add /gotext` |
| 10 | `README.md: Go-CoreLibs fork updates` |
| 11 | `message: added http.Request tag and printer functions` |

### Branching Structure

This repository has two main branches: `upstream` and `trunk`. All Go-CoreLibs
changes are are made to `trunk`. All changes on their way to upstream are in
their own branches (based off of `upstream` and not `trunk`).

All changes made by the Go-Enjin team are licensed under the same terms as
the upstream project.

### Tagging Versions

This repository follows the same pattern as the upstream project with the
exception that all upstream tags are included here suffixed with `-upstream`,
making them beta releases. The actual tagged releases include the Go-Enjin
changes on top of the upstream versions. This is done to minimize confusion
and simplify the build systems used by Go-Enjin.

For example:

- upstream released `v0.14.0`
- `trunk` is essentially hard-reset to the upstream changes
- `trunk` is tagged `v0.14.0-upstream`
- all Go-Enjin changes are re-applied to `trunk`
- `trunk` is tagged `v0.14.0`

Currently there is no evidence of upstream doing any patch-level versions so
this model works without much effort. Should upstream decide to start using
patch-level versions, this model still works because the beta releases are
always upstream and the official releases are always the latest upstream
plus the Go-Enjin changes.

### Long-Term Support

This package is now a part of the Go-CoreLibs project and will be maintained
for as long as Go-Enjin and Go-Curses projects need it.

## CLDR Versioning

It is important that the Unicode version used in `x/text` matches the one used
by your Go compiler. The `x/text` repository supports multiple versions of
Unicode and will match the version of Unicode to that of the Go compiler. At the
moment this is supported for Go compilers from version 1.7.

## Download/Install

The easiest way to install is to run `go get -u github.com/go-corelibs/x-text`. You can
also manually git clone the repository to `$GOPATH/src/github.com/go-corelibs/x-text`.

## Contribute
To submit changes to this repository, see http://golang.org/doc/contribute.html.

To generate the tables in this repository (except for the encoding tables),
run go generate from this directory. By default tables are generated for the
Unicode version in core and the CLDR version defined in
github.com/go-corelibs/x-text/unicode/cldr.

Running go generate will as a side effect create a DATA subdirectory in this
directory, which holds all files that are used as a source for generating the
tables. This directory will also serve as a cache.

## Testing
Run

    go test ./...

from this directory to run all tests. Add the "-tags icu" flag to also run
ICU conformance tests (if available). This requires that you have the correct
ICU version installed on your system.

TODO:
- updating unversioned source files.

## Generating Tables

To generate the tables in this repository (except for the encoding
tables), run `go generate` from this directory. By default tables are
generated for the Unicode version in core and the CLDR version defined in
github.com/go-corelibs/x-text/unicode/cldr.

Running go generate will as a side effect create a DATA subdirectory in this
directory which holds all files that are used as a source for generating the
tables. This directory will also serve as a cache.

## Versions
To update a Unicode version run

    UNICODE_VERSION=x.x.x go generate

where `x.x.x` must correspond to a directory in https://www.unicode.org/Public/.
If this version is newer than the version in core it will also update the
relevant packages there. The idna package in x/net will always be updated.

To update a CLDR version run

    CLDR_VERSION=version go generate

where `version` must correspond to a directory in
https://www.unicode.org/Public/cldr/.

Note that the code gets adapted over time to changes in the data and that
backwards compatibility is not maintained.
So updating to a different version may not work.

The files in DATA/{iana|icu|w3|whatwg} are currently not versioned.

## Report Issues / Send Patches

This repository uses Gerrit for code changes. To learn how to submit changes to
this repository, see https://golang.org/doc/contribute.html.

The main issue tracker for the image repository is located at
https://github.com/golang/go/issues. Prefix your issue with "x/text:" in the
subject line, so it is easy to find.
