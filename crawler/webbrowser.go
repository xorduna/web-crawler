package crawler

import (
	"io"
	"net/http"
)

type WebBrowser struct {
}

func NewWebBrowser() *WebBrowser {
	return &WebBrowser{}
}

func (w *WebBrowser) Get(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
