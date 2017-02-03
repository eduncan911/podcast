package podcast

import "encoding/xml"

// Image represents an image.
//
// Podcast feeds contain artwork that is a minimum size of
// 1400 x 1400 pixels and a maximum size of 3000 x 3000 pixels,
// 72 dpi, in JPEG or PNG format with appropriate file
// extensions (.jpg, .png), and in the RGB colorspace. To optimize
// images for mobile devices, Apple recommends compressing your
// image files.
type Image struct {
	XMLName xml.Name `xml:"image"`
	// TODO: is it URL or Link? which is it?
	URL    string `xml:"url"`
	Title  string `xml:"title,omitempty"`
	Link   string `xml:"link,omitempty"`
	Width  int    `xml:"width,omitempty"`
	Height int    `xml:"height,omitempty"`
}
