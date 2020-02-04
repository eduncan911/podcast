package podcast

import "encoding/xml"

// Specifications: https://help.apple.com/itc/podcasts_connect/#/itcb54353390
//

// ICategory is a 2-tier classification system for iTunes.
type ICategory struct {
	XMLName     xml.Name `xml:"itunes:category"`
	Text        string   `xml:"text,attr"`
	ICategories []*ICategory
}

// IImage represents an iTunes image.
//
// Podcast feeds contain artwork that is a minimum size of
// 1400 x 1400 pixels and a maximum size of 3000 x 3000 pixels,
// 72 dpi, in JPEG or PNG format with appropriate file
// extensions (.jpg, .png), and in the RGB colorspace. To optimize
// images for mobile devices, Apple recommends compressing your
// image files.
type IImage struct {
	XMLName xml.Name `xml:"itunes:image"`
	HREF    string   `xml:"href,attr"`
}

// ISummary is a 4000 character rich-text field for the itunes:summary tag.
//
// This is rendered as CDATA which allows for HTML tags such as `<a href="">`.
type ISummary struct {
	XMLName xml.Name `xml:"itunes:summary"`
	Text    string   `xml:",cdata"`
}
