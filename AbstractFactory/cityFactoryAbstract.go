package main

import "fmt"

func getCityFactory(cityRef string) (iCityFactory, error) {
  if cityRef == "Ab" {
      return &cityAbFactory{}, nil
  }
  // if brand == "nike" {
  //     return &nike{}, nil
  // }
  return nil, fmt.Errorf("Wrong city reference passed")
}
