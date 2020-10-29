package podcast

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/pkg/errors"
)

const (
	pVersion = "1.3.1"
)

// Podcast represents a podcast.
type Podcast struct {
	XMLName xml.Name `xml:"channel"`

	// Title is the show title.
	//
	// This is a required tag.
	//
	// It’s important to have a clear, concise name for your podcast. Make your
	// title specific. A show titled Our Community Bulletin is too vague to
	// attract many subscribers, no matter how compelling the content.
	//
	// Pay close attention to the title as Apple Podcasts uses this field fo
	// search.
	//
	// If you include a long list of keywords in an attempt to game podcast
	// search, your show may be removed from the Apple directory.
	Title string `xml:"title"`

	// Link is the associated with a podcast.
	//
	// Do not specify HTML here.  Use RAW https:// urls.
	Link string `xml:"link"`

	// Description is text containing one or more sentences describing
	// your podcast to potential listeners.
	//
	// This is a required tag.
	//
	// Limit: 4000 characters
	//
	// Note that this field is a CDATA encoded field which allows for rich text
	// such as html links: `<a href="http://www.apple.com">Apple</a>`.
	//
	// Use podcast.New(...) to populate this field correctly.
	Description *Description

	// Language is language spoken on the show.
	//
	// This is a required tag.
	//
	// Because Apple Podcasts is available in territories around the world,
	// it is critical to specify the language of a podcast. Apple Podcasts
	// only supports values from the ISO 639 list (two-letter language codes,
	// with some possible modifiers, such as "en-us").
	//
	// Invalid language codes will cause your feed to fail Apple validation.
	Language string `xml:"language,omitempty"`

	// IAuthor is the group responsible for creating the show.
	//
	// Note the difference of this itunes tag from Harvard's RSS definition
	// for author.  Harvard's requirement is that this is in the format of:
	//
	// Full Name (emai@address)
	//
	// Whereas iTunes defines this field as the following:
	//
	// Show author most often refers to the parent company or network of a
	// podcast, but it can also be used to identify the host(s) if none exists.
	//
	// Author information is especially useful if a company or organization
	// publishes multiple podcasts. Providing this information will allow
	// listeners to see all shows created by the same entity.
	IAuthor string `xml:"itunes:author,omitempty"`

	// IImage is of type podcast.Image.
	//
	// This is a required tag.
	//
	// Use podcast.AddImage(...) to populate this field correctly.
	IImage *IImage

	// IExplicit defines the parental advisory information.
	//
	// This is a required tag.
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

	// IOwner is the podcast owner contact information.
	//
	// Use AddAuthor(...) to set this field correctly.
	//
	// Note: This tag information is for administrative communication about
	// the podcast and isn’t displayed in Apple Podcasts. Please make sure
	// the email address is active and monitored.
	IOwner *Author

	// ICategories is an array of specific iTunes categories.
	//
	// This is a required tag.
	//
	// Use podcast.AddCategories(...) to properly populate.
	ICategories []*ICategory

	// ITitle is a Situational tag that is the title specific for Apple
	// Podcasts.
	//
	// Use podcast.New(...) to properly set this field.
	ITitle string `xml:"itunes:title,omitempty"`

	// IType is a Situational tag of the type of the show.
	//
	// If your show is Serial you must use this tag.
	//
	// Use podcast.AddType(...) to properly set this field.
	IType *IType

	// Copyright is a Situational tag for the show's copyright details.
	Copyright string `xml:"copyright,omitempty"`

	// INewFeedURL is a Situational tag to specify new podcast RSS Feed URL.
	//
	// If you change the URL of your podcast feed, you should use this tag in
	// your new feed.
	//
	// Use this tag to manually change the URL where your podcast is located.
	//
	// You should maintain your old feed until you have migrated your existing
	// subscribers. For more information, see Update your RSS feed URL.
	//
	// Note: This tag reports new feed URLs to Apple Podcasts and isn’t
	// displayed in Apple Podcasts.
	INewFeedURL string `xml:"itunes:new-feed-url,omitempty"`

	// IBlock is a Situational tag to show or hide the status of the the
	// podcasts.
	//
	// If you want your show removed from the Apple directory, use this tag.
	//
	// Specifying the tag with a "Yes" value prevents that this show from
	// appearing in Apple Podcasts.
	//
	// Specifying any value other than Yes has no effect.
	IBlock string `xml:"itunes:block,omitempty"`

	// IComplete is a Situational tag for the podcast update status.
	//
	// If you will never publish another episode to your show, use this tag.
	//
	// Specifying the <itunes:complete> tag with a Yes value indicates that a
	// podcast is complete and you will not post any more episodes in the
	// future.
	//
	// Specifying any value other than Yes has no effect.
	IComplete string `xml:"itunes:complete,omitempty"`

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

	AtomLink       *AtomLink
	Category       string `xml:"category,omitempty"`
	Cloud          string `xml:"cloud,omitempty"`
	Docs           string `xml:"docs,omitempty"`
	Generator      string `xml:"generator,omitempty"`
	IDuration      string `xml:"itunes:duration,omitempty"`
	ISubtitle      string `xml:"itunes:subtitle,omitempty"`
	ISummary       *ISummary
	Image          *Image
	LastBuildDate  string `xml:"lastBuildDate,omitempty"`
	ManagingEditor string `xml:"managingEditor,omitempty"`
	PubDate        string `xml:"pubDate,omitempty"`
	Rating         string `xml:"rating,omitempty"`
	SkipHours      string `xml:"skipHours,omitempty"`
	SkipDays       string `xml:"skipDays,omitempty"`
	TextInput      *TextInput
	TTL            int    `xml:"ttl,omitempty"`
	WebMaster      string `xml:"webMaster,omitempty"`

	// Items is a collection of 0..n episodes for this podcast.
	Items  []*Item
	encode func(w io.Writer, o interface{}) error
}

