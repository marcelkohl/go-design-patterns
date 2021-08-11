package main

import "fmt"

type cassette5 struct {
  cassette
}

func (c *cassette5) handle(remainingValue int) {
  valueHandled := ((remainingValue / 5) * 5)
  fmt.Println("cassete 5 for ", remainingValue, " handled ", valueHandled)

  if c.next != nil {
    c.next.handle(remainingValue - valueHandled)
  }
}
