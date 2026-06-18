package rss

import "swiss/internal/xml"

// ToJSON converts an RSS feed (which is XML) into JSON.
func ToJSON(in string) (string, error) {
	return xml.ToJSON(in)
}
