package crawler

import (
	"fmt"
	"sync"
	"web-crawler/lib"
)

type FastCrawler struct {
	browser Browser
}

func NewFastCrawler(browser Browser) *FastCrawler {
	return &FastCrawler{
		browser: browser,
	}
}

func (c *FastCrawler) visit(url string) ([]Link, error) {
	reader, err := c.browser.Get(url)
	if err != nil {
		return nil, err
	}
	return ExtractLinksFromHtml(reader)
}

func (c *FastCrawler) Crawl(parentUrl string, url string, visitedUrls lib.SafeVisited) {
	var wg sync.WaitGroup
	wg.Add(1)
	fastCrawl(parentUrl, url, visitedUrls, &wg)
	wg.Wait()
}

func fastCrawl(parentUrl string, url string, visitedUrls lib.SafeVisited, wg *sync.WaitGroup) {
	defer wg.Done()
	links, err := extractLinksFromUrl(url)
	if err == nil {
		for _, l := range links {

			// check if it is an external link
			if l.Host != parentUrl && l.Host != "" {
				fmt.Printf("Skipping external link: %s\n", l.FullLink())
				continue
			}

			//check if we have visited this url before
			if !visitedUrls.IsVisited(l.FullLink()) {
				fmt.Println("Visiting: ", l.FullLink())
				visitedUrls.AddVisited(l.FullLink())
				wg.Add(1)
				go fastCrawl(parentUrl, l.FullLink(), visitedUrls, wg)

			} else {
				fmt.Printf("Already visited: %s\n", l.FullLink())
			}
		}
	}
}