// New instantiates a Podcast with required parameters.
//
// Nil-able fields are optional but recommended as they are formatted
// to the expected proper formats.
func New(title, link, description string,
	pubDate, lastBuildDate *time.Time) Podcast {
	return Podcast{
		Title:         title,
		ITitle:        title,
		Description:   parseDescription(description),
		Link:          link,
		Generator:     fmt.Sprintf("go podcast v%s (github.com/eduncan911/podcast)", pVersion),
		PubDate:       parseDateRFC1123Z(pubDate),
		LastBuildDate: parseDateRFC1123Z(lastBuildDate),
		Language:      "en-us",

		// setup dependency (could inject later)
		encode: encoder,
	}
}

// AddAuthor adds the specified Author to the podcast's IOwner and
// Harvard's ManagingEditor tags.
func (p *Podcast) AddAuthor(name, email string) {
	if len(email) == 0 {
		return
	}
	a := &Author{
		Name:  name,
		Email: email,
	}
	p.IOwner = a
	p.ManagingEditor = parseAuthorNameEmail(a)
}

// AddAtomLink adds a FQDN reference to an atom feed.
func (p *Podcast) AddAtomLink(href string) {
	if len(href) == 0 {
		return
	}
	p.AtomLink = &AtomLink{
		HREF: href,
		Rel:  "self",
		Type: "application/rss+xml",
	}
}

// AddCategory adds the category to the Podcast.
//
// ICategory can be listed multiple times.
//
// Calling this method multiple times will APPEND the category to the existing
// list, if any, including ICategory.
//
// Note that Apple iTunes has a specific list of categories that only can be
// used and will invalidate the feed if deviated from the list.  The list
// changes occassionally.  Please refer to the following link for the updated
// list:
//
// https://help.apple.com/itc/podcasts_connect/#/itc9267a2f12
func (p *Podcast) AddCategory(category string, subCategories []string) {
	if len(category) == 0 {
		return
	}

	// RSS 2.0 Category only supports 1-tier
	if len(p.Category) > 0 {
		p.Category = p.Category + "," + category
	} else {
		p.Category = category
	}

	icat := ICategory{Text: category}
	for _, c := range subCategories {
		if len(c) == 0 {
			continue
		}
		icat2 := ICategory{Text: c}
		icat.ICategories = append(icat.ICategories, &icat2)
	}
	p.ICategories = append(p.ICategories, &icat)
}

// AddImage adds the specified Image to the Podcast.
//
// Podcast feeds contain artwork that is a minimum size of
// 1400 x 1400 pixels and a maximum size of 3000 x 3000 pixels,
// 72 dpi, in JPEG or PNG format with appropriate file
// extensions (.jpg, .png), and in the RGB colorspace. To optimize
// images for mobile devices, Apple recommends compressing your
// image files.
func (p *Podcast) AddImage(url string) {
	if len(url) == 0 {
		return
	}
	p.Image = &Image{
		URL:   url,
		Title: p.Title,
		Link:  p.Link,
	}
	p.IImage = &IImage{HREF: url}
}

