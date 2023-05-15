package main

import (
	"io"
	"net/http"
)

type WebExtractor struct {
}

func NewWebExtractor() *WebExtractor {
	return &WebExtractor{}
}

func (w *WebExtractor) Get(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
