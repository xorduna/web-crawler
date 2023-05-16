package crawler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiskBrowser_Get(t *testing.T) {

	//make a test list of links
	testCases := []struct {
		url      string
		numLinks int
	}{
		{
			url:      "https://fakesite.com/index.html",
			numLinks: 5,
		},
		{
			url:      "https://fakesite.com/",
			numLinks: 5,
		},
	}

	fe := NewDiskBrowser("https://fakesite.com")

	for _, tc := range testCases {

		links, err := fe.Get(tc.url)
		if err != nil {
			t.Error("Expected no error")
		}

		assert.NotNil(t, links)
		assert.Nil(t, err)
		assert.Len(t, links, tc.numLinks)

	}
}
