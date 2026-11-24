package main

import (
    "fmt"
    "net/http"
)

type Employee struct {
    ID       int
    Name     string
    Email    string
    Position string
}

func sayHello() string {
    return "Welcome to Go Backend"
}

func main() {
    message := sayHello()

    fmt.Println(message)

    employee := Employee{
        ID:       1,
        Name:     "Akash",
        Email:    "akash@example.com",
        Position: "Software Engineer",
    }

    fmt.Println(employee.Name)
    fmt.Println(employee.Email)


    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Server is healthy!")
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to Go Backend Development!")
    })

    fmt.Println("Server running on port 8080")

    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        fmt.Println("Server error:", err)
    }
}
