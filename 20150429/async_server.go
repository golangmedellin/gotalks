// +build ignore

package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
)

// START SERVER OMIT
var asyncCount int64
var atomicCount int64

func main() {
	http.HandleFunc("/async", func(res http.ResponseWriter, req *http.Request) {
		asyncCount++
		fmt.Fprintf(os.Stdout, "ASYNC: %d\n", asyncCount)
	})

	http.HandleFunc("/atomic", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(os.Stdout, "ATOMIC: %d\n", atomic.AddInt64(&atomicCount, 1))
	})

	// error handling omitted for demostrative purposes but NOT DO THIS
	http.ListenAndServe(":3000", nil)
}

// END SERVER OMIT
