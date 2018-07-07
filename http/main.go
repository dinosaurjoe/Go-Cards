package main

import "fmt"
import "net/http"
import "os"

func main() {
  resp, err := http.Get("http://bing.com")

  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }

  bs := make([]byte, 99999)
  resp.Body.Read(bs)
  fmt.Println(string(bs))
}
