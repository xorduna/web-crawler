package crawler

import (
	"fmt"
	"net/http"
	"web-crawler/lib"
)

func isInArray(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}

func extractLinksFromUrl(url string) ([]Link, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ExtractLinksFromHtml(resp.Body)
}

type RecursiveCrawler struct {
	browser Browser
}

func NewCrawler(browser Browser) *RecursiveCrawler {
	return &RecursiveCrawler{
		browser: browser,
	}
}

func (c *RecursiveCrawler) visit(url string) ([]Link, error) {
	reader, err := c.browser.Get(url)
	if err != nil {
		return nil, err
	}
	return ExtractLinksFromHtml(reader)
}

func (c *RecursiveCrawler) Crawl(parentUrl string, url string, visited lib.SafeVisited) {
	parentLink, err := ParseLink(url)
	if err != nil {
		return
	}

	links, err := c.visit(url)
	if err == nil {
		for _, l := range links {
			//check if we have visited this url before
			if l.Host != parentUrl && l.Host != "" {
				fmt.Printf("Skipping external link: %s\n", l.FullLink())
				continue
			}
			if !visited.IsVisited(l.FullLink()) {
				fmt.Println("Visiting: ", l.FullLink())
				visited.AddVisited(l.FullLink())

				if l.Host == "" {
					l.Host = parentLink.Host
				}
				if l.Path == "" {
					l.Path = parentLink.Path
				}
				c.Crawl(parentUrl, l.FullLink(), visited)

			} else {
				fmt.Printf("Already visited: %s\n", l.FullLink())
			}
		}
	}
}
