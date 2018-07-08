package main

import (
"fmt"
"bufio"
"os"
)

func main() {
  input := input()
  fmt.Println(input)
}

func input() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    return text
}