// AddItem adds the podcast episode.  It returns a count of Items added or any
// errors in validation that may have occurred.
//
// This method takes the "itunes overrides" approach to populating
// itunes tags according to the overrides rules in the specification.
// This not only complies completely with iTunes parsing rules; but, it also
// displays what is possible to be set on an individual episode level – if you
// wish to have more fine grain control over your content.
//
// This method imposes strict validation of the Item being added to confirm
// to Podcast and iTunes specifications.
//
// Article minimal requirements are:
//
//   * Title
//   * Description
//   * Link
//
// Audio, Video and Downloads minimal requirements are:
//
//   * Title
//   * Description
//   * Enclosure (HREF, Type and Length all required)
//
// The following fields are always overwritten (don't set them):
//
//   * GUID
//   * PubDateFormatted
//   * AuthorFormatted
//   * Enclosure.TypeFormatted
//   * Enclosure.LengthFormatted
//
// Recommendations:
//
//   * Just set the minimal fields: the rest get set for you.
//   * Always set an Enclosure.Length, to be nice to your downloaders.
//   * Follow Apple's best practices to enrich your podcasts:
//     https://help.apple.com/itc/podcasts_connect/#/itc2b3780e76
//   * For specifications of itunes tags, see:
//     https://help.apple.com/itc/podcasts_connect/#/itcb54353390
//
func (p *Podcast) AddItem(i Item) (int, error) {
	// initial guards for required fields
	if len(i.Title) == 0 || len(i.Description) == 0 {
		return len(p.Items), errors.New("Title and Description are required")
	}
	if i.Enclosure != nil {
		if len(i.Enclosure.URL) == 0 {
			return len(p.Items),
				errors.New(i.Title + ": Enclosure.URL is required")
		}
		if i.Enclosure.Type.String() == enclosureDefault {
			return len(p.Items),
				errors.New(i.Title + ": Enclosure.Type is required")
		}
	} else if len(i.Link) == 0 {
		return len(p.Items),
			errors.New(i.Title + ": Link is required when not using Enclosure")
	}

	// corrective actions and overrides
	//
	i.PubDateFormatted = parseDateRFC1123Z(i.PubDate)
	i.AuthorFormatted = parseAuthorNameEmail(i.Author)
	if i.Enclosure != nil {
		if len(i.GUID) == 0 {
			i.GUID = i.Enclosure.URL // yep, GUID is the Permlink URL
		}

		if i.Enclosure.Length < 0 {
			i.Enclosure.Length = 0
		}
		i.Enclosure.LengthFormatted = strconv.FormatInt(i.Enclosure.Length, 10)
		i.Enclosure.TypeFormatted = i.Enclosure.Type.String()

		// allow Link to be set for article references to Downloads,
		// otherwise set it to the enclosurer's URL.
		if len(i.Link) == 0 {
			i.Link = i.Enclosure.URL
		}
	} else {
		i.GUID = i.Link // yep, GUID is the Permlink URL
	}

	// iTunes it
	//
	if len(i.IAuthor) == 0 {
		switch {
		case i.Author != nil:
			i.IAuthor = i.Author.Email
		case len(p.IAuthor) != 0:
			i.Author = &Author{Email: p.IAuthor}
			i.IAuthor = p.IAuthor
		case len(p.ManagingEditor) != 0:
			i.Author = &Author{Email: p.ManagingEditor}
			i.IAuthor = p.ManagingEditor
		}
	}
	if i.IImage == nil {
		if p.Image != nil {
			i.IImage = &IImage{HREF: p.Image.URL}
		}
	}

	p.Items = append(p.Items, &i)
	return len(p.Items), nil
}

// AddPubDate adds the datetime as a parsed PubDate.
//
// UTC time is used by default.
func (p *Podcast) AddPubDate(datetime *time.Time) {
	p.PubDate = parseDateRFC1123Z(datetime)
}

// AddLastBuildDate adds the datetime as a parsed PubDate.
//
// UTC time is used by default.
func (p *Podcast) AddLastBuildDate(datetime *time.Time) {
	p.LastBuildDate = parseDateRFC1123Z(datetime)
}

