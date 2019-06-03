//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package fake

import (
	"container/list"
	"strconv"
	"sync"
)

const (
	cycles        = 3
	callsPerCycle = 100
)

// CacheSize determines how big the cache can grow
const CacheSize = 100

// KeyStoreCacheLoader is an interface for the KeyStoreCache
type KeyStoreCacheLoader interface {
	// Load implements a function where the cache should gets it's content from
	Load(string) string
}

// KeyStoreCache is a LRU cache for string key-value pairs
type KeyStoreCache struct {
	Cache map[string]string
	Pages list.List
	Load  func(string) string
	mut   sync.Mutex
}

// New creates a new KeyStoreCache
func New(load KeyStoreCacheLoader) *KeyStoreCache {
	return &KeyStoreCache{
		Load:  load.Load,
		Cache: make(map[string]string),
	}
}

// Get gets the key from cache, loads it from the source if needed
func (k *KeyStoreCache) Get(key string) string {
	k.mut.Lock()
	defer k.mut.Unlock()
	val, ok := k.Cache[key]

	// Miss - load from database and save it in cache

	if !ok {
		val = k.Load(key)
		k.Cache[key] = val
		k.Pages.PushFront(key)

		// if cache is full remove the least used item
		if len(k.Cache) > CacheSize {
			delete(k.Cache, k.Pages.Back().Value.(string))
			k.Pages.Remove(k.Pages.Back())
		}
	}

	return val
}

// Loader implements KeyStoreLoader
type Loader struct {
	DB *MockDB
}

// Load gets the data from the database
func (l *Loader) Load(key string) string {
	val, err := l.DB.Get(key)
	if err != nil {
		panic(err)
	}

	return val
}

func Run() *KeyStoreCache {
	loader := Loader{
		DB: GetMockDB(),
	}
	cache := New(&loader)

	RunMockServer(cache)

	return cache
}

// RunMockServer simulates a running server, which accesses the
// key-value database through our cache
func RunMockServer(cache *KeyStoreCache) {
	var wg sync.WaitGroup

	for c := 0; c < cycles; c++ {
		wg.Add(1)
		go func() {
			for i := 0; i < callsPerCycle; i++ {

				cache.Get("Test" + strconv.Itoa(i))

			}
			wg.Done()
		}()
	}

	wg.Wait()
}
