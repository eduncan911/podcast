package podcast

import (
	"encoding/xml"
	"fmt"
	"time"
	"unicode/utf8"
)

// Item represents a single entry in a podcast.
//
// Article minimal requirements are:
// - Title
// - Description
// - Link
//
// Audio minimal requirements are:
// - Title
// - Description
// - Enclosure (HREF, Type and Length all required)
//
// Recommendations:
// - Setting the minimal fields sets most of other fields, including iTunes.
// - Use the Published time.Time setting instead of PubDate.
// - Always set an Enclosure.Length, to be nice to your downloaders.
// - Use Enclosure.Type instead of setting TypeFormatted for valid extensions.
type Item struct {
	XMLName xml.Name `xml:"item"`

	// GUID is the episode’s globally unique identifier (GUID). It is
	// recommended to set this tag on each item.
	//
	// It is very important that each episode have a unique GUID and
	// that it never changes, even if an episode’s metadata, like
	// title or enclosure URL, do change.
	GUID string `xml:"guid"`

	// Title is an episode title. It is a required field per iTunes
	// definitions.
	Title string `xml:"title"`

	// Link is an episode link URL.
	//
	// Do not use HTMl tags.  Only raw URLs such as https:// are allowed.
	Link string `xml:"link"`

	// Description is text containing one or more sentences describing
	// your episode to potential listeners.
	//
	// Use item.AddDescription(...) to populate this field correctly.
	Description *Description

	// PubDate is the date and time when an episode was released. It is
	// recommended to set this tag on each item.
	//
	// Use item.AddPubDate(...) to populate this field correctly.
	PubDate *time.Time `xml:"-"`

	// PubDateFormatted is deprecated.  Do not populate nor read this
	// string as it will be removed in a future release.
	PubDateFormatted string `xml:"pubDate,omitempty"`

	// Enclosure is of type podcast.Enclosure. It is a required field per
	// iTunes definitions.
	//
	// Use item.AddEnclosure(...) to populate this field correctly.
	Enclosure *Enclosure

	// IDuration is the duration of an episode.
	//
	// Use item.AddDuration(...) to populate this field correctly.
	IDuration string `xml:"itunes:duration,omitempty"`

	// IImage is of type podcast.Image.
	//
	// Use item.AddImage(...) to populate this field correctly.
	IImage *IImage

	// IExplicit defines the episode parental advisory information.
	//
	// Where the explicit value can be one of the following:
	//
	// "true" : If you specify true, indicating the presence of explicit
	// content, Apple Podcasts displays an Explicit parental advisory
	// graphic for your episode.
	// Episodes containing explicit material aren’t available in some Apple
	// Podcasts territories.
	//
	// "false" : If you specify false, indicating that the episode does not
	// contain explicit language or adult content, Apple Podcasts displays
	// a Clean parental advisory graphic for your episode.
	IExplicit string `xml:"itunes:explicit,omitempty"`

	// ITitle is a Situational episode title specific for Apple Podcasts.
	//
	// This tag is a string containing a clear concise name of your
	// episode on Apple Podcasts.
	//
	// Don’t specify the episode number or season number in this tag. Instead,
	// specify those details in the appropriate tags IEpisode and ISeason>.
	//
	// Also, don’t repeat the title of your show within your episode title.
	//
	// Separating episode and season number from the title makes it possible
	// for Apple to easily index and order content from all shows.
	ITitle string `xml:"itunes:title,omitempty"`

	// IEpisode is a Situational tag for the episode number.
	//
	// If all your episodes have numbers and you would like them to be ordered
	// based on them use this tag for each one.
	//
	// Episode numbers are optional for type episodic shows, but are
	// mandatory for serial shows.
	//
	// Where episode is a non-zero integer (1, 2, 3, etc.) representing your
	// episode number.
	IEpisode string `xml:"itunes:episode,omitempty"`

	// ISeason is a Situational tag for the episode season number.
	//
	// If an episode is within a season use this tag.
	//
	// Where season is a non-zero integer (1, 2, 3, etc.) representing your
	// season number.
	//
	// To allow the season feature for shows containing a single season, if
	// only one season exists in the RSS feed, Apple Podcasts doesn’t display
	// a season number. When you add a second season to the RSS feed, Apple
	// Podcasts displays the season numbers.
	ISeason string `xml:"itunes:season,omitempty"`

	// IEpisodeType is a Situational tag for the episode type.
	//
	// If an episode is a trailer or bonus content, use this tag.
	//
	// Use AddEpisodeType(...) to populate this field correctly.
	IEpisodeType *IEpisodeType

	// IBlock is a Situational tag to show or hide the status of the episode.
	//
	// If you want an episode removed from the Apple directory, use this tag.
	//
	// Specifying the tag with a "Yes" value prevents that episode from
	// appearing in Apple Podcasts.
	//
	// For example, you might want to block a specific episode if you know
	// that its content would otherwise cause the entire podcast to be
	// removed from Apple Podcasts.
	//
	// Specifying any value other than Yes has no effect.
	IBlock string `xml:"itunes:block,omitempty"`

	// As of April 2019, the following tags are no longer listed in iTunes'
	// supported tags.  However, most are still listed under Harvard's
	// definition of RSS feed tags with channels and episodes.  See Harvard's
	// defnitions for more info.
	//
	// https://cyber.harvard.edu/rss/rss.html
	//
	// This does not mean they are not supported; it is just that Apple
	// has chosen to remove their listings and descriptions.
	//

	Author             *Author `xml:"-"`
	AuthorFormatted    string  `xml:"author,omitempty"`
	Category           string  `xml:"category,omitempty"`
	Comments           string  `xml:"comments,omitempty"`
	IAuthor            string  `xml:"itunes:author,omitempty"`
	IIsClosedCaptioned string  `xml:"itunes:isClosedCaptioned,omitempty"`
	IOrder             string  `xml:"itunes:order,omitempty"`
	ISubtitle          string  `xml:"itunes:subtitle,omitempty"`
	ISummary           *ISummary
	Source             string `xml:"source,omitempty"`
}

