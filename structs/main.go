package main

import "fmt"

type person struct {
  firstName string
  lastName string
}

func main() {
  joe := person{firstName: "Joe", lastName:  "Schafer"}
  fmt.Println(joe)
}
