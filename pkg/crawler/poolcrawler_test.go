package crawler

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"web-crawler/pkg/lib"
)

func TestPooledCrawler_Crawl(t *testing.T) {

	fakeExtractor := NewDiskBrowser("https://fakesite.com")

	c := NewPooledCrawler(fakeExtractor, 4, false)

	visitedUrls := lib.NewSafeMap()

	c.Crawl("https://fakesite.com", "https://fakesite.com", visitedUrls)

	for _, l := range visitedUrls.List() {
		t.Logf("Visited: %s", l)
		assert.Contains(t, fakeSiteLinks, l)
	}

	assert.Len(t, visitedUrls.List(), len(fakeSiteLinks))
}
