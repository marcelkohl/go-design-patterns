package main

import "fmt"

func main() {
  c5 := &cassette5{}
  c10 := &cassette10{}
  c10.setNext(c5)
  c20 := &cassette20{}
  c20.setNext(c10)
  c50 := &cassette50{}
  c50.setNext(c20)
  c100 := &cassette100{}
  c100.setNext(c50)

  c100.handle(185);

  fmt.Println("done")
}
