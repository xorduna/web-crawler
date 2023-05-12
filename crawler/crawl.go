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

func Crawl(parentUrl string, url string, visited lib.SafeVisited) {

	links, err := extractLinksFromUrl(url)
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
				Crawl(parentUrl, l.FullLink(), visited)

			} else {
				fmt.Printf("Already visited: %s\n", l.FullLink())
			}
		}
	}
}
