// +build ignore

package main

import (
    . "github.com/sescobb27/meetup/handlers"
    "net/http"
)

func main() {
    server := http.NewServeMux()
    server.Handle("/", Get(Index_Handler))
    server.Handle("/users", Get(All_Handler))
    server.Handle("/user/new", Post(New_Handler))
    http.ListenAndServe(":8000", server)
}

func Index_Handler(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("Hello World"))
}

func New_Handler(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("Creating User"))
}

func All_Handler(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("All Users"))
}
