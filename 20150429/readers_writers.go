// +build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

// SharedResource is the type representing shared resource.
// it's contents could be anything
type SharedResource struct {
	string
}

// Access struct wraps pointer to SharedResource with embedded RWMutex
type Access struct {
	sync.RWMutex
	res *SharedResource
}

// reader reads three times at about 1 second intervals, taking
// 2 second to perform the read
func reader(name string, acc *Access, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("reader", name, "ready to read")
		acc.RLock()
		fmt.Println("reader", name, "reading")
		time.Sleep(2 * time.Second)
		msg := acc.res.string
		acc.RUnlock()
		fmt.Println("reader", name, "read:", msg)
	}
	wg.Done()
}

// writer writes twice
func writer(name string, acc *Access, wg *sync.WaitGroup) {
	claim := []string{"once", "again"}
	for i := 0; i < 2; i++ {
		time.Sleep(1 * time.Second)
		msg := name + " was here " + claim[i]
		fmt.Println("writer", name, "ready to write")
		acc.Lock()
		fmt.Println("writer", name, "writing:", msg)
		time.Sleep(2 * time.Second)
		acc.res.string = msg
		acc.Unlock()
		fmt.Println("writer", name, "done")
	}
	wg.Done()
}

func main() {
	acc := &Access{res: &SharedResource{"zero"}}
	fmt.Println("Initial value:", acc.res.string)

	wg := new(sync.WaitGroup)
	wg.Add(5) // three readers and two writers

	go reader("A", acc, wg)
	go reader("B", acc, wg)
	go reader("C", acc, wg)

	go writer("X", acc, wg)
	go writer("Y", acc, wg)

	wg.Wait()
	fmt.Println("Final value:", acc.res.string)
}
