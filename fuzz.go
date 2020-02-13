// +build gofuzz

package podcast

import (
	"bytes"
	"encoding/binary"
	"time"
)

func FuzzItemAddDuration(data []byte) int {
	input, read := binary.Varint(data)
	if input <= 0 && read == 0 {
		// error converting []byte into int64
		return 0
	}
	i := newItem(data)

	i.AddDuration(input)

	p := newPodcast(data)
	if _, err := p.AddItem(i); err != nil {
		return 0
	}
	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzItemAddEnclosure(data []byte) int {
	url := string(data)
	length, read := binary.Varint(data)
	if length <= 0 && read == 0 {
		// error converting []byte into int64
		return 0
	}
	i := newItem(data)

	i.AddEnclosure(url, MP3, length)

	p := newPodcast(data)
	if _, err := p.AddItem(i); err != nil {
		return 0
	}
	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzItemAddImage(data []byte) int {
	i := newItem(data)

	i.AddImage(string(data))

	p := newPodcast(data)
	if _, err := p.AddItem(i); err != nil {
		return 0
	}
	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzItemAddPubDate(data []byte) int {
	t := time.Time{}
	if err := t.GobDecode(data); err != nil {
		return 0
	}
	i := newItem(data)

	i.AddPubDate(&t)

	p := newPodcast(data)
	if _, err := p.AddItem(i); err != nil {
		return 0
	}
	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzItemAddSummary(data []byte) int {
	i := newItem(data)

	i.AddSummary(string(data))

	p := newPodcast(data)
	if _, err := p.AddItem(i); err != nil {
		return 0
	}
	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzPodcastNew(data []byte) int {
	p := newPodcast(data)

	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzPodcastAddAtomLink(data []byte) int {
	p := newPodcast(data)

	p.AddAtomLink(string(data))

	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzPodcastAddAuthor(data []byte) int {
	p := newPodcast(data)

	p.AddAuthor(string(data), string(data))

	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzPodcastAddCategory(data []byte) int {
	p := newPodcast(data)

	subs := make([]string, 3)
	subs[0] = string(data)
	subs[1] = string(data)
	subs[2] = string(data)
	p.AddCategory(string(data), subs)

	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzPodcastAddImage(data []byte) int {
	p := newPodcast(data)

	p.AddImage(string(data))

	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzPodcastAddItem(data []byte) int {
	p := newPodcast(data)
	i := newItem(data)

	if _, err := p.AddItem(i); err != nil {
		return 0
	}

	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzPodcastAddLastBuildDate(data []byte) int {
	p := newPodcast(data)
	t := time.Time{}
	if err := t.GobDecode(data); err != nil {
		return 0
	}

	p.AddLastBuildDate(&t)

	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzPodcastAddPubDate(data []byte) int {
	p := newPodcast(data)
	t := time.Time{}
	if err := t.GobDecode(data); err != nil {
		return 0
	}

	p.AddPubDate(&t)

	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzPodcastAddSubTitle(data []byte) int {
	p := newPodcast(data)

	p.AddSubTitle(string(data))

	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzPodcastAddSummary(data []byte) int {
	p := newPodcast(data)

	p.AddSummary(string(data))

	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func FuzzPodcastBytes(data []byte) int {
	p := newPodcast(data)

	p.Bytes()

	return 1
}

func FuzzPodcastEncode(data []byte) int {
	p := newPodcast(data)

	var buf bytes.Buffer
	if err := p.Encode(&buf); err != nil {
		return 0
	}

	return 1
}

func newPodcast(data []byte) Podcast {
	return New(
		string(data),
		string(data),
		string(data),
		nil, nil)
}

func newItem(data []byte) Item {
	// Article minimal requirements are:
	// - Title
	// - Description
	// - Link
	//
	// Audio minimal requirements are:
	// - Title
	// - Description
	// - Enclosure (HREF, Type and Length all required)
	//
	return Item{
		Title:       string(data),
		Description: string(data),
		Link:        string(data),
	}
}
