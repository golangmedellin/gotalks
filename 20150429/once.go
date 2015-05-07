// +build ignore

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// START ONCE OMIT
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	var once = &sync.Once{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				fmt.Println("Only once")
			})
			fmt.Println("greeting")
			wg.Done()
		}()
	}
	wg.Wait()
}

// END ONCE OMIT
