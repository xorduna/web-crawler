package crawler

import (
	"fmt"
	"web-crawler/pkg/lib"
)

type RecursiveCrawler struct {
	browser   Browser
	verbosity bool
}

func NewCrawler(browser Browser, verbosity bool) *RecursiveCrawler {
	return &RecursiveCrawler{
		browser:   browser,
		verbosity: verbosity,
	}
}

func (c *RecursiveCrawler) Crawl(parentURL string, url string, visited lib.SafeVisited) {
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
		// Skip external links
		if l.Host != parentURL && l.Host != "" {
			if c.verbosity {
				fmt.Printf("Skipping external link: %s\n", l.FullLink())
			}

			continue
		}

		// check if we have visited this url before
		if !visited.IsVisited(l.FullLink()) {
			if c.verbosity {
				fmt.Println("Visiting: ", l.FullLink())
			}

			visited.AddVisited(l.FullLink())

			// If the link is relative, we need to add the parent's host and path
			if l.Host == "" {
				l.Host = parentLink.Host
			}

			if l.Path == "" {
				l.Path = parentLink.Path
			}

			c.Crawl(parentURL, l.FullLink(), visited)
		} else if c.verbosity {
			fmt.Printf("Already visited: %s\n", l.FullLink())
		}
	}
}
