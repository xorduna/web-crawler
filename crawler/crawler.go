package crawler

import "web-crawler/lib"

type Crawler interface {
	Crawl(url string, baseSite string, visitedUrls lib.SafeVisited)
}
