package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World from Go Application from Prod\n")
    })

    http.ListenAndServe(":9000", nil)
}

