package podcast_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"time"

	"github.com/eduncan911/podcast"
)

func Example() {
	now := time.Date(2017, time.February, 1, 7, 51, 0, 0, time.Local)

	// instantiate a new Podcast
	p := podcast.New(
		"Sample Podcasts",
		"http://example.com/",
		"An example Podcast",
		&now, &now,
	)

	// add some channel properties
	p.ISubtitle = "A simple Podcast"
	p.AddImage(podcast.Image{URL: "http://example.com/podcast.jpg"})
	p.AddAuthor(podcast.Author{
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
	})

	for i := int64(0); i < 2; i++ {
		n := strconv.FormatInt(i, 10)

		// create an Item
		item := podcast.Item{
			Title:       "Episode " + n,
			Description: "Description for Episode " + n,
			ISubtitle:   "A simple episode " + n,
			PubDate:     &now,
		}
		// add a Download to the Item
		item.AddEnclosure(
			"http://example.com/"+n+".mp3", podcast.MP3, 55*(i+1))

		// add the Item and check for validation errors
		if _, err := p.AddItem(item); err != nil {
			os.Stderr.WriteString("item validation error: " + err.Error())
		}
	}

	// Podcast.Encode writes to an io.Writer
	if err := p.Encode(os.Stdout); err != nil {
		os.Stderr.WriteString("error writing to stdout: " + err.Error())
	}

	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
	//   <channel>
	//     <title>Sample Podcasts</title>
	//     <link>http://example.com/</link>
	//     <description>An example Podcast</description>
	//     <generator>go podcast v1.0.0 (github.com/eduncan911/podcast)</generator>
	//     <language>en-us</language>
	//     <lastBuildDate>Wed, 01 Feb 2017 07:51:00 -0500</lastBuildDate>
	//     <managingEditor>jane.doe@example.com (Jane Doe)</managingEditor>
	//     <pubDate>Wed, 01 Feb 2017 07:51:00 -0500</pubDate>
	//     <image>
	//       <url>http://example.com/podcast.jpg</url>
	//       <title></title>
	//       <link></link>
	//     </image>
	//     <itunes:author>jane.doe@example.com (Jane Doe)</itunes:author>
	//     <itunes:subtitle>A simple Podcast</itunes:subtitle>
	//     <itunes:image href="http://example.com/podcast.jpg"></itunes:image>
	//     <item>
	//       <guid>http://example.com/0.mp3</guid>
	//       <title>Episode 0</title>
	//       <link>http://example.com/0.mp3</link>
	//       <description>Description for Episode 0</description>
	//       <pubDate>Wed, 01 Feb 2017 07:51:00 -0500</pubDate>
	//       <enclosure url="http://example.com/0.mp3" length="55" type="audio/mpeg"></enclosure>
	//       <itunes:author>jane.doe@example.com (Jane Doe)</itunes:author>
	//       <itunes:subtitle>A simple episode 0</itunes:subtitle>
	//       <itunes:image href="http://example.com/podcast.jpg"></itunes:image>
	//       <itunes:duration>55</itunes:duration>
	//     </item>
	//     <item>
	//       <guid>http://example.com/1.mp3</guid>
	//       <title>Episode 1</title>
	//       <link>http://example.com/1.mp3</link>
	//       <description>Description for Episode 1</description>
	//       <pubDate>Wed, 01 Feb 2017 07:51:00 -0500</pubDate>
	//       <enclosure url="http://example.com/1.mp3" length="110" type="audio/mpeg"></enclosure>
	//       <itunes:author>jane.doe@example.com (Jane Doe)</itunes:author>
	//       <itunes:subtitle>A simple episode 1</itunes:subtitle>
	//       <itunes:image href="http://example.com/podcast.jpg"></itunes:image>
	//       <itunes:duration>110</itunes:duration>
	//     </item>
	//   </channel>
	// </rss>
}

func Example_httpHandlers() {

	// ResponseWriter example using Podcast.Encode(w io.Writer).
	//
	httpHandler := func(w http.ResponseWriter, r *http.Request) {

		// instantiate a new Podcast
		p := podcast.New(
			"eduncan911 Podcasts",
			"http://eduncan911.com/",
			"An example Podcast",
			&pubDate, &pubDate,
		)

		// add some channel properties
		p.AddAuthor(podcast.Author{Name: "Jane Doe", Email: "me@janedoe.com"})
		p.AddImage(podcast.Image{URL: "http://janedoe.com/i.jpg"})

		for i := int64(0); i < 3; i++ {
			n := strconv.FormatInt(i, 10)

			// create an Item
			item := podcast.Item{
				Title:       "Episode " + n,
				Link:        "http://example.com/" + n + ".mp3",
				Description: "Description for Episode " + n,
				PubDate:     &pubDate,
			}
			// add a Download to the Item
			item.AddEnclosure(
				"http://example.com/"+n+".mp3", podcast.MP3, 55*(i+1))

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
	// <rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
	//   <channel>
	//     <title>eduncan911 Podcasts</title>
	//     <link>http://eduncan911.com/</link>
	//     <description>An example Podcast</description>
	//     <generator>go podcast v1.0.0 (github.com/eduncan911/podcast)</generator>
	//     <language>en-us</language>
	//     <lastBuildDate>Wed, 01 Feb 2017 08:21:52 -0500</lastBuildDate>
	//     <managingEditor>me@janedoe.com (Jane Doe)</managingEditor>
	//     <pubDate>Wed, 01 Feb 2017 08:21:52 -0500</pubDate>
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
	//       <pubDate>Wed, 01 Feb 2017 08:21:52 -0500</pubDate>
	//       <enclosure url="http://example.com/0.mp3" length="55" type="audio/mpeg"></enclosure>
	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//       <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//       <itunes:duration>55</itunes:duration>
	//     </item>
	//     <item>
	//       <guid>http://example.com/1.mp3</guid>
	//       <title>Episode 1</title>
	//       <link>http://example.com/1.mp3</link>
	//       <description>Description for Episode 1</description>
	//       <pubDate>Wed, 01 Feb 2017 08:21:52 -0500</pubDate>
	//       <enclosure url="http://example.com/1.mp3" length="110" type="audio/mpeg"></enclosure>
	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//       <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//       <itunes:duration>110</itunes:duration>
	//     </item>
	//     <item>
	//       <guid>http://example.com/2.mp3</guid>
	//       <title>Episode 2</title>
	//       <link>http://example.com/2.mp3</link>
	//       <description>Description for Episode 2</description>
	//       <pubDate>Wed, 01 Feb 2017 08:21:52 -0500</pubDate>
	//       <enclosure url="http://example.com/2.mp3" length="165" type="audio/mpeg"></enclosure>
	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//       <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//       <itunes:duration>165</itunes:duration>
	//     </item>
	//   </channel>
	// </rss>

}
