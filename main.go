package main

import (
	"fmt"
	"log"
	"net/url"
	"runtime"
	"time"
	"web-crawler/crawler"
	"web-crawler/lib"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//nolint:lll
func recursiveCrawl(parentURL string, startURL string, visitedUrls lib.SafeVisited, browser crawler.Browser, verbosity bool) {
	c := crawler.NewCrawler(browser, verbosity)
	c.Crawl(parentURL, startURL, visitedUrls)
}

//nolint:lll
func poolCrawl(parentURL string, startURL string, visitedUrls lib.SafeVisited, browser crawler.Browser, verbosity bool, workers int) {
	c := crawler.NewPooledCrawler(browser, workers, verbosity)
	c.Crawl(parentURL, startURL, visitedUrls)
}

//nolint:lll
func fastCrawl(parentURL string, startURL string, visitedUrls lib.SafeVisited, browser crawler.Browser, verbosity bool) {
	crawler := crawler.NewFastCrawler(browser, verbosity)
	crawler.Crawl(parentURL, startURL, visitedUrls)
}

var (
	engine    string
	workers   int
	verbosity bool
)

func runCrawler(cmd *cobra.Command, args []string) {
	inputURL := args[0]

	// Validate URL
	parsedURL, err := url.ParseRequestURI(inputURL)
	if err != nil {
		log.Fatalf("Invalid URL: %s", inputURL)
	}

	// Check if URL scheme is specified
	if parsedURL.Scheme == "" {
		log.Fatalf("Invalid URL: Scheme not specified")
	}

	parentURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)
	fmt.Printf("Parent URL: %s\n", parentURL)

	epoch := time.Now()
	visitedUrls := lib.NewSafeMap()

	browser := crawler.NewWebBrowser()

	switch engine {
	case "recursive":
		recursiveCrawl(parentURL, inputURL, visitedUrls, browser, verbosity)
	case "fast":
		fastCrawl(parentURL, inputURL, visitedUrls, browser, verbosity)
	case "pooled":
		poolCrawl(parentURL, inputURL, visitedUrls, browser, verbosity, workers)
	default:
		log.Fatal("Invalid engine specified")
	}

	fmt.Println(" ====== Visited urls: ======")

	for i, v := range visitedUrls.List() {
		fmt.Printf("%d: %s\n", i, v)
	}

	fmt.Printf("Time taken: %s\n", time.Since(epoch))
}

//nolint:errcheck
func main() {
	rootCmd := &cobra.Command{
		Use:   "crawler [url]",
		Short: "A CLI for running crawlers",
		Args:  cobra.ExactArgs(1),
		Run:   runCrawler,
	}

	rootCmd.Flags().StringVarP(&engine, "engine", "e", "recursive",
		"Crawling engine (recursive, fast, pooled)")
	viper.BindPFlags(rootCmd.Flags())

	rootCmd.Flags().IntVarP(&workers, "workers", "w", runtime.GOMAXPROCS(0),
		"Number of workers (only applicable for pooled engine)")
	viper.BindPFlag("workers", rootCmd.Flags().Lookup("workers"))

	rootCmd.Flags().BoolVarP(&verbosity, "verbose", "v", false,
		"Verbose output")
	viper.BindPFlag("verbose", rootCmd.Flags().Lookup("verbose"))

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
