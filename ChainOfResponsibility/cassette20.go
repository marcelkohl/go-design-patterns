package main

import "fmt"

type cassette20 struct {
  cassette
}

func (c *cassette20) handle(remainingValue int) {
  valueHandled := ((remainingValue / 20) * 20)
  fmt.Println("cassete 20 for ", remainingValue, " handled ", valueHandled)

  if c.next != nil {
    c.next.handle(remainingValue - valueHandled)
  }
}