// AddSubTitle adds the iTunes subtitle that is displayed with the title
// in iTunes.
//
// Note that this field should be just a few words long according to Apple.
// This method will truncate the string to 64 chars if too long with "..."
func (p *Podcast) AddSubTitle(subTitle string) {
	count := utf8.RuneCountInString(subTitle)
	if count == 0 {
		return
	}
	if count > 64 {
		s := []rune(subTitle)
		subTitle = string(s[0:61]) + "..."
	}
	p.ISubtitle = subTitle
}

// AddSummary adds the iTunes summary.
//
// Limit: 4000 characters
//
// Note that this field is a CDATA encoded field which allows for rich text
// such as html links: `<a href="http://www.apple.com">Apple</a>`.
func (p *Podcast) AddSummary(summary string) {
	count := utf8.RuneCountInString(summary)
	if count == 0 {
		return
	}
	if count > 4000 {
		s := []rune(summary)
		summary = string(s[0:4000])
	}
	p.ISummary = &ISummary{
		Text: summary,
	}
}

func (p *Podcast) AddType(type *Type) {
	// Its values can be one of the following:
	//
	// Episodic (default). Specify episodic when episodes are intended to be
	// consumed without any specific order. Apple Podcasts will present newest
	// episodes first and display the publish date (required) of each episode.
	// If organized into seasons, the newest season will be presented first
	// - otherwise, episodes will be grouped by year published, newest first.
	//
	// For new subscribers, Apple Podcasts adds the newest, most recent episode
	// in their Library.
	//
	// Serial. Specify serial when episodes are intended to be consumed in
	// sequential order. Apple Podcasts will present the oldest episodes
	// first and display the episode numbers (required) of each episode. If
	// organized into seasons, the newest season will be presented first and
	// <itunes:episode> numbers must be given for each episode.
	//
	// For new subscribers, Apple Podcasts adds the first episode to their
	// Library, or the entire current season if using seasons.
	
}

// Bytes returns an encoded []byte slice.
func (p *Podcast) Bytes() []byte {
	return []byte(p.String())
}

// Encode writes the bytes to the io.Writer stream in RSS 2.0 specification.
func (p *Podcast) Encode(w io.Writer) error {
	if _, err := w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")); err != nil {
		return errors.Wrap(err, "podcast.Encode: w.Write return error")
	}

	atomLink := ""
	if p.AtomLink != nil {
		atomLink = "http://www.w3.org/2005/Atom"
	}
	wrapped := podcastWrapper{
		ITUNESNS: "http://www.itunes.com/dtds/podcast-1.0.dtd",
		ATOMNS:   atomLink,
		Version:  "2.0",
		Channel:  p,
	}
	return p.encode(w, wrapped)
}

// String encodes the Podcast state to a string.
func (p *Podcast) String() string {
	b := new(bytes.Buffer)
	if err := p.Encode(b); err != nil {
		return "String: podcast.write returned the error: " + err.Error()
	}
	return b.String()
}

// // Write implements the io.Writer interface to write an RSS 2.0 stream
// // that is compliant to the RSS 2.0 specification.
// func (p *Podcast) Write(b []byte) (n int, err error) {
// 	buf := bytes.NewBuffer(b)
// 	if err := p.Encode(buf); err != nil {
// 		return 0, errors.Wrap(err, "Write: podcast.encode returned error")
// 	}
// 	return buf.Len(), nil
// }

type podcastWrapper struct {
	XMLName  xml.Name `xml:"rss"`
	Version  string   `xml:"version,attr"`
	ATOMNS   string   `xml:"xmlns:atom,attr,omitempty"`
	ITUNESNS string   `xml:"xmlns:itunes,attr"`
	Channel  *Podcast
}

var encoder = func(w io.Writer, o interface{}) error {
	e := xml.NewEncoder(w)
	e.Indent("", "  ")
	if err := e.Encode(o); err != nil {
		return errors.Wrap(err, "podcast.encoder: e.Encode returned error")
	}
	return nil
}

var parseAuthorNameEmail = func(a *Author) string {
	var author string
	if a != nil {
		author = a.Email
		if len(a.Name) > 0 {
			author = fmt.Sprintf("%s (%s)", a.Email, a.Name)
		}
	}
	return author
}

var parseDescription = func(d string) *Description {
	count := utf8.RuneCountInString(d)
	if count > 4000 {
		s := []rune(d)
		d = string(s[0:4000])
	}
	return &Description{
		Text: d,
	}
}
