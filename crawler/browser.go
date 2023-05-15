package crawler

import "io"

type Browser interface {
	Get(url string) (io.Reader, error)
}
