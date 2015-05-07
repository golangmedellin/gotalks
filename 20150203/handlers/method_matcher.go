package handlers

import (
    "fmt"
    "net/http"
)

func methodMatcher(method string, handler http.HandlerFunc) http.HandlerFunc {
    var err string
    return func(res http.ResponseWriter, req *http.Request) {
        if method != req.Method {
            err = fmt.Sprintf("Expected %s Http Method, Got: %s Http Method", method, req.Method)
            http.Error(res, err, http.StatusBadRequest)
        } else {
            handler.ServeHTTP(res, req)
        }
    }
}

func Get(handler http.HandlerFunc) http.HandlerFunc {
    return methodMatcher("GET", handler)
}

func Post(handler http.HandlerFunc) http.HandlerFunc {
    return methodMatcher("POST", handler)
}

func Put(handler http.HandlerFunc) http.HandlerFunc {
    return methodMatcher("PUT", handler)
}

func Delete(handler http.HandlerFunc) http.HandlerFunc {
    return methodMatcher("DELETE", handler)
}

func Patch(handler http.HandlerFunc) http.HandlerFunc {
    return methodMatcher("PATCH", handler)
}
