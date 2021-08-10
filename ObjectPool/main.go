package main

import "fmt"

func main() {
  pool := (samplesPool{}).new(2)

  obj, _ := pool.acquire()
  fmt.Println("got 1st object: ", obj.getId())

  fmt.Println("released 1st object: ", obj.getId())
  pool.release(obj)

  obj2, _ := pool.acquire()
  fmt.Println("got 2nd object (must be the same as 1st as we released it before acquiring a new one)", obj2.getId())

  obj3, _ := pool.acquire()
  fmt.Println("got 3rd object (must be new)", obj3.getId())

  obj4, err := pool.acquire()

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("got 4th object (must not happen as we reached limit of 2)", obj4.getId())
  }
}
