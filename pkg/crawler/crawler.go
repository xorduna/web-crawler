package crawler

import (
	"web-crawler/pkg/lib"
)

type Crawler interface {
	Crawl(url string, baseSite string, visitedUrls lib.SafeVisited)
}
