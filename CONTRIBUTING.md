# Contributing to this repository

## Getting Started

Before you begin:

* Stay compatible with Go 1.7.  Do not use [newer language changes](https://golang.org/doc/go1.13#language) in later versions of Go.
* This repo targets 100% code coverage.  See **Testing Strategy, Documentation, and Examples** below for some cool details.
* Ensure you have the most common golang Linters enabled and running as all PRs run them.

## The Short Short Version

* **Create PRs against `develop` branch only.** The default branch is `master` for those that use `go get -u`.
* Check your PR for C.I. testing results once they complete (see **Testing Strategy, Documentation, and Examples** below). Once merged to `develop`, the author (me) will need to create a release branch. `go mod` and Github Releases are tagged from `master` once the release branch is completed.
* The `README.md` is auto-generated via the `Makefile`.  If you need to update it, please change `doc.go` and/or comments in code.

## Coding Guidelines

This repo tries to adhere to the basic set of Golang rules in the community:

* This repo strictly follows the [Effective Go](https://golang.org/doc/effective_go.html) coding guidelines.  
* This repo follows the `gofmt` formatting rules. Ensure it's enabled in your IDE of choice.
* Functions cannot have more than 20 lines of code, or it will flag external package auditors.  The `podcast.AddItem(...)` func being an exception for classroom and lunch-n-learn training scenarios.
* Avoid adding additional packages.  This repo is supposed to be stand-alone (with the exception of tests).
* When needing to modify comments, `doc.go`, Example text, please follow [Godoc-tricks](https://github.com/fluhus/godoc-tricks).

Please try to follow Chris Beams' [How to Write a Git Commit Message](https://chris.beams.io/posts/git-commit/), **specifically when opening the PR**. This saves me from having to rename your PR title.  :)  What you do once the PR is open and described is fair game, commit at will.

## SemVer and Breaking Changes

This repo uses [Semantic Versioning](https://semver.org/) for all tags and releases.  

Be cautious of any public API changes.  Additions are welcome; however, changes and deletions to the public API are greatly discouraged unless required (for example, the April 2019 iTunes' tag changes).

If a breaking change must occur, branch from `develop` like normal and call it out in the PR description.

## Testing Strategy, Documentation, and Examples

I use this repository as a showcase of a new testing methodology I called, `Minimalist Testing`.  As a matter of fact, it's the whole reason I created this repo.  :)

### In Summary

* 1st: Use [Golang Examples](https://blog.golang.org/examples) for all positive outcomes **against the public API only, not internal**, and against every public API method.  See `examples_tests.go`.
* 2nd: Test negative outcomes, again, **against the public API only, not internal** using standard testing.  Please try to follow `podcast_test.go` as an example.
* 3rd: Any remaining code-coverage that cannot be covered via the above, can be added into `_internals_test.go` files as needed.

Surprising, following the 1st and 2nd rules above, in that order, has an enormous amount of benefits:

* Usually covers about 92% to 97% of all code coverage!  Only thing left are some low-hanging internal private methods.
* Automatic full examples, already written out, and [documented publicly](https://godoc.org/github.com/eduncan911/podcast#pkg-examples) from just writing examples in the 1st step above.  All because you simply used `examples` to test positive outcomes, and not unit tests.
* You get a feel of how clean (or dirty) your public API is via your IDE's intellitype as you are coding against the public API.

Then you end up focusing on just a few tweaks to that 5% of remaining internal private methods, covering those edge cases and error conditions.

## How to pull this off

Basically, just focus on writing Examples instead of just purely writing unit tests.  It's fun, and you actually get to see your code functional as you grow and actually see the outputs.  You don't have to focus on TDD principals: just write your code and start writing your first Example, until it works.  It's almost like [Cowboy Coding](https://en.wikipedia.org/wiki/Cowboy_coding), except you see how you will be using your public API in real-time (and largely make changes right then and there to clean this or that up - which has saved loads of time for me).

When you need to test some negative outcomes (like guards and error handling), then add a few normal unit tests - but only against the public API (see first bullet point below).  This surprisingly lends well to covering the vast majority of packages.

* Test the public API by using a secondary package called `package podcast_test`.  Take note that `podcast_test.go` and `examples_test.go` do not use the default package; but instead, they use a different `package podcast_test`.  This means only the public API methods are exposed to your coding.
* Test as much as you can via the public API.  If you cannot reach into an internal use-case, then use `_internals_test.go` files as needed.
* Internal private tests, again only about 5% remaining, are usually the only ones that need some unique IoC technique.  

### Inversion of Control

I know, a bad word in Golang; but, a necessary evil to obtain 100% code coverage.

An example of D.I. is with the internal `podcast.encode(...)` pointer receiver on type `podcast`. By defining a field that I implement at runtime, instead of just writing a private func, I can assign the internal method `encoder(...)` that matches the method's signature at runtime (within the `New()` method).  However, at testing time, I set it internally to something else for my test setup conditions.

We can take a look at `podcast_internal_test.go` for an example of defining a custom `encode()` pointer receiver at testing time:

```
func TestStringError(t *testing.T) {
	t.Parallel()

	// arrange
	e := "TestEncodeError error result"
	p := Podcast{}
	p.encode = func(w io.Writer, o interface{}) error {
		return errors.New(e)
	}

	// act
	r := p.String()

	// assert
	assert.Contains(t, r, e)
}
```

As you can see, `p.encode` gets assigned an inline method under our control.  In this case, I want to test the error handling of `p.String()`.  The only way to do that is to override the `encode()` functionality that would otherwise just be a private method.

### Clean Golang Examples

Pay particular attention to how clean the Example test code is (see `examples_test.go`) as the test code actually shows up as documentation.

### Enough!  How to run tests

Simple: All Golang Examples, public API unit tests, and internal private tests can all be run simply by:

```
$ go test -v -cover
...
PASS
coverage: 100.0% of statements
ok  	github.com/eduncan911/podcast	0.005s
```

You also get to narrow down to individual tests, or even groups of tests:

```
$ go test -v -cover -run Example
...
PASS
coverage: 70.7% of statements
ok  	github.com/eduncan911/podcast	0.004s

$ go test -v -cover -run Test
PASS
coverage: 87.2% of statements
ok  	github.com/eduncan911/podcast	0.005s
```

Did you notice that both Positive tests (Examples) and Negative tests (non-internal) overlap dramatically in code coverage?  This is an added bonus in capturing additional edge cases with minimal testing code.

## Testing Tips

One tip is to enable Test-On-Save, with the addition of Show Coverage, of your favorite IDE. So when you save a file, it immediately runs all tests in your package, as well as adding the `-cover` attribute to include code coverage.  The advantage to this setting is most IDEs give a visual indicator of what code has and hasn't been covered. 

## My PR Tests are Failing

First, check the outcome of each test.  It usually has deep error messages to help debug.

If you want to run the C.I. tests locally, you can see the latest version of C.I. testing here:

https://github.com/eduncan911/podcast/blob/master/.github/workflows/go-cicd.yaml#L21

## Thank you!

Most of all, thank you for your help!

---

## How to Publish

I branch and publish this repo with the following steps.  Note that [I have automatic rebasing in my `~/.gitconfig`](https://github.com/eduncan911/dotfiles/blob/master/.gitconfig#L1-L2) when doing all `git pulls`.

I still prefer to do this manually for now; but, I may move it to a Github action later.

```
git checkout develop
git checkout -b feature-x
# commit everything
# open PR against develop branch
# squash and merge to develop branch
git checkout develop
git pull origin develop
git checkout master
git pull origin master
# push develop onto of master
git push origin develop:master
# Create Release/Tag from Github UI, since it's much prettier
```

TODO:
  * Automate release notes via commits and PR notes
