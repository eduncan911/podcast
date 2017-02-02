

# podcast
`import "github.com/eduncan911/podcast"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
Package `podcast` is an iTunes and RSS 2.0 podcast generator for GoLang that
enforces strict compliance by using its simple interface.

[![GoDoc](<a href="https://godoc.org/github.com/eduncan911/podcast?status.svg">https://godoc.org/github.com/eduncan911/podcast?status.svg</a>)](<a href="https://godoc.org/github.com/eduncan911/podcast">https://godoc.org/github.com/eduncan911/podcast</a>) [![Build Status](<a href="https://travis-ci.org/eduncan911/podcast.svg?branch=master">https://travis-ci.org/eduncan911/podcast.svg?branch=master</a>)](<a href="https://travis-ci.org/eduncan911/podcast">https://travis-ci.org/eduncan911/podcast</a>) [![Go Report Card](<a href="https://goreportcard.com/badge/github.com/eduncan911/podcast">https://goreportcard.com/badge/github.com/eduncan911/podcast</a>)](<a href="https://goreportcard.com/report/github.com/eduncan911/podcast">https://goreportcard.com/report/github.com/eduncan911/podcast</a>)

Full documentation with detailed examples located at [![GoDoc](<a href="https://godoc.org/github.com/eduncan911/podcast?status.svg">https://godoc.org/github.com/eduncan911/podcast?status.svg</a>)](<a href="https://godoc.org/github.com/eduncan911/podcast">https://godoc.org/github.com/eduncan911/podcast</a>)

Usage


	$ go get -u github.com/eduncan911/podcast

The API exposes a number of method receivers on structs that implements the
logic required to comply with the specifications and ensure a compliant feed.
A number of overrides occur to help with iTunes visibility of your episodes.

See the detailed Examples in the GoDocs for complete usage.

### Extensiblity
In no way are you restricted in having full control over your feeds.  You may
choose to skip the API methods and instead use the structs directly.  The
fields have been grouped by RSS 2.0 and iTunes fields.

iTunes specific fields are all prefixed with the letter `I`.

### References
RSS 2.0: <a href="https://cyber.harvard.edu/rss/rss.html">https://cyber.harvard.edu/rss/rss.html</a>

Podcasts: <a href="https://help.apple.com/itc/podcasts_connect/#/itca5b22233">https://help.apple.com/itc/podcasts_connect/#/itca5b22233</a>




## <a name="pkg-index">Index</a>
* [type Author](#Author)
* [type Enclosure](#Enclosure)
* [type EnclosureType](#EnclosureType)
  * [func (et EnclosureType) String() string](#EnclosureType.String)
* [type ICategory](#ICategory)
* [type IImage](#IImage)
* [type Image](#Image)
* [type Item](#Item)
  * [func (i *Item) AddEnclosure(url string, enclosureType EnclosureType, lengthInSeconds int64)](#Item.AddEnclosure)
* [type Podcast](#Podcast)
  * [func New(title, link, description string, pubDate, lastBuildDate *time.Time) Podcast](#New)
  * [func (p *Podcast) AddAuthor(a Author)](#Podcast.AddAuthor)
  * [func (p *Podcast) AddCategory(category string, subCategories []string)](#Podcast.AddCategory)
  * [func (p *Podcast) AddImage(i Image)](#Podcast.AddImage)
  * [func (p *Podcast) AddItem(i Item) (int, error)](#Podcast.AddItem)
  * [func (p *Podcast) Bytes() []byte](#Podcast.Bytes)
  * [func (p *Podcast) Encode(w io.Writer) error](#Podcast.Encode)
  * [func (p *Podcast) String() string](#Podcast.String)
* [type TextInput](#TextInput)

#### <a name="pkg-examples">Examples</a>
* [Package](#example_)
* [New](#example_New)
* [Podcast.AddAuthor](#example_Podcast_AddAuthor)
* [Podcast.AddCategory](#example_Podcast_AddCategory)
* [Podcast.AddImage](#example_Podcast_AddImage)
* [Podcast.AddItem](#example_Podcast_AddItem)
* [Podcast.Bytes](#example_Podcast_Bytes)
* [Package (Encode)](#example__encode)

#### <a name="pkg-files">Package files</a>
[author.go](/src/github.com/eduncan911/podcast/author.go) [doc.go](/src/github.com/eduncan911/podcast/doc.go) [enclosure.go](/src/github.com/eduncan911/podcast/enclosure.go) [image.go](/src/github.com/eduncan911/podcast/image.go) [item.go](/src/github.com/eduncan911/podcast/item.go) [itunes.go](/src/github.com/eduncan911/podcast/itunes.go) [podcast.go](/src/github.com/eduncan911/podcast/podcast.go) [textinput.go](/src/github.com/eduncan911/podcast/textinput.go) 






## <a name="Author">type</a> [Author](/src/target/author.go?s=149:287#L1)
``` go
type Author struct {
    XMLName xml.Name `xml:"itunes:owner"`
    Name    string   `xml:"itunes:name"`
    Email   string   `xml:"itunes:email"`
}
```
Author represents a named author and email.

For iTunes compiance, both Name and Email are required.










## <a name="Enclosure">type</a> [Enclosure](/src/target/enclosure.go?s=814:1118#L36)
``` go
type Enclosure struct {
    XMLName         xml.Name      `xml:"enclosure"`
    URL             string        `xml:"url,attr"`
    Length          int64         `xml:"-"`
    LengthFormatted string        `xml:"length,attr"`
    Type            EnclosureType `xml:"-"`
    TypeFormatted   string        `xml:"type,attr"`
}
```
Enclosure represents a download enclosure.










## <a name="EnclosureType">type</a> [EnclosureType](/src/target/enclosure.go?s=274:296#L11)
``` go
type EnclosureType int
```
EnclosureType specifies the type of the enclosure.


``` go
const (
    M4A EnclosureType = iota
    M4V
    MP4
    MP3
    MOV
    PDF
    EPUB
)
```
EnclosureType specifies the type of the enclosure.










### <a name="EnclosureType.String">func</a> (EnclosureType) [String](/src/target/enclosure.go?s=371:410#L14)
``` go
func (et EnclosureType) String() string
```
String returns the MIME type encoding of the specified EnclosureType.




## <a name="ICategory">type</a> [ICategory](/src/target/itunes.go?s=645:782#L12)
``` go
type ICategory struct {
    XMLName     xml.Name `xml:"itunes:category"`
    Text        string   `xml:"text,attr"`
    ICategories []*ICategory
}
```
ICategory is a 2-tier classification system for iTunes.










## <a name="IImage">type</a> [IImage](/src/target/itunes.go?s=487:584#L6)
``` go
type IImage struct {
    XMLName xml.Name `xml:"itunes:image"`
    HREF    string   `xml:"href,attr"`
}
```
IImage represents an iTunes image.

Podcast feeds contain artwork that is a minimum size of
1400 x 1400 pixels and a maximum size of 3000 x 3000 pixels,
72 dpi, in JPEG or PNG format with appropriate file
extensions (.jpg, .png), and in the RGB colorspace. To optimize
images for mobile devices, Apple recommends compressing your
image files.










## <a name="Image">type</a> [Image](/src/target/image.go?s=398:656#L3)
``` go
type Image struct {
    XMLName xml.Name `xml:"image"`
    // TODO: is it URL or Link? which is it?
    URL    string `xml:"url"`
    Title  string `xml:"title"`
    Link   string `xml:"link"`
    Width  int    `xml:"width,omitempty"`
    Height int    `xml:"height,omitempty"`
}
```
Image represents an image.

Podcast feeds contain artwork that is a minimum size of
1400 x 1400 pixels and a maximum size of 3000 x 3000 pixels,
72 dpi, in JPEG or PNG format with appropriate file
extensions (.jpg, .png), and in the RGB colorspace. To optimize
images for mobile devices, Apple recommends compressing your
image files.










## <a name="Item">type</a> [Item](/src/target/item.go?s=606:1746#L15)
``` go
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
    IAuthor   string `xml:"itunes:author,omitempty"`
    ISubtitle string `xml:"itunes:subtitle,omitempty"`
    // TODO: CDATA
    ISummary           string `xml:"itunes:summary,omitempty"`
    IImage             *IImage
    IDuration          string `xml:"itunes:duration,omitempty"`
    IExplicit          string `xml:"itunes:explicit,omitempty"`
    IIsClosedCaptioned string `xml:"itunes:isClosedCaptioned,omitempty"`
    IOrder             string `xml:"itunes:order,omitempty"`
}
```
Item represents a single entry in a podcast.

Article minimal requirements are:
- Title
- Description
- Link

Audio minimal requirements are:
- Title
- Description
- Enclosure (HREF, Type and Length all required)

Recommendations:
- Setting the minimal fields sets most of other fields, including iTunes.
- Use the Published time.Time setting instead of PubDate.
- Always set an Enclosure.Length, to be nice to your downloaders.
- Use Enclosure.Type instead of setting TypeFormatted for valid extensions.










### <a name="Item.AddEnclosure">func</a> (\*Item) [AddEnclosure](/src/target/item.go?s=1813:1906#L43)
``` go
func (i *Item) AddEnclosure(
    url string, enclosureType EnclosureType, lengthInSeconds int64)
