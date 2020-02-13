SHELL = /bin/bash

GITHUB_REPO:=eduncan911/podcast

README:
	godoc2ghmd -play -ex -verify_import_links=0 github.com/$(GITHUB_REPO) > README.md.tmp
	echo "[![GoDoc](https://godoc.org/github.com/$(GITHUB_REPO)?status.svg)](https://godoc.org/github.com/$(GITHUB_REPO))" > README.md	
	echo "[![Build Status](https://github.com/$(GITHUB_REPO)/workflows/go-cicd/badge.svg)](https://github.com/$(GITHUB_REPO)/actions?workflow=go-cicd)" >> README.md
	echo "[![Coverage Status](https://coveralls.io/repos/github/$(GITHUB_REPO)/badge.svg?branch=master)](https://coveralls.io/github/$(GITHUB_REPO)?branch=master)" >> README.md
	echo "[![Go Report Card](https://goreportcard.com/badge/github.com/$(GITHUB_REPO))](https://goreportcard.com/report/github.com/$(GITHUB_REPO))" >> README.md
	echo "[![MIT License](https://img.shields.io/npm/l/mediaelement.svg)](https://eduncan911.mit-license.org/)" >> README.md
	echo  >>README.md
	cat README.md.tmp >> README.md
	rm README.md.tmp

clean:
	rm -rf corpus crashers suppressions workdir podcast-fuzz.zip
