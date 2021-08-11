package main

type cassette struct {
  next cassetteInterface
}

func (c *cassette) handle(remainingValue int) {
  if c.next != nil {
    c.next.handle(remainingValue)
  }
}

func (c *cassette) setNext(cassette cassetteInterface) {
  c.next = cassette
}
