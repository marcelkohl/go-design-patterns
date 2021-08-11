package main

import "fmt"

type cassette50 struct {
  cassette
}

func (c *cassette50) handle(remainingValue int) {
  valueHandled := ((remainingValue / 50) * 50)
  fmt.Println("cassete 50 for ", remainingValue, " handled ", valueHandled)

  if c.next != nil {
    c.next.handle(remainingValue - valueHandled)
  }
}
