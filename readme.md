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

### Development process

To implement this webcrawler, I followed the principle of make it work, make it right, make it faster. First i focused on getting a single threaded recursive crawler. Then I focused on test, and finally I focused on make it faster by adding concurrency.

Those are the main tasks in which I divided the development process:
- Create a link parser that is able to detect if a link can be followed or not. It also extracts host, server and document to deal with local links
- Create a link extractor that gets all links from a given document using default golang html parser
- Create a simple recursive crawler that crawls a website and returns a list of links found.
- Create a memory safe map to store the links found and avoid that they are visited twice. It can be implemented with a slice or a map. I created an interface to maybe change it in the future.
- Create a test for the crawler using an interface to mock the webserver.
- Create a concurrent crawler that spawns a goroutine for each link found
- Create a concurrent crawler that uses a pool of workers to process the links
- Wrap everything with a cli interface using cobra
- Clean up the code and add comments

### Possible improvements

- Remove some repeated code in the crawlers
- Add more tests for more complex sites
- handle redirects properly
- ensure that body is closed when opening the url. Before it was closed on the same function, but since I added the DiskBrowser to mock the webserver I had to close it somewhere else and io.Reader does not provide a Close method.
- change the verbosity flag and the if verbosity printf to use a logger with a loglevel. I didn't do it because it needed external dependencies. I wanted to add verbosity because in any production software you need to go deeper if needed.

### External dependencies explained

As requested in the exercise, I was not allowed to use any external dependencies to implement the crawlwer. However I used some external dependencies that I explain here:
- cobra / viper for cli, because I am really used to them and it makes really easy to parse arguments
- testify for testing, because it makes really easy to assert results
- x/net/html for html parsing, because it is the standard library for html parsing in golang