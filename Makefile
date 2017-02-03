README:
	godoc2md github.com/eduncan911/podcast > README.md.tmp
	echo "[![GoDoc](https://godoc.org/github.com/eduncan911/podcast?status.svg)](https://godoc.org/github.com/eduncan911/podcast) [![Build Status](https://travis-ci.org/eduncan911/podcast.svg?branch=master)](https://travis-ci.org/eduncan911/podcast) [![Coverage Status](https://coveralls.io/repos/github/eduncan911/podcast/badge.svg?branch=master)](https://coveralls.io/github/eduncan911/podcast?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/eduncan911/podcast)](https://goreportcard.com/report/github.com/eduncan911/podcast)" > README.md
	cat README.md.tmp >> README.md
	rm README.md.tmp
