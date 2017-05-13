package podcast_test

import (
	"fmt"

	"github.com/eduncan911/podcast"
)

func ExampleItem_AddPubDate() {
	p := podcast.New("title", "link", "description", nil, nil)
	i := podcast.Item{
		Title:       "item title",
		Description: "item desc",
		Link:        "item link",
	}
	d := pubDate.AddDate(0, 0, -11)

	// add the pub date
	i.AddPubDate(&d)

	if i.PubDate != nil {
		fmt.Println(i.PubDateFormatted, *i.PubDate)
	}
	p.AddItem(i) // this should not override with Podcast.PubDate
	fmt.Println(i.PubDateFormatted, *i.PubDate)
	// Output:
	// Tue, 24 Jan 2017 08:21:52 +0000 2017-01-24 08:21:52 +0000 UTC
	// Tue, 24 Jan 2017 08:21:52 +0000 2017-01-24 08:21:52 +0000 UTC
}

func ExampleItem_AddDuration() {
	i := podcast.Item{
		Title:       "item title",
		Description: "item desc",
		Link:        "item link",
	}
	d := int64(533)

	// add the Duration in Seconds
	i.AddDuration(d)

	fmt.Println(i.IDuration)
	// Output:
	// 533
}
