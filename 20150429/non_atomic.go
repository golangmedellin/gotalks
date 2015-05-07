// +build ignore

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	// START NON-ATOMIC OMIT
	var v int32
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			v++
			v++
			v++
			v++
			v++
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			v++
			v++
			v++
			v++
			v++
			wg.Done()
		}()
	}
	// END NON-ATOMIC OMIT
	wg.Wait()
	fmt.Println(v)
}
