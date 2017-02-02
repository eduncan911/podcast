package podcast

import "encoding/xml"

// Speicfication: https://help.apple.com/itc/podcasts_connect/#/itcb54353390
//

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

// ICategory is a 2-tier classification system for iTunes.
type ICategory struct {
	XMLName     xml.Name `xml:"itunes:category"`
	Text        string   `xml:"text,attr"`
	ICategories []*ICategory
}