```
AddEnclosure adds the downloadable asset to the podcast Item.




## <a name="Podcast">type</a> [Podcast](/src/target/podcast.go?s=176:1774#L9)
``` go
type Podcast struct {
    XMLName        xml.Name `xml:"channel"`
    Title          string   `xml:"title"`
    Link           string   `xml:"link"`
    Description    string   `xml:"description"`
    Category       string   `xml:"category,omitempty"`
    Cloud          string   `xml:"cloud,omitempty"`
    Copyright      string   `xml:"copyright,omitempty"`
    Docs           string   `xml:"docs,omitempty"`
    Generator      string   `xml:"generator,omitempty"`
    Language       string   `xml:"language,omitempty"`
    LastBuildDate  string   `xml:"lastBuildDate,omitempty"`
    ManagingEditor string   `xml:"managingEditor,omitempty"`
    PubDate        string   `xml:"pubDate,omitempty"`
    Rating         string   `xml:"rating,omitempty"`
    SkipHours      string   `xml:"skipHours,omitempty"`
    SkipDays       string   `xml:"skipDays,omitempty"`
    TTL            int      `xml:"ttl,omitempty"`
    WebMaster      string   `xml:"webMaster,omitempty"`
    Image          *Image
    TextInput      *TextInput

    // https://help.apple.com/itc/podcasts_connect/#/itcb54353390
    IAuthor   string `xml:"itunes:author,omitempty"`
    ISubtitle string `xml:"itunes:subtitle,omitempty"`
    // TODO: CDATA
    ISummary    string `xml:"itunes:summary,omitempty"`
    IBlock      string `xml:"itunes:block,omitempty"`
    IImage      *IImage
    IDuration   string  `xml:"itunes:duration,omitempty"`
    IExplicit   string  `xml:"itunes:explicit,omitempty"`
    IComplete   string  `xml:"itunes:complete,omitempty"`
    INewFeedURL string  `xml:"itunes:new-feed-url,omitempty"`
    IOwner      *Author // Author is formatted for itunes as-is
    ICategories []*ICategory

    Items []*Item
}
```
Podcast represents a podcast.







### <a name="New">func</a> [New](/src/target/podcast.go?s=1940:2025#L52)
``` go
func New(title, link, description string,
    pubDate, lastBuildDate *time.Time) Podcast
