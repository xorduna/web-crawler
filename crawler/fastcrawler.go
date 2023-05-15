package crawler

import (
	"fmt"
	"sync"
	"web-crawler/lib"
)

type FastCrawler struct {
	browser   Browser
	verbosity bool
}

func NewFastCrawler(browser Browser, verbosity bool) *FastCrawler {
	return &FastCrawler{
		browser:   browser,
		verbosity: verbosity,
	}
}

func (c *FastCrawler) Crawl(parentURL string, url string, visitedUrls lib.SafeVisited) {
	var wg sync.WaitGroup

	wg.Add(1)
	fastCrawl(c.browser, parentURL, url, visitedUrls, &wg)
	wg.Wait()
}

func fastCrawl(browser Browser, parentURL string, url string, visitedUrls lib.SafeVisited, wg *sync.WaitGroup) {
	defer wg.Done()

	links, err := visit(browser, url)
	if err != nil {
		return
	}

	// first loop is to print the links
	fmt.Printf("Visited: %s\n", url)
	for _, l := range links {
		fmt.Printf(" - %s\n", l.FullLink())
	}

	// second loop is to crawl the links
	for _, l := range links {
		// check if it is an external link
		if l.Host != parentURL && l.Host != "" {
			fmt.Printf("Skipping external link: %s\n", l.FullLink())

			continue
		}

		// check if we have visited this url before
		if !visitedUrls.IsVisited(l.FullLink()) {
			fmt.Println("Visiting: ", l.FullLink())
			visitedUrls.AddVisited(l.FullLink())
			wg.Add(1)

			go fastCrawl(browser, parentURL, l.FullLink(), visitedUrls, wg)
		} else {
			fmt.Printf("Already visited: %s\n", l.FullLink())
		}
	}
}
