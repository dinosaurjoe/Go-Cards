package main

import "fmt"
import "net/http"
import "os"

func main() {
  resp, err := http.Get("http://google.com")

  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }

  fmt.Println(resp)
}
