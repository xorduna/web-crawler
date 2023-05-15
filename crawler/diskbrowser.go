package crawler

import (
	"io"
	"os"
	"strings"
)

type DiskBrowser struct {
	baseSite string
}

func NewDiskBrowser(baseSite string) *DiskBrowser {
	return &DiskBrowser{
		baseSite: baseSite,
	}
}

func (f *DiskBrowser) Get(url string) (io.Reader, error) {
	// remove host from url
	path := strings.ReplaceAll(url, f.baseSite, "")
	if path == "" || path == "/" {
		path = "/index.html"
	}

	filePath := "fixtures/fakesite/" + path

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	reader := io.Reader(file)

	return reader, nil
}
