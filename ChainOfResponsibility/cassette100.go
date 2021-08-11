package main

import "fmt"

type cassette100 struct {
  cassette
}

func (c *cassette100) handle(remainingValue int) {
  valueHandled := ((remainingValue / 100) * 100)
  fmt.Println("cassete 100 for ", remainingValue, " handled ", valueHandled)

  if c.next != nil {
    c.next.handle(remainingValue - valueHandled)
  }
}
