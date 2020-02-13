// Package podcast generates a fully compliant iTunes and RSS 2.0 podcast feed
// for GoLang using a simple API.
//
// Full documentation with detailed examples located at https://godoc.org/github.com/eduncan911/podcast
//
// Usage
//
// To use, `go get` and `import` the package like your typical GoLang library.
//
//     $ go get -u github.com/eduncan911/podcast
//
//     import "github.com/eduncan911/podcast"
//
// The API exposes a number of method receivers on structs that implements the
// logic required to comply with the specifications and ensure a compliant feed.
// A number of overrides occur to help with iTunes visibility of your episodes.
//
// Notably, the `Podcast.AddItem` function performs most
// of the heavy lifting by taking the `Item` input and performing
// validation, overrides and duplicate setters through the feed.
//
// Full detailed Examples of the API are at https://godoc.org/github.com/eduncan911/podcast.
//
// Go Modules
//
// This library is supported on GoLang 1.7 and higher.
//
// We have implemented Go Modules support and the CI pipeline shows it working with
// new installs, tested with Go 1.13.  To keep 1.7 compatibility, we use
// `go mod vendor` to maintain the `vendor/` folder for older 1.7 and later runtimes.
//
// If either runtime has an issue, please create an Issue and I will address.
//
// Extensibility
//
// For version 1.x, you are not restricted in having full control over your feeds.
// You may choose to skip the API methods and instead use the structs directly.  The
// fields have been grouped by RSS 2.0 and iTunes fields with iTunes specific fields
// all prefixed with the letter `I`.
//
// However, do note that the 2.x version currently in progress will break this
// extensibility and enforce API methods going forward. This is to ensure that the feed
// can both be marshalled, and unmarshalled back and forth (current 1.x branch can only
// be unmarshalled - hence the work for 2.x).
//
// Fuzzing Inputs
//
// `go-fuzz` has been added in 1.4.1, covering all exported API methods.  They have been
// ran extensively and no issues have come out of them yet (most tests were ran overnight,
// over about 11 hours with zero crashes).
//
// If you wish to help fuzz the inputs, with Go 1.13 or later you can run `go-fuzz` on any
// of the inputs.
//
//   go get -u github.com/dvyukov/go-fuzz/go-fuzz
//   go get -u github.com/dvyukov/go-fuzz/go-fuzz-build
//   go get -u github.com/eduncan911/podcast
//   cd $GOPATH/src/github.com/eduncan911/podcast
//   go-fuzz-build
//   go-fuzz -func FuzzPodcastAddItem
//
// To obtain a list of available funcs to pass, just run `go-fuzz` without any parameters:
//
//   $ go-fuzz
//   2020/02/13 07:27:32 -func flag not provided, but multiple fuzz functions available:
//   FuzzItemAddDuration, FuzzItemAddEnclosure, FuzzItemAddImage, FuzzItemAddPubDate,
//   FuzzItemAddSummary, FuzzPodcastAddAtomLink, FuzzPodcastAddAuthor, FuzzPodcastAddCategory,
//   FuzzPodcastAddImage, FuzzPodcastAddItem, FuzzPodcastAddLastBuildDate, FuzzPodcastAddPubDate,
//   FuzzPodcastAddSubTitle, FuzzPodcastAddSummary, FuzzPodcastBytes, FuzzPodcastEncode,
//   FuzzPodcastNew
//
// If you do find an issue, please raise an issue immediately and I will quickly address.
//
// Roadmap
//
// The 1.x branch is now mostly in maintenance mode, open to PRs.  This means no
// more planned features on the 1.x feature branch is expected. With the success of 6
// iTunes-accepted podcasts I have published with this library, and with the feedback from
// the community, the 1.x releases are now considered stable.
//
// The 2.x branch's primary focus is to allow for bi-direction marshalling both ways.
// Currently, the 1.x branch only allows unmarshalling to a serial feed.  An attempt to marshall
// a serialized feed back into a Podcast form will error or not work correctly.  Note that while
// the 2.x branch is targeted to remain backwards compatible, it is true if using the public
// API funcs to set parameters only.  Several of the underlying public fields are being removed
// in order to accommodate the marshalling of serialized data.  Therefore, a version 2.x is denoted
// for this release.
//
// Versioning
//
// We use SemVer versioning schema.  You can rest assured that pulling 1.x branches will
// remain backwards compatible now and into the future.
//
// However, the new 2.x branch, while keeping the same API, is expected break those that
// bypass the API methods and use the underlying public properties instead.
//
// Release Notes
//
// v1.4.2
//   * Slim down Go Modules for consumers (#32)
//
// v1.4.1
//   * Implement fuzz logic testing of exported funcs (#31)
//   * Upgrade CICD Pipeline Tooling (#31)
//   * Update documentation for 1.x and 2.3 (#31)
//   * Allow godoc2ghmd to run without network (#31)
//
// v1.4.0
//   * Add Go Modules, Update vendor folder (#26, #25)
//   * Add C.I. GitHub Actions (#25)
//   * Add additional error checks found by linters (#25)
//   * Go Fmt enclosure_test.go (#25)
//
// v1.3.2
//   * Correct count len of UTF8 strings (#9)
//   * Implement duration parser (#8)
//   * Fix Github and GoDocs Markdown (#14)
//   * Move podcast.go Private Methods to Respected Files (#12)
//   * Allow providing GUID on Podcast (#15)
//
// v1.3.1
//   * increased itunes compliance after feedback from Apple:
//     - specified what categories should be set with AddCategory().
//     - enforced title and link as part of Image.
//   * added Podcast.AddAtomLink() for more broad compliance to readers.
//
// v1.3.0
//   * fixes Item.Duration being set incorrectly.
//   * changed Item.AddEnclosure() parameter definition (Bytes not Seconds!).
//   * added Item.AddDuration formatting and override.
//   * added more documentation surrounding Item.Enclosure{}
//
// v1.2.1
//   * added Podcast.AddSubTitle() and truncating to 64 chars.
//   * added a number of Guards to protect against empty fields.
//
// v1.2.0
//   * added Podcast.AddPubDate() and Podcast.AddLastBuildDate() overrides.
//   * added Item.AddImage() to mask some cumbersome addition of IImage.
//   * added Item.AddPubDate to simply datetime setters.
//   * added more examples (mostly around Item struct).
//   * tweaked some documentation.
//
// v1.1.0
//   * Enabling CDATA in ISummary fields for Podcast and Channel.
//
// v1.0.0
//   * Initial release.
//   * Full documentation, full examples and complete code coverage.
//
// References
//
// RSS 2.0: https://cyber.harvard.edu/rss/rss.html
//
// Podcasts: https://help.apple.com/itc/podcasts_connect/#/itca5b22233
//
package podcast
