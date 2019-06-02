//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package main

import (
	"testing"

	"github.com/kennykarnama/go-concurrency-exercises/2-race-in-cache/fake"
)

func TestMain(t *testing.T) {
	cache := fake.Run()

	cacheLen := len(cache.Cache)
	pagesLen := cache.Pages.Len()
	if cacheLen != fake.CacheSize {
		t.Errorf("Incorrect cache size %v", cacheLen)
	}
	if pagesLen != fake.CacheSize {
		t.Errorf("Incorrect pages size %v", pagesLen)
	}
}
