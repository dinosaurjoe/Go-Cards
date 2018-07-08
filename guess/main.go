package main

import (
"fmt"
"bufio"
"os"
"math/rand"
"time"
"strings"
)

func main() {
  input := input()
  fmt.Println(input)
  newSuit := newSuit()
  fmt.Println(newSuit)
  
  if strings.TrimRight(input, "\n") == newSuit {
    fmt.Println("You have guessed the suit correctly")
  } else {
    fmt.Println("You have guessed incorrectly, try again!")
  }
}

func input() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    return text
}


func newSuit() string {

  cardSuits := []string{"hearts", "spades", "clubs", "diamonds"}


  rand.Seed(time.Now().UnixNano())

         // flip the coin
  suit := cardSuits[rand.Intn(len(cardSuits))]

  return suit
}
