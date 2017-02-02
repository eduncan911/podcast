package podcast_test

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/eduncan911/podcast"
)

var (
	pubDate = time.Date(2017, time.February, 1, 8, 21, 52, 0, time.Local)
)

func ExampleNew() {
	ti, l, d := "title", "link", "description"

	// instantiate a new Podcast
	p := podcast.New(ti, l, d, &pubDate, &pubDate)

	fmt.Println(p.Title, p.Link, p.Description, p.Language)
	fmt.Println(p.PubDate, p.LastBuildDate)
	// Output:
	// title link description en-us
	// Wed, 01 Feb 2017 08:21:52 -0500 Wed, 01 Feb 2017 08:21:52 -0500
}

func ExamplePodcast_AddAuthor() {
	p := podcast.New("title", "link", "description", nil, nil)

	// add the Author
	p.AddAuthor(podcast.Author{
		Name:  "the name",
		Email: "me@test.com",
	})

	fmt.Println(p.ManagingEditor)
	fmt.Println(p.IAuthor)
	// Output:
	// me@test.com (the name)
	// me@test.com (the name)
}

func ExamplePodcast_AddCategory() {
	p := podcast.New("title", "link", "description", nil, nil)

	// add the Category
	p.AddCategory("Bombay", nil)
	p.AddCategory("American", []string{"Longhair", "Shorthair"})
	p.AddCategory("Siamese", nil)

	fmt.Println(len(p.ICategories), len(p.ICategories[1].ICategories))
	fmt.Println(p.Category)
	// Output:
	// 3 2
	// Bombay,American,Siamese
}

func ExamplePodcast_AddImage() {
	p := podcast.New("title", "link", "description", nil, nil)

	// add the Image
	p.AddImage(podcast.Image{
		URL: "http://example.com/image.jpg",
	})

	if p.Image != nil && p.IImage != nil {
		fmt.Println(p.Image.URL)
		fmt.Println(p.IImage.HREF)
	}
	// Output:
	// http://example.com/image.jpg
	// http://example.com/image.jpg
}

func ExamplePodcast_AddItem() {
	p := podcast.New("title", "link", "description", &pubDate, &pubDate)
	p.AddAuthor(podcast.Author{Name: "the name", Email: "me@test.com"})
	p.AddImage(podcast.Image{URL: "http://example.com/image.jpg"})

	// create an Item
	item := podcast.Item{
		Title:       "Episode 1",
		Description: "Description for Episode 1",
		ISubtitle:   "A simple episode 1",
		PubDate:     &pubDate,
	}
	item.AddEnclosure(
		"http://example.com/1.mp3",
		podcast.MP3,
		183,
	)
	// add the Item
	if _, err := p.AddItem(item); err != nil {
		fmt.Println("item validation error: " + err.Error())
	}

	if len(p.Items) != 1 {
		fmt.Println("expected 1 item in the collection")
	}
	pp := p.Items[0]
	fmt.Println(
		pp.GUID, pp.Title, pp.Link, pp.Description, pp.Author,
		pp.AuthorFormatted, pp.Category, pp.Comments, pp.Source,
		pp.PubDate, pp.PubDateFormatted, *pp.Enclosure,
		pp.IAuthor, pp.IDuration, pp.IExplicit, pp.IIsClosedCaptioned,
		pp.IOrder, pp.ISubtitle, pp.ISummary)
	// Output:
	// http://example.com/1.mp3 Episode 1 http://example.com/1.mp3 Description for Episode 1 &{{ }  me@test.com (the name)}     2017-02-01 08:21:52 -0500 EST Wed, 01 Feb 2017 08:21:52 -0500 {{ } http://example.com/1.mp3 183 183 audio/mpeg audio/mpeg} me@test.com (the name) 183    A simple episode 1
}

func ExamplePodcast_Bytes() {
	pubDate := time.Date(2017, time.February, 1, 9, 11, 0, 0, time.Local)

	p := podcast.New(
		"eduncan911 Podcasts",
		"http://eduncan911.com/",
		"An example Podcast",
		&pubDate, &pubDate,
	)
	p.AddAuthor(podcast.Author{Name: "Jane Doe", Email: "me@janedoe.com"})
	p.AddImage(podcast.Image{URL: "http://janedoe.com/i.jpg"})

	for i := int64(0); i < 4; i++ {
		n := strconv.FormatInt(i, 10)

		item := podcast.Item{
			Title:       "Episode " + n,
			Link:        "http://example.com/" + n + ".mp3",
			Description: "Description for Episode " + n,
			PubDate:     &pubDate,
		}
		if _, err := p.AddItem(item); err != nil {
			fmt.Println(item.Title, ": error", err.Error())
			break
		}
	}

	// call Podcast.Bytes() to return a byte array
	os.Stdout.Write(p.Bytes())

	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
	//   <channel>
	//     <title>eduncan911 Podcasts</title>
	//     <link>http://eduncan911.com/</link>
	//     <description>An example Podcast</description>
	//     <generator>go podcast v1.0.0 (github.com/eduncan911/podcast)</generator>
	//     <language>en-us</language>
	//     <lastBuildDate>Wed, 01 Feb 2017 09:11:00 -0500</lastBuildDate>
	//     <managingEditor>me@janedoe.com (Jane Doe)</managingEditor>
	//     <pubDate>Wed, 01 Feb 2017 09:11:00 -0500</pubDate>
	//     <image>
	//       <url>http://janedoe.com/i.jpg</url>
	//       <title></title>
	//       <link></link>
	//     </image>
	//     <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//     <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//     <item>
	//       <guid>http://example.com/0.mp3</guid>
	//       <title>Episode 0</title>
	//       <link>http://example.com/0.mp3</link>
	//       <description>Description for Episode 0</description>
	//       <pubDate>Wed, 01 Feb 2017 09:11:00 -0500</pubDate>
	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//       <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//     </item>
	//     <item>
	//       <guid>http://example.com/1.mp3</guid>
	//       <title>Episode 1</title>
	//       <link>http://example.com/1.mp3</link>
	//       <description>Description for Episode 1</description>
	//       <pubDate>Wed, 01 Feb 2017 09:11:00 -0500</pubDate>
	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//       <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//     </item>
	//     <item>
	//       <guid>http://example.com/2.mp3</guid>
	//       <title>Episode 2</title>
	//       <link>http://example.com/2.mp3</link>
	//       <description>Description for Episode 2</description>
	//       <pubDate>Wed, 01 Feb 2017 09:11:00 -0500</pubDate>
	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//       <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//     </item>
	//     <item>
	//       <guid>http://example.com/3.mp3</guid>
	//       <title>Episode 3</title>
	//       <link>http://example.com/3.mp3</link>
	//       <description>Description for Episode 3</description>
	//       <pubDate>Wed, 01 Feb 2017 09:11:00 -0500</pubDate>
	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//       <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//     </item>
	//   </channel>
	// </rss>

}
