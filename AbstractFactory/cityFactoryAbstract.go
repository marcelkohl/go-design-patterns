package main

import "fmt"

func getCityFactory(cityRef string) (iCityFactory, error) {
  if cityRef == "Ab" {
      return &cityAbFactory{}, nil
  } else if cityRef == "Xy" {
      return &cityXyFactory{}, nil
  }
  return nil, fmt.Errorf("Wrong city reference passed")
}
