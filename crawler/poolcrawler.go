package crawler

import (
	"fmt"
	"sync"
	"web-crawler/lib"
)

type PooledCrawler struct {
	browser    Browser
	numWorkers int
	verbosity  bool
}

func NewPooledCrawler(browser Browser, numWorkers int, verbosity bool) *PooledCrawler {
	return &PooledCrawler{
		browser:    browser,
		numWorkers: numWorkers,
		verbosity:  verbosity,
	}
}

func (c *PooledCrawler) Crawl(parentURL string, url string, visitedUrls lib.SafeVisited) {
	linkBuffer := make(chan string)
	wg := &sync.WaitGroup{}
	workersWg := &sync.WaitGroup{}

	if c.verbosity {
		fmt.Printf("starting %d workers", c.numWorkers)
	}

	for w := 0; w < c.numWorkers; w++ {
		workersWg.Add(1)

		workerFunc := func(child int) {
			defer workersWg.Done()

			for link := range linkBuffer {
				if c.verbosity {
					fmt.Printf("worker:%d - received link : %s\n", child, link)
				}

				c.linkCrawl(child, parentURL, link, visitedUrls, linkBuffer, wg)
			}

			if c.verbosity {
				fmt.Printf("worker:%d - received shutdown signal\n", child)
			}
		}

		go workerFunc(w)
	}

	wg.Add(1)
	linkBuffer <- url

	wg.Wait()
	close(linkBuffer)
	if c.verbosity {
		fmt.Println("parent: sent shutdown signal")
	}

	workersWg.Wait()
}

//nolint:lll
func (c *PooledCrawler) linkCrawl(workerNum int, parentURL string, url string, visitedUrls lib.SafeVisited, buffer chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	links, err := visit(c.browser, url)
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
				fmt.Printf("worker:%d - Skipping external link: %s\n", workerNum, l.FullLink())
			}

			continue
		}

		// Add this link to the buffer if we haven't visited it before
		if !visitedUrls.IsVisited(l.FullLink()) {
			if c.verbosity {
				fmt.Println("Visiting: ", l.FullLink())
			}
			visitedUrls.AddVisited(l.FullLink())
			wg.Add(1)

			link := l.FullLink()
			go func(link string) {
				buffer <- link
			}(link)
		} else {
			if c.verbosity {
				fmt.Printf("worker:%d - Already visited: %s\n", workerNum, l.FullLink())
			}
		}
	}
}
