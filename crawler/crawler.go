package crawler

import "web-crawler/lib"

type Crawler interface {
	Crawl(url string, baseSite string, visitedUrls lib.SafeVisited)
}

func visit(browser Browser, url string) ([]Link, error) {
	reader, err := browser.Get(url)
	if err != nil {
		return nil, err
	}

	return ExtractLinksFromHTML(reader)
}
