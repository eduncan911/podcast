package podcast

import "encoding/xml"

// Atom represents the Atom reference link.
type Atom struct {
	XMLName xml.Name `xml:"atom:link"`
	HREF    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}
