package main

type cassetteInterface interface {
  handle(remainingValue int)
  setNext(cassette cassetteInterface)
}
