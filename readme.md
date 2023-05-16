## Web Crawler

This is a very simple webcrawler that crawls a given website and returns a list of all the links found on that website.

### How to build

To build the crawler, simply run the following command:

```bash
make build
```

### How to run

This webcrawler comes in three flavors:

- Recursively: standard recursive crawler without any concurrency
- Fast: concurrent crawler that spaws a goroutine for each link found
- Pooled: concurrent crawler that uses a pool of workers to process the links

```bash
    webcrawler <url> --engine=[recursive|fast|pooled]
```

With pooled engine, you can also specify the number of workers to use, by default is set to the number of cores of the machine:

```bash
    webcrawler <url> --engine=pooled --workers=10
```

### How to test

To run the tests, simply run the following command:

```bash
make test
```

Linter has not been set very strict for simplicity. To run the linter, simply run the following command:

```bash
make lint
```

### Code structure

Code structure is very simple. It is divided in two main packages:

- pkg/lib: contains a safe map to store visited links and the crawlers.
- pkg/crawler: contains the main code of the crawer.
- cmd: contains a small cli interface to run the crawler.

The crawler itself is divided into 4 main components:

- Link parser: parses a link and returns the host, path and document. It also validates that is not a mailto, javascript or anchor.
- Link extractor: extracts all links from a given io.Reader interface.
- Browser: opens a given url and returns the list of links. It is available as webbrowser and diskbrowser (for mocking purposes)
- Crawler: crawls a given url using a browser. It comes in three flavours: recursive, fast and pooled.

### Development process

To implement this webcrawler, I followed the principle of make it work, make it right, make it faster. First I focused on getting a single threaded recursive crawler. Then I focused on test, and finally I focused on make it faster by adding concurrency.

Those are the main tasks in which I divided the development process:
- Create a link parser that is able to detect if a link can be followed or not. It also extracts host, server and document to deal with local links
- Create a link extractor that gets all links from a given document using default golang html parser
- Create a simple recursive crawler that crawls a website and returns a list of links found.
- Create a memory safe map to store the links found and avoid that they are visited twice. It can be implemented with a slice or a map. I created an interface to maybe change it in the future.
- Create a test for the crawler using an interface to mock the webserver.
- Create a concurrent crawler that spawns a goroutine for each link found (make it fast)
- Create a concurrent crawler that uses a pool of workers to process the links (make it right again)
- Wrap everything with a cli interface using cobra
- Clean up the code and add comments

### Possible improvements

Code is never finished. There are always things to improve. Here are some of the things that I would improve if I had more time:

- Remove some repeated code in the crawlers
- Add more tests for more complex sites.
- Handle redirects properly.
- Change the verbosity flag and the if verbosity printf to use a logger with a loglevel. I didn't do it because it needed external dependencies. I wanted to add verbosity because in any production software you need to go deeper if needed.
- Improve some naming, specially the ones related with the extraction of links
- reorganize folders, for example Safevisited initially was intended to be a generic safe map, but then it was much more practical to named it SafeVisited since it just stores visited links. Also, crawler package might be divided into smaller subpackages.
- As any production software, we should add some metrics. I just added a small print statement with the amount of time that took to crawl the website, but we could add some metrics like the amount of links found, the amount of links visited, the amount of links that failed, bytes downloaded, etc.

### External dependencies explained

As requested in the exercise, I was not allowed to use any external dependencies to implement the crawlwer. However, I used some external dependencies that I explain here:
- cobra / viper for cli, because I am really used to them, and it makes really easy to parse arguments
- testify for testing, because it makes really easy to assert results. Its the library I am most comfortable with.
- x/net/html for html parsing, because it is the standard library for html parsing in golang

