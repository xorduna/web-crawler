package crawler

import (
	"fmt"
	"sync"
	"web-crawler/lib"
)

func FastCrawl(parentUrl string, url string, visitedUrls lib.SafeVisited, wg *sync.WaitGroup) {
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
				go FastCrawl(parentUrl, l.FullLink(), visitedUrls, wg)

			} else {
				fmt.Printf("Already visited: %s\n", l.FullLink())
			}
		}
	}
}
