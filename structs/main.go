package main

import "fmt"

type contactInfo struct {
  github string
  zipCode int
}

type person struct {
  firstName string
  lastName string
  contactInfo
}

func main() {
  joe := person{
    firstName: "Joe",
    lastName: "Schafer",
    contactInfo: contactInfo{
      github: "dinosaurjoe",
      zipCode: 60622,
    },
  }

  joe.updateName("Joseph")
  joe.print()
}

func (p person) updateName(newFirstName string) {
  p.firstName = newFirstName
}

func (p person) print() {
  fmt.Printf("%+v", p)
}
