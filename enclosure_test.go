package podcast_test

import (
	"github.com/eduncan911/podcast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnclosureTypeM4A(t *testing.T) {
	t.Parallel()

	// act
	v := podcast.M4A.String()

	// assert
	assert.EqualValues(t, "audio/x-m4a", v)
}

func TestEnclosureTypeM4V(t *testing.T) {
	t.Parallel()

	// act
	v := podcast.M4V.String()

	// assert
	assert.EqualValues(t, "video/x-m4v", v)
}

func TestEnclosureTypeMP4(t *testing.T) {
	t.Parallel()

	// act
	v := podcast.MP4.String()

	// assert
	assert.EqualValues(t, "video/mp4", v)
}

func TestEnclosureTypeMP3(t *testing.T) {
	t.Parallel()

	// act
	v := podcast.MP3.String()

	// assert
	assert.EqualValues(t, "audio/mpeg", v)
}

func TestEnclosureTypeMOV(t *testing.T) {
	t.Parallel()

	// act
	v := podcast.MOV.String()

	// assert
	assert.EqualValues(t, "video/quicktime", v)
}

func TestEnclosureTypePDF(t *testing.T) {
	t.Parallel()

	// act
	v := podcast.PDF.String()

	// assert
	assert.EqualValues(t, "application/pdf", v)
}

func TestEnclosureTypeEPUB(t *testing.T) {
	t.Parallel()

	// act
	v := podcast.EPUB.String()

	// assert
	assert.EqualValues(t, "document/x-epub", v)
}

func TestEnclosureTypeDefault(t *testing.T) {
	t.Parallel()

	// act
	v := podcast.EnclosureType(99)

	// assert
	assert.EqualValues(t, "application/octet-stream", v.String())
}
