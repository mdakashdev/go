```bash
brew install go
```

```bash
go version
```

এবার Go Module initialize করো:

```bash
go mod init go-api
```

then, main.go তৈরি করো।

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Akash!")
}
```

```go 

package main

import (
    "fmt"
    "net/http"
)

func main() {

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

        fmt.Fprintln(w, "Hello Go!")

    })

    fmt.Println("Server running on port 8080")

    http.ListenAndServe(":8080", nil)
}

```



run
```bash
go run main.go
```


> Laravel-এ যেমন application-এর entry point থাকে public/index.php, Go executable application শুরু হয় main() function থেকে।


# Api

Your first API : http://localhost:8080/health