package crawler

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestDiskBrowser_Get(t *testing.T) {

	//make a test list of links
	testCases := []struct {
		link string
	}{
		{
			link: "https://fakesite.com/index.html",
		},
		{
			link: "https://fakesite.com/",
		},
	}

	fe := NewDiskBrowser("https://fakesite.com")

	for _, tc := range testCases {

		reader, err := fe.Get(tc.link)
		if err != nil {
			t.Error("Expected no error")
		}

		assert.NotNil(t, reader)
		data, err := io.ReadAll(reader)
		assert.NotNil(t, data)
		assert.Nil(t, err)

	}
}
