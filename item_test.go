package podcast_test

import (
	"testing"

	"github.com/eduncan911/podcast"
	"github.com/stretchr/testify/assert"
)

func TestItemAddSummaryTooLong(t *testing.T) {
	t.Parallel()

	// arrange
	i := podcast.Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}
	summary := ""
	for {
		if len(summary) >= 4051 {
			break
		}
		summary = summary + "abc ss 5 "
	}

	// act
	i.AddSummary(summary)

	// assert
	assert.Len(t, i.ISummary.Text, 4000)
}

func TestItemAddImageEmptyUrl(t *testing.T) {
	t.Parallel()

	// arrange
	i := podcast.Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}

	// act
	i.AddImage("")

	// assert
	assert.Nil(t, i.IImage)
}
