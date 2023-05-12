package crawler

import (
	"fmt"
	"regexp"
	"strings"
)

type Link struct {
	Host     string
	Path     string
	Document string
	Query    string
}

func (l Link) FullLink() string {
	if l.Query == "" {
		return fmt.Sprintf("%s%s%s", l.Host, l.Path, l.Document)
	} else {
		return fmt.Sprintf("%s%s%s?%s", l.Host, l.Path, l.Document, l.Query)
	}
}

// Regular expression to match URLs with named captured groups for host, path, and document
// oldRegex := "^(?P<host>^https?:\\/\\/[^\\/]+)?(?P<path>\\/[^?#]+\\/)?(?P<document>[^?#\\/]+)?$"
// (?P<host>https?://[^"/]+)?(?P<path>(?:/|^)(.*?)(?P<document>[^/"\s]+))$
// urlRegex := regexp.MustCompile(`(?P<host>https?://[^"/]+)?(?P<path>(?:/|^)(.*?)(?P<document>[^/"\s]+))$`)

func ParseLink(str string) (Link, error) {
	//check if string starts with "#" or "javascript" to discard javascript commands and anchors
	if strings.HasPrefix(str, "#") || strings.HasPrefix(str, "javascript") || strings.HasPrefix(str, "mailto") {
		return Link{}, fmt.Errorf("invalid URL: %s", str)
	}

	// look for a host in str
	hostRegex := regexp.MustCompile(`(?P<host>https?://[^"/]+)`)
	hostMatch := hostRegex.FindStringSubmatch(str)

	host := ""
	path := str

	if len(hostMatch) > 0 {
		// if host is found, remove it from str
		path = strings.Replace(str, hostMatch[0], "", 1)

		//get host as a variable
		host = hostMatch[0]
	}

	// given a path, get the paht, document and query as named groups
	pathRegex := regexp.MustCompile(`^(?P<path>.*/)?(?P<document>[^\/?]*)?(?:\?(?P<query>.+))?$`)

	// get path, document and query as variables using named groups
	pathMatch := pathRegex.FindStringSubmatch(path)
	paramsMap := make(map[string]string)

	if pathMatch != nil {
		for i, name := range pathRegex.SubexpNames() {
			if i != 0 && name != "" {
				paramsMap[name] = pathMatch[i]
			}
		}
	}

	path = paramsMap["path"]
	document := paramsMap["document"]
	query := paramsMap["query"]

	return Link{
		Host:     host,
		Path:     path,
		Document: document,
		Query:    query,
	}, nil
}
