package main

import (
	"net/http"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			// error omitted for demostrative purposes but NOT DO THIS
			res, _ := http.Get("http://localhost:3000/async")
			res.Body.Close()
			wg.Done()
		}()
	}
	wg.Wait()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			// error omitted for demostrative purposes but NOT DO THIS
			res, _ := http.Get("http://localhost:3000/atomic")
			res.Body.Close()
			wg.Done()
		}()
	}
	wg.Wait()
}
