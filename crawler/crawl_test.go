package crawler

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func GenerateFixtureSite() {
	// Crea les carpetes si no existeixen
	err := os.MkdirAll("carpeta1", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll("carpeta2", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll("carpeta3", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Genera 10 fitxers HTML amb enllaços entre ells
	for i := 1; i <= 10; i++ {
		filename := fmt.Sprintf("fixtures/carpeta%d/pagina%d.html", rand.Intn(3)+1, i)
		link := fmt.Sprintf("../carpeta%d/pagina%d.html", rand.Intn(3)+1, (i+1)%10+1)
		content := fmt.Sprintf("<html><body><h1>Pàgina %d</h1><a href=\"%s\">Enllaç a la pàgina següent</a></body></html>", i, link)

		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}

		_, err = file.WriteString(content)
		if err != nil {
			log.Fatal(err)
		}

	}

	fmt.Println("Fitxers HTML generats amb èxit.")
}

func TestCrawl(t *testing.T) {
	// Set up test data
	parentURL := "http://localhost:8080"
	startingURL := "http://localhost:8080/page1"
	visitedURLs := []string{}

	// Start a test HTTP server and define the handler function
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define the response based on the requested URL
		switch r.URL.Path {
		case "/page1":
			// Respond with a page containing links
			fmt.Fprint(w, `
				<a href="/page2">Page 2</a>
				<a href="/page3">Page 3</a>
				<a href="https://external.com">External Link</a>
			`)
		case "/page2":
			// Respond with a page containing links
			fmt.Fprint(w, `
				<a href="/page3">Page 3</a>
				<a href="/page4">Page 4</a>
			`)
		case "/page3":
			// Respond with a page containing links
			fmt.Fprint(w, `
				<a href="/page4">Page 4</a>
			`)
		case "/page4":
			// Respond with a page containing no links
			fmt.Fprint(w, "Page 4")
		default:
			// Respond with a 404 Not Found for any other URLs
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	// Call the function to be tested, using the test server's URL
	Crawl(parentURL, startingURL, &visitedURLs)

	for _, url := range visitedURLs {
		fmt.Printf("Visited URL: %s\n", url)
	}

	// Verify the expected behavior
	expectedVisitedURLs := []string{
		"http://localhost:8080/page2",
		"http://localhost:8080/page3",
		"http://localhost:8080/page4",
	}

	if len(visitedURLs) != len(expectedVisitedURLs) {
		t.Errorf("Unexpected number of visited URLs. Expected: %d, Got: %d", len(expectedVisitedURLs), len(visitedURLs))
	}

	for i, url := range visitedURLs {
		if url != expectedVisitedURLs[i] {
			t.Errorf("Unexpected visited URL. Expected: %s, Got: %s", expectedVisitedURLs[i], url)
		}
	}
}
