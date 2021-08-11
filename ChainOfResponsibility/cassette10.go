package main

import "fmt"

type cassette10 struct {
  cassette
}

func (c *cassette10) handle(remainingValue int) {
  valueHandled := ((remainingValue / 10) * 10)
  fmt.Println("cassete 10 for ", remainingValue, " handled ", valueHandled)

  if c.next != nil {
    c.next.handle(remainingValue - valueHandled)
  }
}
