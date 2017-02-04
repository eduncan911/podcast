// Package podcast is an iTunes and RSS 2.0 podcast generator for GoLang that
// enforces strict compliance by using its simple interface.
//
// Full documentation with detailed examples located at https://godoc.org/github.com/eduncan911/podcast
//
// Usage
//
//     $ go get -u github.com/eduncan911/podcast
//
// The API exposes a number of method receivers on structs that implements the
// logic required to comply with the specifications and ensure a compliant feed.
// A number of overrides occur to help with iTunes visibility of your episodes.
//
// Notably, the [Podcast.AddItem(i Item)](#Podcast.AddItem) function performs most of the
// heavy lifting by taking the [Item](#Item) input and performing validation, overrides
// and duplicate setters through the feed.
//
// See the detailed Examples in the GoDocs for complete usage.
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
// 1.0.0
// * Initial release.
// * Full documentation, full examples and complete code coverage.
//
package podcast