// AddEnclosure adds the downloadable asset to the podcast Item.
func (i *Item) AddEnclosure(
	url string, enclosureType EnclosureType, lengthInBytes int64) {
	i.Enclosure = &Enclosure{
		URL:    url,
		Type:   enclosureType,
		Length: lengthInBytes,
	}
}

// AddEpisode adds the situational tags with rules to iTunes' episodes.
// Using this function will ensure a properly formatted episode has been
// added to the feed in compliance to iTunes' requirements.
func (i *Item) AddEpisode() {
	// As of April 2019:
	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	//
	// If an episode is a trailer or bonus content, use this tag.
	//
	// Where the episodeType value can be one of the following:
	// - Full (default). Specify full when you are submitting the complete
	// content of your show.
	// - Trailer. Specify trailer when you are submitting a short, promotional
	// piece of content that represents a preview of your current show.
	// - Bonus. Specify bonus when you are submitting extra content for your
	// show (for example, behind the scenes information or interviews with
	// the cast) or cross-promotional content for another show.
	//
	// The rules for using trailer and bonus tags depend on whether the
	// <itunes:season> and <itunes:episode> tags have values:
	//
	// Trailer:
	// - No season or episode number: a show trailer
	// - A season number and no episode number: a season trailer. (Note: an
	// episode trailer should have a different <guid> than the actual episode)
	// - Episode number and optionally a season number: an episode
	// trailer/teaser, later replaced with the actual episode
	//
	// Bonus:
	// - No season or episode number: a show bonus
	// - A season number: a season bonus
	// - Episode number and optionally a season number: a bonus episode related
	// to a specific episode
	//

	i.IEpisodeType = &EpisodeType{}
}

// AddImage adds the image as an iTunes-only IImage.  RSS 2.0 does not have
// the specification of Images at the Item level.
//
// Podcast feeds contain artwork that is a minimum size of
// 1400 x 1400 pixels and a maximum size of 3000 x 3000 pixels,
// 72 dpi, in JPEG or PNG format with appropriate file
// extensions (.jpg, .png), and in the RGB colorspace. To optimize
// images for mobile devices, Apple recommends compressing your
// image files.
func (i *Item) AddImage(url string) {
	if len(url) > 0 {
		i.IImage = &IImage{HREF: url}
	}
}

// AddPubDate adds the datetime as a parsed PubDate.
//
// UTC time is used by default.
func (i *Item) AddPubDate(datetime *time.Time) {
	i.PubDate = datetime
	i.PubDateFormatted = parseDateRFC1123Z(i.PubDate)
}

// AddSummary adds the iTunes summary.
//
// Limit: 4000 characters
//
// Note that this field is a CDATA encoded field which allows for rich text
// such as html links: `<a href="http://www.apple.com">Apple</a>`.
func (i *Item) AddSummary(summary string) {
	count := utf8.RuneCountInString(summary)
	if count > 4000 {
		s := []rune(summary)
		summary = string(s[0:4000])
	}
	i.ISummary = &ISummary{
		Text: summary,
	}
}

// AddDescription adds a rich-text description tag.
//
// Limit: 4000 characters
//
// Note that this field is a CDATA encoded field which allows for rich text
// such as html links: `<a href="http://www.apple.com">Apple</a>`.
func (i *Item) AddDescription(d string) {
	count := utf8.RuneCountInString(d)
	if count > 4000 {
		s := []rune(d)
		d = string(s[0:4000])
	}
	i.Description = &Description{
		Text: d,
	}
}

// AddDuration adds the duration to the iTunes duration field.
func (i *Item) AddDuration(durationInSeconds int64) {
	if durationInSeconds <= 0 {
		return
	}
	i.IDuration = parseDuration(durationInSeconds)
}

var parseDuration = func(duration int64) string {
	h := duration / 3600
	duration = duration % 3600

	m := duration / 60
	duration = duration % 60

	s := duration

	// HH:MM:SS
	if h > 9 {
		return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	}

	// H:MM:SS
	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}

	// MM:SS
	if m > 9 {
		return fmt.Sprintf("%02d:%02d", m, s)
	}

	// M:SS
	return fmt.Sprintf("%d:%02d", m, s)
}

var parseDateRFC1123Z = func(t *time.Time) string {
	if t != nil && !t.IsZero() {
		return t.Format(time.RFC1123Z)
	}
	return time.Now().UTC().Format(time.RFC1123Z)
}
