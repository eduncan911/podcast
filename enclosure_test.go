package podcast_test

import (
	"github.com/eduncan911/podcast"
	"github.com/stretchr/testify/assert"
	"testing"
)

type enclosureTest struct {
	t        podcast.EnclosureType
	expected string
}

var enclosureTests = []enclosureTest{
	enclosureTest{podcast.M4A, "audio/x-m4a"},
	enclosureTest{podcast.M4V, "video/x-m4v"},
	enclosureTest{podcast.MP4, "video/mp4"},
	enclosureTest{podcast.MP3, "audio/mpeg"},
	enclosureTest{podcast.MOV, "video/quicktime"},
	enclosureTest{podcast.PDF, "application/pdf"},
	enclosureTest{podcast.EPUB, "document/x-epub"},
	enclosureTest{podcast.M4A, "audio/x-m4a"},
	enclosureTest{99, "application/octet-stream"},
}

func TestEnclosureTypes(t *testing.T) {
	t.Parallel()
	for _, et := range enclosureTests {
		et := et
		t.Run(et.t.String(), func(t *testing.T) {
			t.Parallel()

			assert.EqualValues(t, et.expected, et.t.String())
		})
	}
}
