SHELL = /bin/bash

GITHUB_REPO := "eduncan911/podcast"

README:
	godoc2ghmd github.com/$(GITHUB_REPO) > README.md.tmp
	echo "[![GoDoc](https://godoc.org/github.com/$(GITHUB_REPO)?status.svg)](https://godoc.org/github.com/$(GITHUB_REPO))" > README.md
	echo "[![Build Status](https://travis-ci.org/$(GITHUB_REPO).svg?branch=master)](https://travis-ci.org/$(GITHUB_REPO))" >> README.md
	echo "[![Coverage Status](https://coveralls.io/repos/github/$(GITHUB_REPO)/badge.svg?branch=master)](https://coveralls.io/github/$(GITHUB_REPO)?branch=master)" >> README.md
	echo "[![Go Report Card](https://goreportcard.com/badge/github.com/$(GITHUB_REPO))](https://goreportcard.com/report/github.com/$(GITHUB_REPO))" >> README.md
	echo "[![GoDoc](https://godoc.org/github.com/$(GITHUB_REPO)?status.svg)](https://godoc.org/github.com/$(GITHUB_REPO))"
	echo "[![MIT License](https://img.shields.io/npm/l/mediaelement.svg)](https://eduncan911.mit-license.org/)" >> README.md
	echo  >>README.md
	cat README.md.tmp >> README.md
	rm README.md.tmp
