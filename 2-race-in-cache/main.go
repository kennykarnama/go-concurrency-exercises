//////////////////////////////////////////////////////////////////////
//
// Given is some code to cache key-value pairs from a database into
// the main memory (to reduce access time). Note that golang's map are
// not entirely thread safe. Multiple readers are fine, but multiple
// writers are not. Change the code to make this thread safe.
//

package main

import (
	"github.com/kennykarnama/go-concurrency-exercises/2-race-in-cache/fake"
)

func main() {
	fake.Run()
}
