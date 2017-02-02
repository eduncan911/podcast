package podcast

import "encoding/xml"

// EnclosureType specifies the type of the enclosure.
const (
	M4A EnclosureType = iota
	M4V
	MP4
	MP3
	MOV
	PDF
	EPUB
)

const (
	enclosureDefault = "application/octet-stream"
)

// EnclosureType specifies the type of the enclosure.
type EnclosureType int

// String returns the MIME type encoding of the specified EnclosureType.
func (et EnclosureType) String() string {
	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	switch et {
	case M4A:
		return "audio/x-m4a"
	case M4V:
		return "video/x-m4v"
	case MP4:
		return "video/mp4"
	case MP3:
		return "audio/mpeg"
	case MOV:
		return "video/quicktime"
	case PDF:
		return "application/pdf"
	case EPUB:
		return "document/x-epub"
	}
	return enclosureDefault
}

// Enclosure represents a download enclosure.
type Enclosure struct {
	XMLName         xml.Name      `xml:"enclosure"`
	URL             string        `xml:"url,attr"`
	Length          int64         `xml:"-"`
	LengthFormatted string        `xml:"length,attr"`
	Type            EnclosureType `xml:"-"`
	TypeFormatted   string        `xml:"type,attr"`
}
