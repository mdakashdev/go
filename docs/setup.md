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

run
```bash
go run main.go
```


> Laravel-এ যেমন application-এর entry point থাকে public/index.php, Go executable application শুরু হয় main() function থেকে।