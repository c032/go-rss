package rss_test

import (
	"fmt"
	"os"
	"testing"

	rss "git.wark.io/www/rss-go"
)

const rss2File = "testdata/rss-2.0.rss"

func TestParse(t *testing.T) {
	var (
		err error

		f *os.File
	)

	f, err = os.Open(rss2File)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var feed *rss.RSS

	feed, err = rss.Parse(f)
	if err != nil {
		t.Fatal(err)
	}
	if feed == nil {
		t.Fatalf("nil feed")
	}

	if got, want := feed.Channel.Title, "Kuro Kurori's Lounge"; got != want {
		t.Errorf("feed.Channel.Title = %q; want %q", got, want)
	}
	if got, want := feed.Channel.RawLastBuildDate, "Wed, 25 Dec 2019 16:34:30 +0000"; got != want {
		t.Errorf("feed.Channel.RawLastBuildDate = %q; want %q", got, want)
	}

	tt := []rss.RSSChannelItem{
		{
			Title: "Arisa \u2013 Chapter 39",
			Link:  "https://kurokurori.wordpress.com/2019/12/26/arisa-chapter-39/",
			Categories: []rss.RSSChannelItemCategory{
				{Content: "Disciple Yandere"},
			},
		},
		{
			Title: "Apricot \u2013 Chapter 22",
			Link:  "https://kurokurori.wordpress.com/2019/12/26/apricot-chapter-22/",
			Categories: []rss.RSSChannelItemCategory{
				{Content: "Demon King"},
			},
		},
		{
			Title: "Liselotte \u2013 Chapter 41",
			Link:  "https://kurokurori.wordpress.com/2019/12/26/liselotte-chapter-41/",
			Categories: []rss.RSSChannelItemCategory{
				{Content: "Soon D'rey"},
			},
		},
		{
			Title: "Rena \u2013 Chapter 29",
			Link:  "https://kurokurori.wordpress.com/2019/12/26/rena-chapter-29/",
			Categories: []rss.RSSChannelItemCategory{
				{Content: "Genocide Online"},
			},
		},
		{
			Title: "Saachi \u2013 Chapter 22",
			Link:  "https://kurokurori.wordpress.com/2019/12/26/saachi-chapter-22/",
			Categories: []rss.RSSChannelItemCategory{
				{Content: "Former Assassin"},
			},
		},
	}

	if got, want := len(feed.Channel.Items), len(tt); got == want {
		for i, item := range feed.Channel.Items {
			if got, want := item.Title, tt[i].Title; got != want {
				t.Errorf("feed.Channel.Items[%d].Title = %q; want %q", i, got, want)
			}

			if got, want := item.Link, tt[i].Link; got != want {
				t.Errorf("feed.Channel.Items[%d].Link = %q; want %q", i, got, want)
			}

			if got, want := fmt.Sprintf("%#v", item.Categories), fmt.Sprintf("%#v", tt[i].Categories); got != want {
				t.Errorf("feed.Channel.Items[%d].Categories = %q; want %q", i, got, want)
			}
		}
	} else {
		t.Errorf("len(feed.Channel.Items) = %d; want %d", got, want)
	}
}
