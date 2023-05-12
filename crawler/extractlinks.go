package crawler

import (
	"golang.org/x/net/html"
	"io"
)

func ExtractLinksFromHtml(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return getLinks(doc), nil
}

func getLinks(n *html.Node) []Link {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {

				link, err := ParseLink(a.Val)
				//fmt.Printf("Checking link %s -> %+v\n", a.Val, link)
				// Check if it is a link
				if err != nil {
					return []Link{}
				} else {
					return []Link{link}
				}
			}
		}
	}

	var links []Link
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, getLinks(c)...)
	}
	return links
}
