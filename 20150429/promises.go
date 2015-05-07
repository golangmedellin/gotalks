// +build ignore

package main

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func panicIfError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// Promise interface for handling async responses anc callbacks
// START PROMISE OMIT
type Promise interface {
	Then(func(value interface{}) Promise) Promise
}

// END PROMISE OMIT

// AsyncRequest implements Promise interface and has a buffered int64 chan
// for syncronization
// START REQUEST OMIT
type AsyncRequest struct {
	Promise
	body chan int64
}

// NewAsyncRequest returns a new AsyncRequest object
// END REQUEST OMIT
func NewAsyncRequest() *AsyncRequest {
	return &AsyncRequest{
		body: make(chan int64),
	}
}

// Then satisfy Promise interface
// START IMPLEMENTATION OMIT
func (aReq *AsyncRequest) Then(chain func(value interface{}) Promise) Promise {
	return chain(<-aReq.body)
}

// Get perform a http.Get request and return a Promise for handling it's response
// asyncronously
func (aReq *AsyncRequest) Get(path string) Promise {
	go func() {
		res, err := http.Get(path)
		defer res.Body.Close()
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			os.Exit(1)
		}
		var answer int64
		err = binary.Read(res.Body, binary.LittleEndian, &answer)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			os.Exit(1)
		}
		aReq.body <- answer
	}()
	return aReq
}

// END IMPLEMENTATION OMIT

func main() {
	req := NewAsyncRequest()
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(r *AsyncRequest, iter int) {
			fmt.Println("Asynchronous Gets")
			r.Get("http://localhost:3000/odd").Then(func(response interface{}) Promise {
				fmt.Fprintf(os.Stdout, "(ODD) From Promise %d value => %v\n", iter, response)
				req2 := NewAsyncRequest()
				return req2.Get("http://localhost:3000/pair").Then(func(response interface{}) Promise {
					fmt.Fprintf(os.Stdout, "(PAIR) From Innert Promise %d value => %v\n", iter, response)
					wg.Done()
					return nil
				})
			})
		}(req, i)
	}
	wg.Wait()
}
