package podcast_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"

	"github.com/eduncan911/podcast"
)

func Example_httpHandlers() {

	// ResponseWriter example using Podcast.Encode(w io.Writer).
	//
	httpHandler := func(w http.ResponseWriter, r *http.Request) {

		// instantiate a new Podcast
		p := podcast.New(
			"eduncan911 Podcasts",
			"http://eduncan911.com/",
			"An example Podcast",
			&pubDate, &updatedDate,
		)

		// add some channel properties
		p.AddAuthor("Jane Doe", "me@janedoe.com")
		p.AddAtomLink("http://eduncan911.com/feed.rss")
		p.AddImage("http://janedoe.com/i.jpg")
		p.AddSummary(`link <a href="http://example.com">example.com</a>`)
		p.IExplicit = "no"

		for i := int64(1); i < 3; i++ {
			n := strconv.FormatInt(i, 10)
			d := pubDate.AddDate(0, 0, int(i))

			// create an Item
			item := podcast.Item{
				Title:       "Episode " + n,
				Link:        "http://example.com/" + n + ".mp3",
				Description: "Description for Episode " + n,
				PubDate:     &d,
			}
			item.AddImage("http://example.com/episode-" + n + ".png")
			item.AddSummary(`item <a href="http://example.com">example.com</a>`)
			// add a Download to the Item
			item.AddEnclosure("http://e.com/"+n+".mp3", podcast.MP3, 55*(i+1))

			// add the Item and check for validation errors
			if _, err := p.AddItem(item); err != nil {
				fmt.Println(item.Title, ": error", err.Error())
				return
			}
		}

		// set the Content Type to that of XML
		w.Header().Set("Content-Type", "application/xml")

		// finally, Encode and write the Podcast to the ResponseWriter.
		//
		// a simple pattern is to handle any errors within this check.
		// alternatively if using middleware, you can just return
		// the Podcast entity as it also implements the io.Writer interface
		// that complies with several middleware packages.
		if err := p.Encode(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	rr := httptest.NewRecorder()
	httpHandler(rr, nil)
	os.Stdout.Write(rr.Body.Bytes())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
	//   <channel>
	//     <title>eduncan911 Podcasts</title>
	//     <link>http://eduncan911.com/</link>
	//     <description>An example Podcast</description>
	//     <generator>go podcast v2.0.0 (github.com/eduncan911/podcast)</generator>
	//     <language>en-us</language>
	//     <lastBuildDate>Mon, 06 Feb 2017 08:21:52 +0000</lastBuildDate>
	//     <managingEditor>me@janedoe.com (Jane Doe)</managingEditor>
	//     <pubDate>Sat, 04 Feb 2017 08:21:52 +0000</pubDate>
	//     <image>
	//       <url>http://janedoe.com/i.jpg</url>
	//       <title>eduncan911 Podcasts</title>
	//       <link>http://eduncan911.com/</link>
	//     </image>
	//     <atom:link href="http://eduncan911.com/feed.rss" rel="self" type="application/rss+xml"></atom:link>
	//     <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//     <itunes:summary><![CDATA[link <a href="http://example.com">example.com</a>]]></itunes:summary>
	//     <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//     <itunes:explicit>no</itunes:explicit>
	//     <item>
	//       <guid>http://e.com/1.mp3</guid>
	//       <title>Episode 1</title>
	//       <link>http://example.com/1.mp3</link>
	//       <description>Description for Episode 1</description>
	//       <pubDate>Sun, 05 Feb 2017 08:21:52 +0000</pubDate>
	//       <enclosure url="http://e.com/1.mp3" length="110" type="audio/mpeg"></enclosure>
	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//       <itunes:summary><![CDATA[item <a href="http://example.com">example.com</a>]]></itunes:summary>
	//       <itunes:image href="http://example.com/episode-1.png"></itunes:image>
	//     </item>
	//     <item>
	//       <guid>http://e.com/2.mp3</guid>
	//       <title>Episode 2</title>
	//       <link>http://example.com/2.mp3</link>
	//       <description>Description for Episode 2</description>
	//       <pubDate>Mon, 06 Feb 2017 08:21:52 +0000</pubDate>
	//       <enclosure url="http://e.com/2.mp3" length="165" type="audio/mpeg"></enclosure>
	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//       <itunes:summary><![CDATA[item <a href="http://example.com">example.com</a>]]></itunes:summary>
	//       <itunes:image href="http://example.com/episode-2.png"></itunes:image>
	//     </item>
	//   </channel>
	// </rss>
}

func Example_ioWriter() {

	// instantiate a new Podcast
	p := podcast.New(
		"Sample Podcasts",
		"http://example.com/",
		"An example Podcast",
		&createdDate, &updatedDate,
	)

	// add some channel properties
	p.ISubtitle = "A simple Podcast"
	p.AddSummary(`link <a href="http://example.com">example.com</a>`)
	p.AddImage("http://example.com/podcast.jpg")
	p.AddAuthor("Jane Doe", "jane.doe@example.com")
	p.AddAtomLink("http://example.com/atom.rss")

	for i := int64(9); i < 11; i++ {
		n := strconv.FormatInt(i, 10)
		d := pubDate.AddDate(0, 0, int(i))

		// create an Item
		item := podcast.Item{
			Title:       "Episode " + n,
			Description: "Description for Episode " + n,
			ISubtitle:   "A simple episode " + n,
			PubDate:     &d,
		}
		item.AddImage("http://example.com/episode-" + n + ".png")
		item.AddSummary(`item k <a href="http://example.com">example.com</a>`)
		// add a Download to the Item
		item.AddEnclosure("http://example.com/"+n+".mp3", podcast.MP3, 55*(i+1))

		// add the Item and check for validation errors
		if _, err := p.AddItem(item); err != nil {
			os.Stderr.WriteString("item validation error: " + err.Error())
		}
	}

	// Podcast.Encode writes to an io.Writer
	if err := p.Encode(os.Stdout); err != nil {
		fmt.Println("error writing to stdout:", err.Error())
	}

	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
	//   <channel>
	//     <title>Sample Podcasts</title>
	//     <link>http://example.com/</link>
	//     <description>An example Podcast</description>
	//     <generator>go podcast v2.0.0 (github.com/eduncan911/podcast)</generator>
	//     <language>en-us</language>
	//     <lastBuildDate>Mon, 06 Feb 2017 08:21:52 +0000</lastBuildDate>
	//     <managingEditor>jane.doe@example.com (Jane Doe)</managingEditor>
	//     <pubDate>Wed, 01 Feb 2017 08:21:52 +0000</pubDate>
	//     <image>
	//       <url>http://example.com/podcast.jpg</url>
	//       <title>Sample Podcasts</title>
	//       <link>http://example.com/</link>
	//     </image>
	//     <atom:link href="http://example.com/atom.rss" rel="self" type="application/rss+xml"></atom:link>
	//     <itunes:author>jane.doe@example.com (Jane Doe)</itunes:author>
	//     <itunes:subtitle>A simple Podcast</itunes:subtitle>
	//     <itunes:summary><![CDATA[link <a href="http://example.com">example.com</a>]]></itunes:summary>
	//     <itunes:image href="http://example.com/podcast.jpg"></itunes:image>
	//     <item>
	//       <guid>http://example.com/9.mp3</guid>
	//       <title>Episode 9</title>
	//       <link>http://example.com/9.mp3</link>
	//       <description>Description for Episode 9</description>
	//       <pubDate>Mon, 13 Feb 2017 08:21:52 +0000</pubDate>
	//       <enclosure url="http://example.com/9.mp3" length="550" type="audio/mpeg"></enclosure>
	//       <itunes:author>jane.doe@example.com (Jane Doe)</itunes:author>
	//       <itunes:subtitle>A simple episode 9</itunes:subtitle>
	//       <itunes:summary><![CDATA[item k <a href="http://example.com">example.com</a>]]></itunes:summary>
	//       <itunes:image href="http://example.com/episode-9.png"></itunes:image>
	//     </item>
	//     <item>
	//       <guid>http://example.com/10.mp3</guid>
	//       <title>Episode 10</title>
	//       <link>http://example.com/10.mp3</link>
	//       <description>Description for Episode 10</description>
	//       <pubDate>Tue, 14 Feb 2017 08:21:52 +0000</pubDate>
	//       <enclosure url="http://example.com/10.mp3" length="605" type="audio/mpeg"></enclosure>
	//       <itunes:author>jane.doe@example.com (Jane Doe)</itunes:author>
	//       <itunes:subtitle>A simple episode 10</itunes:subtitle>
	//       <itunes:summary><![CDATA[item k <a href="http://example.com">example.com</a>]]></itunes:summary>
	//       <itunes:image href="http://example.com/episode-10.png"></itunes:image>
	//     </item>
	//   </channel>
	// </rss>
}
