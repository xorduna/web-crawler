package crawler

import (
	"fmt"
	"runtime"
	"sync"
	"web-crawler/lib"
)

func linkCrawl(parentUrl string, url string, visitedUrls lib.SafeVisited, buffer chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	links, err := extractLinksFromUrl(url)
	if err == nil {
		for _, l := range links {

			fmt.Printf("Checking link %s -> %+v\n", l.FullLink(), l)

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
				link := l.FullLink()
				go func(link string) {
					buffer <- link
				}(link)
			} else {
				fmt.Printf("Already visited: %s\n", l.FullLink())
			}
		}
	}
}

func PooledCrawler(parentUrl string, url string, visitedUrls lib.SafeVisited) {
	linkBuffer := make(chan string)
	wg := &sync.WaitGroup{}
	g := runtime.GOMAXPROCS(0)
	fmt.Printf("starting %d workers", g)
	workersWg := &sync.WaitGroup{}
	for c := 0; c < g; c++ {
		workersWg.Add(1)
		go func(child int) {
			defer workersWg.Done()
			for link := range linkBuffer {
				fmt.Printf("worker %d : recv'd link : %s\n", child, link)
				linkCrawl(parentUrl, link, visitedUrls, linkBuffer, wg)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	//linkCrawl("http://www.example.com", url, visitedUrls, linkBuffer, wg)
	wg.Add(1)
	linkBuffer <- url

	wg.Wait()
	close(linkBuffer)
	fmt.Println("parent : sent shutdown signal")
	workersWg.Wait()

	fmt.Println("All buffers ready")
}
