// +build ignore

package main

import (
	"encoding/binary"
	"math/rand"
	"net/http"
)

// START SERVER OMIT

func main() {
	odd := make([]int64, 100)
	pair := make([]int64, 100)
	for i := 0; i < 100; i++ {
		pair[i] = int64(i)
		odd[i] = int64(i + 1)
		i++
	}
	rand.Seed(100)
	http.HandleFunc("/odd", func(res http.ResponseWriter, req *http.Request) {
		// error handling omitted for demostrative purposes but NOT DO THIS
		binary.Write(res, binary.LittleEndian, odd[rand.Int63n(100)])
	})

	http.HandleFunc("/pair", func(res http.ResponseWriter, req *http.Request) {
		// error handling omitted for demostrative purposes but NOT DO THIS
		binary.Write(res, binary.LittleEndian, pair[rand.Int63n(100)])
	})

	// error handling omitted for demostrative purposes but NOT DO THIS
	http.ListenAndServe(":3000", nil)
}

// END SERVER OMIT
