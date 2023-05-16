package crawler

import (
	"net/http"
)

type WebBrowser struct {
}

func NewWebBrowser() *WebBrowser {
	return &WebBrowser{}
}

func (w *WebBrowser) Get(url string) ([]Link, error) {
	//nolint: gosec, bodyclose, noctx
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ExtractLinksFromHTML(resp.Body)
}
