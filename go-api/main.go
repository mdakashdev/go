package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello Go")
    })

    fmt.Println("Server running on port 8080")

    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        fmt.Println("Server error:", err)
    }
}
