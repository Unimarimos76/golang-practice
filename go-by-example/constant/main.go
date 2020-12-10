package main

import (
  "fmt"
  "math"
  "reflect"
)

const s string = "constant"

func main() {
    fmt.Println(reflect.TypeOf(s))

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(reflect.TypeOf(d))

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))

}
