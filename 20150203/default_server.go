// +build ignore

package main

import (
    . "github.com/sescobb27/meetup/handlers"
    "net/http"
)

// START OMIT
func main() {
    server := http.NewServeMux()
    server.Handle("/categories", Get(GetCategories))
    server.Handle("/locations", Get(GetLocations))
    server.Handle("/css/",
        http.StripPrefix("/css/",
            http.FileServer(
                http.Dir("resources/css"))))
    server.Handle("/js/",
        http.StripPrefix("/js/",
            http.FileServer(
                http.Dir("resources/js"))))
    http.ListenAndServe(":3000", server)
}

// END OMIT
