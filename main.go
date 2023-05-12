package main

import (
	"fmt"
	"sync"
	"time"
	"web-crawler/crawler"
	"web-crawler/lib"
)

func recursiveCrawl(startUrl string, visitedUrls lib.SafeVisited) {
	crawler.Crawl(startUrl, startUrl, visitedUrls)
}

func poolCrawl(startUrl string, visitedUrls lib.SafeVisited) {
	crawler.PooledCrawler(startUrl, startUrl, visitedUrls)
}

func fastCrawl(startUrl string, visitedUrls lib.SafeVisited) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	crawler.FastCrawl(startUrl, startUrl, visitedUrls, &wg)
	wg.Wait()
}

func main() {
	fmt.Println("Wola!")

	//visitedUrls := []string{}

	//startUrl := "https://parserdigital.com"
	//crawler.Crawl(startUrl, startUrl, &visitedUrls)

	epoch := time.Now()

	visitedUrls := lib.NewSafeMap()
	startUrl := "https://parserdigital.com"

	//recursiveCrawl(startUrl, visitedUrls)
	poolCrawl(startUrl, visitedUrls)
	//fastCrawl(startUrl, visitedUrls)

	fmt.Println(" ====== Visited urls: ======")
	for i, v := range visitedUrls.List() {
		fmt.Printf("%d: %s\n", i, v)
	}

	fmt.Printf("Time taken: %s\n", time.Since(epoch))

}
