package main

import (
"fmt"
"bufio"
"os"
"math/rand"
"time"
)

func main() {
  input := input()
  fmt.Println(input)
  newSuit := newSuit()
  fmt.Println(newSuit)
}

func input() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    return text
}


func newSuit() string {

  cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}


  rand.Seed(time.Now().UnixNano())

         // flip the coin
  suit := cardSuits[rand.Intn(len(cardSuits))]

  return suit
}