```
New instantiates a Podcast with required parameters.

Nil-able fields are optional but recommended as they are formatted
to the expected proper formats.





### <a name="Podcast.AddAuthor">func</a> (\*Podcast) [AddAuthor](/src/target/podcast.go?s=2403:2440#L67)
``` go
func (p *Podcast) AddAuthor(a Author)
```
AddAuthor adds the specified Author to the podcast.




### <a name="Podcast.AddCategory">func</a> (\*Podcast) [AddCategory](/src/target/podcast.go?s=2631:2701#L75)
``` go
func (p *Podcast) AddCategory(category string, subCategories []string)
```
AddCategory adds the cateories to the Podcast in comma delimited format.

subCategories are optional.




### <a name="Podcast.AddImage">func</a> (\*Podcast) [AddImage](/src/target/podcast.go?s=3478:3513#L103)
``` go
func (p *Podcast) AddImage(i Image)
```
AddImage adds the specified Image to the Podcast.

Podcast feeds contain artwork that is a minimum size of
1400 x 1400 pixels and a maximum size of 3000 x 3000 pixels,
72 dpi, in JPEG or PNG format with appropriate file
extensions (.jpg, .png), and in the RGB colorspace. To optimize
images for mobile devices, Apple recommends compressing your
image files.




### <a name="Podcast.AddItem">func</a> (\*Podcast) [AddItem](/src/target/podcast.go?s=4960:5006#L145)
``` go
func (p *Podcast) AddItem(i Item) (int, error)
```
AddItem adds the podcast episode.  It returns a count of Items added or any
errors in validation that may have occurred.

This method takes the "itunes overrides" approach to populating
itunes tags according to the overrides rules in the specification.
This not only complies completely with iTunes parsing rules; but, it also
displays what is possible to be set on an individial eposide level - if you
wish to have more fine grain control over your content.

This method imposes strict validation of the Item being added to confirm
to Podcast and iTunes specifications.

Article minimal requirements are:
* Title
* Description
* Link

Audio, Video and Downloads minimal requirements are:
* Title
* Description
* Enclosure (HREF, Type and Length all required)

The following fields are always overwritten (don't set them):
* GUID
* PubDateFormatted
* AuthorFormatted
* Enclosure.TypeFormatted
* Enclosure.LengthFormatted

Recommendations:
* Just set the minimal fields: the rest get set for you.
* Always set an Enclosure.Length, to be nice to your downloaders.
* Follow Apple's best practices to enrich your podcasts:


	<a href="https://help.apple.com/itc/podcasts_connect/#/itc2b3780e76">https://help.apple.com/itc/podcasts_connect/#/itc2b3780e76</a>

* For specifications of itunes tags, see:


	<a href="https://help.apple.com/itc/podcasts_connect/#/itcb54353390">https://help.apple.com/itc/podcasts_connect/#/itcb54353390</a>




### <a name="Podcast.Bytes">func</a> (\*Podcast) [Bytes](/src/target/podcast.go?s=6829:6861#L213)
``` go
func (p *Podcast) Bytes() []byte
```
Bytes returns an encoded []byte slice.




### <a name="Podcast.Encode">func</a> (\*Podcast) [Encode](/src/target/podcast.go?s=6971:7014#L218)
``` go
func (p *Podcast) Encode(w io.Writer) error
```
Encode writes the bytes to the io.Writer stream in RSS 2.0 specification.




### <a name="Podcast.String">func</a> (\*Podcast) [String](/src/target/podcast.go?s=7091:7124#L223)
``` go
func (p *Podcast) String() string
```
String encodes the Podcast state to a string.




## <a name="TextInput">type</a> [TextInput](/src/target/textinput.go?s=77:290#L1)
``` go
type TextInput struct {
    XMLName     xml.Name `xml:"textInput"`
    Title       string   `xml:"title"`
    Description string   `xml:"description"`
    Name        string   `xml:"name"`
    Link        string   `xml:"link"`
}
```
TextInput represents text inputs.














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
