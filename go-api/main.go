package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Employee struct {
    ID       int
    Name     string
    Email    string
    Position string
}

type Worker struct {
    ID int `json:"uuid"`
    Name string `json:"name"`
    Email string `json:"email"`
    Position string `json:"position"`
}

func workerHandler(w http.ResponseWriter, r *http.Request) {
    worker := Worker{
        ID:       1,
        Name:     "Akash",
        Email:    "akash@example.com",
        Position: "Software Engineer",
    }

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(worker)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

    response := map[string]string{
        "status":  "success",
        "message": "Server is healthy!",
    }

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(response)
}



func sayHello() string {
    return "Welcome to Go Backend"
}

func main() {
    //function sayHello
    message := sayHello()
    fmt.Println(message)

    //struct
    employee := Employee{
        ID:       1,
        Name:     "Akash",
        Email:    "akash@example.com",
        Position: "Software Engineer",
    }
    fmt.Println(employee.Name)
    fmt.Println(employee.Email)

    //handle-worker using json
    http.HandleFunc("/worker", workerHandler)

    http.HandleFunc("/health", healthHandler)

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to Go Backend Development!")
    })

    fmt.Println("Server running on port 8080")

    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        fmt.Println("Server error:", err)
    }
}
