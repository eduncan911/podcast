package podcast

import (
	"encoding/xml"
	"time"
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
	XMLName          xml.Name   `xml:"item"`
	GUID             string     `xml:"guid"`
	Title            string     `xml:"title"`
	Link             string     `xml:"link"`
	Description      string     `xml:"description"`
	Author           *Author    `xml:"-"`
	AuthorFormatted  string     `xml:"author,omitempty"`
	Category         string     `xml:"category,omitempty"`
	Comments         string     `xml:"comments,omitempty"`
	Source           string     `xml:"source,omitempty"`
	PubDate          *time.Time `xml:"-"`
	PubDateFormatted string     `xml:"pubDate,omitempty"`
	Enclosure        *Enclosure

	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	IAuthor            string `xml:"itunes:author,omitempty"`
	ISubtitle          string `xml:"itunes:subtitle,omitempty"`
	ISummary           *ISummary
	IImage             *IImage
	IDuration          string `xml:"itunes:duration,omitempty"`
	IExplicit          string `xml:"itunes:explicit,omitempty"`
	IIsClosedCaptioned string `xml:"itunes:isClosedCaptioned,omitempty"`
	IOrder             string `xml:"itunes:order,omitempty"`
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
// such as html links: <a href="http://www.apple.com">Apple</a>.
func (i *Item) AddSummary(summary string) {
	if len(summary) > 4000 {
		s := []rune(summary)
		summary = string(s[0:4000])
	}
	i.ISummary = &ISummary{
		Text: summary,
	}
}

// AddDuration adds the duration to the iTunes duration field.
func (i *Item) AddDuration(durationInSeconds int64) {
	if durationInSeconds <= 0 {
		return
	}
	i.IDuration = parseDuration(durationInSeconds)
}
