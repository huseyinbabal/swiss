package rss

import (
	"strings"
	"testing"
)

func TestToJSON(t *testing.T) {
	got, err := ToJSON(`<rss><channel><title>Feed</title></channel></rss>`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, `"title": "Feed"`) {
		t.Errorf("got %q", got)
	}
}
