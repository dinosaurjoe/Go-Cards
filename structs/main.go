package main

import "fmt"

type contactInfo struct {
  github string
  zipCode int
}

type person struct {
  firstName string
  lastName string
  contact contactInfo
}

func main() {
  joe := person{
    firstName: "Joe",
    lastName: "Schafer",
    contact: contactInfo{
      github: "dinosaurjoe",
      zipCode: 60622
    }
  }
}
