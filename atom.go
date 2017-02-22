package podcast

import "encoding/xml"

// AtomLink represents the Atom reference link.
type AtomLink struct {
	XMLName xml.Name `xml:"atom:link"`
	HREF    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}
