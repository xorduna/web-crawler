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
	c.fastCrawl(parentURL, url, visitedUrls, &wg)
	wg.Wait()
}

func (c *FastCrawler) fastCrawl(parentURL string, url string, visitedUrls lib.SafeVisited, wg *sync.WaitGroup) {
	defer wg.Done()

	parentLink, err := ParseLink(url)
	if err != nil {
		return
	}

	links, err := c.browser.Get(url)
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
			if c.verbosity {
				fmt.Printf("Skipping external link: %s\n", l.FullLink())
			}

			continue
		}

		// check if we have visited this url before
		if !visitedUrls.IsVisited(l.FullLink()) {
			if c.verbosity {
				fmt.Println("Visiting: ", l.FullLink())
			}

			visitedUrls.AddVisited(l.FullLink())
			wg.Add(1)

			// If the link is relative, we need to add the parent's host and path
			if l.Host == "" {
				l.Host = parentLink.Host
			}

			if l.Path == "" {
				l.Path = parentLink.Path
			}

			go c.fastCrawl(parentURL, l.FullLink(), visitedUrls, wg)
		} else if c.verbosity {
			fmt.Printf("Already visited: %s\n", l.FullLink())
		}
	}
}
