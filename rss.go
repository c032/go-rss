package rss

import (
	"encoding/xml"
	"io"
)

type RSS struct {
	Version string     `xml:"version,attr"`
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	Title            string           `xml:"title"`
	Description      string           `xml:"description"`
	Link             string           `xml:"link"`
	RawLastBuildDate string           `xml:"lastBuildDate"`
	RawPubDate       string           `xml:"pubDate"`
	RawTTL           string           `xml:"ttl"`
	Items            []RSSChannelItem `xml:"item"`
}

type RSSChannelItem struct {
	Title       string                   `xml:"title"`
	Description string                   `xml:"description"`
	Link        string                   `xml:"link"`
	GUID        RSSChannelItemGUID       `xml:"guid"`
	RawPubDate  string                   `xml:"pubDate"`
	Categories  []RSSChannelItemCategory `xml:"category"`
}

type RSSChannelItemCategory struct {
	Content string `xml:",chardata"`
}

type RSSChannelItemGUID struct {
	Content        string `xml:",chardata"`
	RawIsPermaLink string `xml:"isPermaLink,attr"`
}

func Parse(r io.Reader) (*RSS, error) {
	d := xml.NewDecoder(r)

	feed := &RSS{}

	err := d.Decode(&feed)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
