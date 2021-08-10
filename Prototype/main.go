package main

import "fmt"

func main() {
  shape1 := circle{
    name: "circle",
    radius: 2,
    color: "blue",
    border: false,
  }
  fmt.Println(shape1.name, shape1.color)

  shape2 := shape1.clone()
  fmt.Println(shape2.name, shape2.color)
}
