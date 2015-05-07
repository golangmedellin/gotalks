// +build ignore

package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// START OMIT
func main() {
    res, _ := http.Get("http://localhost:3000/categories")
    fmt.Printf("Response Status Code: %d\n", res.StatusCode)
    fmt.Printf("Response Content-Type: %s\n", res.Header.Get("Content-Type"))
    defer res.Body.Close()
    categories := []interface{}{}
    err := json.NewDecoder(res.Body).Decode(&categories)
    if err != nil {
        panic(err)
    }
    fmt.Println("Response Body:")
    for _, category := range categories {
        fmt.Printf("%v\n", category)
    }
}

// END OMIT
