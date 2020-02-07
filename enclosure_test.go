package podcast_test

import (
	"testing"

	"github.com/eduncan911/podcast"
	"github.com/stretchr/testify/assert"
)

type enclosureTest struct {
	t        podcast.EnclosureType
	expected string
}

var enclosureTests = []enclosureTest{
	{podcast.M4A, "audio/x-m4a"},
	{podcast.M4V, "video/x-m4v"},
	{podcast.MP4, "video/mp4"},
	{podcast.MP3, "audio/mpeg"},
	{podcast.MOV, "video/quicktime"},
	{podcast.PDF, "application/pdf"},
	{podcast.EPUB, "document/x-epub"},
	{podcast.M4A, "audio/x-m4a"},
	{99, "application/octet-stream"},
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
