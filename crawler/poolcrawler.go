package crawler

import (
	"fmt"
	"sync"
	"web-crawler/lib"
)

func linkCrawl(workerNum int, parentUrl string, url string, visitedUrls lib.SafeVisited, buffer chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	links, err := extractLinksFromUrl(url)
	if err == nil {
		for _, l := range links {

			// check if it is an external link
			if l.Host != parentUrl && l.Host != "" {
				fmt.Printf("worker:%d - Skipping external link: %s\n", workerNum, l.FullLink())
				continue
			}

			//Add this link to the buffer if we haven't visited it before
			if !visitedUrls.IsVisited(l.FullLink()) {
				fmt.Println("Visiting: ", l.FullLink())
				visitedUrls.AddVisited(l.FullLink())
				wg.Add(1)
				link := l.FullLink()
				go func(link string) {
					buffer <- link
				}(link)
			} else {
				fmt.Printf("worker:%d - Already visited: %s\n", workerNum, l.FullLink())
			}
		}
	}
}

type PooledCrawler struct {
	browser    Browser
	numWorkers int
}

func NewPooledCrawler(browser Browser, numWorkers int) *PooledCrawler {
	return &PooledCrawler{
		browser:    browser,
		numWorkers: numWorkers,
	}
}

func (c *PooledCrawler) visit(url string) ([]Link, error) {
	reader, err := c.browser.Get(url)
	if err != nil {
		return nil, err
	}
	return ExtractLinksFromHtml(reader)
}

func (c *PooledCrawler) Crawl(parentUrl string, url string, visitedUrls lib.SafeVisited) {
	linkBuffer := make(chan string)
	wg := &sync.WaitGroup{}
	fmt.Printf("starting %d workers", c.numWorkers)
	workersWg := &sync.WaitGroup{}
	for w := 0; w < c.numWorkers; w++ {
		workersWg.Add(1)
		go func(child int) {
			defer workersWg.Done()
			for link := range linkBuffer {
				fmt.Printf("worker:%d - received link : %s\n", child, link)
				linkCrawl(child, parentUrl, link, visitedUrls, linkBuffer, wg)
			}
			fmt.Printf("worker:%d - received shutdown signal\n", child)
		}(w)
	}

	wg.Add(1)
	linkBuffer <- url

	wg.Wait()
	close(linkBuffer)
	fmt.Println("parent: sent shutdown signal")
	workersWg.Wait()

}
