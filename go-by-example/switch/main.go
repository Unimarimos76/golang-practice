package main

import (
  "fmt"
)

func main() {
  i := 2
  fmt.Print("Write ", i, " as ")
  switch i {
    case 1 :
      fmt.Println("one")
    case 2 :
      fmt.Println("two")
    case 3 :
      fmt.Println("three")
  }

  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("I`m a bool")
    case int:
      fmt.Println("I`m an int")
    default:
      fmt.Println("Dont know type %T", t)
    }
  }

  whatAmI(true)
  whatAmI(1)
}
