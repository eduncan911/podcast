README:
	godoc2ghmd github.com/eduncan911/podcast > README.md.tmp
	echo "[![GoDoc](https://godoc.org/github.com/eduncan911/podcast?status.svg)](https://godoc.org/github.com/eduncan911/podcast)" > README.md
	echo "[![Build Status](https://travis-ci.org/eduncan911/podcast.svg?branch=master)](https://travis-ci.org/eduncan911/podcast)" >> README.md
	echo "[![Coverage Status](https://coveralls.io/repos/github/eduncan911/podcast/badge.svg?branch=master)](https://coveralls.io/github/eduncan911/podcast?branch=master)" >> README.md
	echo "[![Go Report Card](https://goreportcard.com/badge/github.com/eduncan911/podcast)](https://goreportcard.com/report/github.com/eduncan911/podcast)" >> README.md
	echo "[![GoDoc](https://godoc.org/github.com/eduncan911/podcast?status.svg)](https://godoc.org/github.com/eduncan911/podcast)"
	echo "[![MIT License](https://img.shields.io/npm/l/mediaelement.svg)](https://eduncan911.mit-license.org/)" >> README.md
	echo  >>README.md
	cat README.md.tmp >> README.md
	rm README.md.tmp
