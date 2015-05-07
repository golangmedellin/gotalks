// +build ignore

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	// START ATOMIC OMIT
	var v int32
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&v, 1)
			atomic.AddInt32(&v, 1)
			atomic.AddInt32(&v, 1)
			atomic.AddInt32(&v, 1)
			atomic.AddInt32(&v, 1)
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			atomic.AddInt32(&v, 1)
			atomic.AddInt32(&v, 1)
			atomic.AddInt32(&v, 1)
			atomic.AddInt32(&v, 1)
			atomic.AddInt32(&v, 1)
			wg.Done()
		}()
	}
	// END ATOMIC OMIT
	wg.Wait()
	fmt.Println(v)
}
