package lib

import (
	"sync"
)

type SafeVisited interface {
	AddVisited(url string)
	IsVisited(url string) bool
	List() []string
}

type SafeMap struct {
	visitedUrls map[string]bool
	lock        sync.RWMutex
}

func NewSafeMap() SafeVisited {
	return &SafeMap{
		visitedUrls: make(map[string]bool),
	}
}

func (s *SafeMap) AddVisited(url string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.visitedUrls[url] = true
}

func (s *SafeMap) IsVisited(url string) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.visitedUrls[url]
	return ok
}

func (s *SafeMap) List() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	var urls []string
	for k, _ := range s.visitedUrls {
		urls = append(urls, k)
	}
	return urls
}
