package crawler

import (
	"testing"
	"web-crawler/pkg/lib"
)

func TestCrawl(t *testing.T) {

	fakeExtractor := NewDiskBrowser("https://fakesite.com")

	c := NewCrawler(fakeExtractor, false)

	visitedUrls := lib.NewSafeMap()

	c.Crawl("https://fakesite.com", "https://fakesite.com", visitedUrls)

	for _, l := range visitedUrls.List() {
		t.Logf("Visited: %s", l)
	}
}
