package podcast

import (
	"encoding/xml"
	"github.com/pkg/errors"
	"strconv"
)

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
	cM4A     = "audio/x-m4a"
	cM4V     = "video/x-m4v"
	cMP4     = "video/mp4"
	cMP3     = "audio/mpeg"
	cMOV     = "video/quicktime"
	cPDF     = "application/pdf"
	cEPUB    = "document/x-epub"
	cDefault = "application/octet-stream"
)

// Enclosure represents a download enclosure.
type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`

	// URL is the downloadable url for the content. (Required)
	URL string `xml:"url,attr"`

	// Length is the size in Bytes of the download. (Required)
	Length EnclosureLength `xml:"length,attr"`

	// Type is MIME type encoding of the download. (Required)
	Type EnclosureType `xml:"type,attr"`
}

// EnclosureLength specifies the length of the enclosure.
type EnclosureLength int64

// MarshalXMLAttr handles the custom formatting from a strongly typed value.
func (et *EnclosureLength) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	v := strconv.FormatInt(int64(*et), 10)
	attr := xml.Attr{
		Name:  name,
		Value: v,
	}
	return attr, nil
}

// UnmarshalXMLAttr handles the custom formatting to a strongly typed value.
func (et *EnclosureLength) UnmarshalXMLAttr(attr xml.Attr) error {
	v, err := strconv.ParseInt(attr.Value, 10, 64)
	if err != nil {
		return errors.Wrap(err, "EnclosureLength UnmarshalXMLAttr could not ParseInt")
	}
	*et = EnclosureLength(v)
	return nil
}

// EnclosureType specifies the type of the enclosure.
type EnclosureType int

// String returns the MIME type encoding of the specified EnclosureType.
func (et EnclosureType) String() string {
	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	switch et {
	case M4A:
		return cM4A
	case M4V:
		return cM4V
	case MP4:
		return cMP4
	case MP3:
		return cMP3
	case MOV:
		return cMOV
	case PDF:
		return cPDF
	case EPUB:
		return cEPUB
	}
	return cDefault
}

// MarshalXMLAttr handles the custom formatting from a strongly typed value.
func (et *EnclosureType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	v := et.String()
	attr := xml.Attr{
		Name:  name,
		Value: v,
	}
	return attr, nil
}

// UnmarshalXMLAttr handles the custom formatting to a strongly typed value.
func (et *EnclosureType) UnmarshalXMLAttr(attr xml.Attr) error {
	v := attr.Value
	switch v {
	case cM4A:
		*et = M4A
	case cM4V:
		*et = M4V
	case cMP4:
		*et = MP4
	case cMP3:
		*et = MP3
	case cMOV:
		*et = MOV
	case cPDF:
		*et = PDF
	case cEPUB:
		*et = EPUB
	default:
		// should we be so strict?
		//
		// if we set a "EnclosureTypeDefault" to the octet-stream it will have
		// to be exported to support the current API that is in use.  to me,
		// that looks sloppy and actually allows people to form invalid feeds by
		// just selecting it for everything.
		//
		// by not allowing a "default enum", like we must do with c# and their
		// parsers, actually keeps the package's API very clean and efficient.
		// but, at the risk of being too strict when parsing existing feeds that
		// may very well be malformed.
		//
		// IF YOU DISAGREE: Please open an Issue and we can discuss, remove,
		//                  change, etc.  because I can easily seeing this
		//                  fail on existing poorly formatted feeds.
		return errors.New("invalid channel.item.enclosure.type found in violation of iTunes specs: " + v)
	}
	return nil
}
