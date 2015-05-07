package main

import (
    "fmt"
    "log"
    "net/http"
)

type Request struct {
    Filename   string
    ResultChan chan string
}

// process requests on background
func WorkerPool(n int) chan *Request {
    requests := make(chan *Request, n)

    for i := 0; i < n; i++ {
        go Worker(requests)
    }

    return requests
}

func Worker(requests chan *Request) {
    for request := range requests {
        // path := methodForFindingFiles()
        path := request.Filename
        request.ResultChan <- path
    }
}

type Server struct {
    Requests chan *Request
}

// START OMIT
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    filename := req.URL.Query().Get("filename")
    fmt.Printf("%v\n", req)
    request := &Request{Filename: filename, ResultChan: make(chan string)}
    s.Requests <- request
    path := <-request.ResultChan
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    fmt.Printf("%v\n\n\n", req)
    http.ServeFile(w, req, path)
}

func main() {
    requests := WorkerPool(5)
    server := &Server{Requests: requests}
    log.Fatal(http.ListenAndServe(":5000", server))
}

// END OMIT
