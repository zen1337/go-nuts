package main

import (
  "fmt"
  "strconv"
)

func main() {
  var a int = 1
  fmt.Printf("%v is a %T\n", a, a)
  b := a + a
  c := true
  if c {
    c := b
    fmt.Printf("%v is a %T\n", c, c)
  }


  fmt.Printf("%v", (b + b))
  fmt.Println(b + b)
  fmt.Println(strconv.FormatUint(345345345, 16))
  fmt.Println(strconv.Itoa(345345345))
  fmt.Println(strconv.FormatBool(isBool(0)) + "\n")
  fmt.Printf("%v is a %T\n", c, c)
}

func test(someVar int) string {
  x := string(someVar)
  return x
}

func isBool(x int) bool {
  b := x != 1
  return b
}
