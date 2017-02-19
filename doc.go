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
// Notably, the [Podcast.AddItem(i Item)](#Podcast.AddItem) function performs most
// of the heavy lifting by taking the [Item](#Item) input and performing
// validation, overrides and duplicate setters through the feed.
//
// See the detailed Examples of the API at https://godoc.org/github.com/eduncan911/podcast.
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
// Release Notes
//
// 1.3.0
// * fixes Item.Duration being set incorrectly.
// * changed Item.AddEnclosure() parameter definition (Bytes not Seconds!).
// * added Item.AddDuration formatting and override.
// * added more documentation surrounding Item.Enclosure{}
//
// 1.2.1
// * added Podcast.AddSubTitle() and truncating to 64 chars.
// * added a number of Guards to protect against empty fields.
//
// 1.2.0
// * added Podcast.AddPubDate() and Podcast.AddLastBuildDate() overrides.
// * added Item.AddImage() to mask some cumbersome addition of IImage.
// * added Item.AddPubDate to simply datetime setters.
// * added more examples (mostly around Item struct).
// * tweaked some documentation.
//
// 1.1.0
// * Enabling CDATA in ISummary fields for Podcast and Channel.
//
// 1.0.0
// * Initial release.
// * Full documentation, full examples and complete code coverage.
//
package podcast
