package main

import (
"fmt"
"bufio"
"strings"
"os"
)

type bot interface {
  getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
  eb := englishBot{}
  sb := spanishBot{}

  input := input()

  if strings.TrimRight(input, "\n") == "hola" {
      printGreeting(sb)
    } else if strings.TrimRight(input, "\n") == "hello" {
      printGreeting(eb)
      } else {
        fmt.Println("que?")
      }
}

func input() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    return text
}


func printGreeting(b bot) {
  fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
  // custom logic for generating english greeting
  return "Hello."
}

func (spanishBot) getGreeting() string {
  return "Hola."
}
