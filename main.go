package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello India from Go Application From Stage brance 1\n")
    })

    http.ListenAndServe(":9000", nil)
}

