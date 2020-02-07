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
// Extensibility
//
// In no way are you restricted in having full control over your feeds.  You may
// choose to skip the API methods and instead use the structs directly.  The
// fields have been grouped by RSS 2.0 and iTunes fields.
//
// iTunes specific fields are all prefixed with the letter `I`.
//
// References
//
// RSS 2.0: https://cyber.harvard.edu/rss/rss.html
//
// Podcasts: https://help.apple.com/itc/podcasts_connect/#/itca5b22233
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
// 1.4.0
//   * Add C.I. GitHub Actions (#25)
//   * Add Go Modules (#26)
//
// 1.3.2
//   * Correct count len of UTF8 strings (#9)
//   * Implement duration parser (#8)
//   * Fix Github and GoDocs Markdown (#14)
//   * Move podcast.go Private Methods to Respected Files (#12)
//   * Allow providing GUID on Podcast (#15)
//
// 1.3.1
//   * increased itunes compliance after feedback from Apple:
//     - specified what categories should be set with AddCategory().
//     - enforced title and link as part of Image.
//   * added Podcast.AddAtomLink() for more broad compliance to readers.
//
// 1.3.0
//   * fixes Item.Duration being set incorrectly.
//   * changed Item.AddEnclosure() parameter definition (Bytes not Seconds!).
//   * added Item.AddDuration formatting and override.
//   * added more documentation surrounding Item.Enclosure{}
//
// 1.2.1
//   * added Podcast.AddSubTitle() and truncating to 64 chars.
//   * added a number of Guards to protect against empty fields.
//
// 1.2.0
//   * added Podcast.AddPubDate() and Podcast.AddLastBuildDate() overrides.
//   * added Item.AddImage() to mask some cumbersome addition of IImage.
//   * added Item.AddPubDate to simply datetime setters.
//   * added more examples (mostly around Item struct).
//   * tweaked some documentation.
//
// 1.1.0
//   * Enabling CDATA in ISummary fields for Podcast and Channel.
//
// 1.0.0
//   * Initial release.
//   * Full documentation, full examples and complete code coverage.
//
package podcast
