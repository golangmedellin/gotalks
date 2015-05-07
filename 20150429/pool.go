// +build ignore

// Pool’s intended use is for free lists maintained in global variables,
// typically accessed by multiple goroutines simultaneously.
// Using a Pool instead of a custom free list allows the runtime to reclaim
// entries from the pool when it makes sense to do so. An appropriate use
// of sync.Pool is to create a pool of temporary buffers shared between
// independent clients of a global resource. On the other hand, if a free list
// is maintained as part of an object used only by a single client and freed
// when the client completes, implementing that free list as a Pool is not appropriate.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	p := sync.Pool{
		New: func() interface{} {
			return "New Pool"
		},
	}

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			p.Put("Put Value")
			time.Sleep(100 * time.Millisecond)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(p.Get())
			time.Sleep(50 * time.Millisecond)
		}
		wg.Done()
	}()

	wg.Wait()
}
