package podcast

import "encoding/xml"

// Description is a 4000 character rich-text field for the channel and
// podcast description tags.
//
// This is rendered as CDATA which allows for HTML tags such as `<a href="">`.
type Description struct {
	XMLName xml.Name `xml:"description"`
	Text    string   `xml:",cdata"`
}
