package crawler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLink(t *testing.T) {
	testCases := []struct {
		input    string
		expected Link
		err      error
	}{
		{
			input: "https://www.example.com/path/to/page.html",
			expected: Link{
				Host:     "https://www.example.com",
				Path:     "/path/to/",
				Document: "page.html",
			},
			err: nil,
		},
		{
			input: "https://www.example.com/path/",
			expected: Link{
				Host:     "https://www.example.com",
				Path:     "/path/",
				Document: "",
			},
			err: nil,
		}, {
			input: "https://example.com",
			expected: Link{
				Host:     "https://example.com",
				Path:     "",
				Document: "",
			},
			err: nil,
		},
		{
			input: "www.example.com",
			expected: Link{
				Host:     "",
				Path:     "",
				Document: "www.example.com",
			},
			err: nil,
		},
		{
			input: "example.com",
			expected: Link{
				Host:     "",
				Path:     "",
				Document: "example.com",
			},
			err: nil,
		},
		{
			input: "/path/to/document.html",
			expected: Link{
				Host:     "",
				Path:     "/path/to/",
				Document: "document.html",
			},
			err: nil,
		},
		{
			input: "/path",
			expected: Link{
				Host:     "",
				Path:     "/",
				Document: "path",
			},
			err: nil,
		},
		{
			input:    "document.html",
			expected: Link{Host: "", Path: "", Document: "document.html"},
			err:      nil,
		},
		{
			input: "/path/to/document.html?myquery=1",
			expected: Link{
				Host:     "",
				Path:     "/path/to/",
				Document: "document.html",
				Query:    "myquery=1",
			},
			err: nil,
		},
		{
			input: "/path?myquery=1",
			expected: Link{
				Host:     "",
				Path:     "/",
				Document: "path",
				Query:    "myquery=1",
			},
			err: nil,
		},
		{
			input: "document.html?myquery=1",
			expected: Link{
				Host:     "",
				Path:     "",
				Document: "document.html",
				Query:    "myquery=1",
			},
			err: nil,
		},
		{
			input:    "javascript:alert('hello')",
			expected: Link{},
			err:      fmt.Errorf("invalid URL: javascript:alert('hello')"),
		},
		{
			input:    "mailto:user@email.com",
			expected: Link{},
			err:      fmt.Errorf("invalid URL: mailto:user@email.com"),
		}, {
			input:    "#myanchor",
			expected: Link{},
			err:      fmt.Errorf("invalid URL: #myanchor"),
		},
	}

	for _, tc := range testCases {
		actual, err := ParseLink(tc.input)
		//fmt.Printf("%s -> %+v\n", err, actual)

		assert.Equal(t, tc.expected, actual, "Unexpected result for input string: %q", tc.input)
		assert.Equal(t, tc.err, err, "Unexpected error for input string: %q", tc.input)
		if err == nil {
			assert.Equal(t, tc.input, actual.FullLink())
		}
	}
}
