package crawler

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractLinks(t *testing.T) {
	testCases := []struct {
		input    string
		expected Link
		err      error
	}{
		{
			input: "fixtures/links/testlinks-1.html",
			expected: Link{
				Host:     "https://www.example.com",
				Path:     "/path/to/",
				Document: "page.html",
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		inputPath := filepath.Join(tc.input)

		r, err := os.Open(inputPath)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		links, err := ExtractLinksFromHTML(r)
		assert.NotNil(t, links)
		assert.Nil(t, err)
		for _, l := range links {
			t.Logf("link: %+v", l)
		}

	}
}
