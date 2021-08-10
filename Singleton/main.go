package main

import (
  "fmt"
)

func main() {
  single := (singleton{}).new()

  for i := 0; i < 10; i++ {
    go single.getInstance()
  }

  fmt.Scanln()
}
