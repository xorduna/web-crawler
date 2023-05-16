package lib

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSafeMap(t *testing.T) {

	sm := NewSafeMap()

	wg := sync.WaitGroup{}

	numDifferentUrls := 10

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(50)))
			sm.AddVisited(fmt.Sprintf("http://example-%d.com", a%numDifferentUrls))
			assert.True(t, sm.IsVisited(fmt.Sprintf("http://example-%d.com", a%numDifferentUrls)))
		}(i)
	}

	wg.Wait()

}
