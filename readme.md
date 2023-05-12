## Web Crawler

This is a very simple webcrawler that crawls a given website and returns a list of all the links found on that website.

### How to run

To run the crawler, simply run the following command:

```bash
go run main.go
```

It runs in three flavors:

- Recursively: standard recursive crawler without any concurrency
- Fast: concurrent crawler that spaws a goroutine for each link found
- Pooled: concurrent crawler that uses a pool of workers to process the links

### Development process

Those are the main tasks in which i divided the development process:

- Create a link parser that is able to detect wether a link can be followed or not. It also extracts host, server and document to deal with local links
- Create a link extractor that gets all links from a given document using default golang html parser
- Create a memory safe map to store the links found and avoid that they are visited twice. It can be implemented with a slice or a map. I created an interface to maybe change it in the future.
- Create a simple recursive crawler that crawls a website and returns a list of links found.
- Create a test for the crawler using an interface to mock the webserver. This site was generated automatically using a small go program.
- Create a concurrent crawler that spawns a goroutine for each link found
- Create a concurrent crawler that uses a pool of workers to process the links
- Clean up the code and add comments