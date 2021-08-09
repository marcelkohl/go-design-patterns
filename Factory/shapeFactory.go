package main

import "fmt"

func getShape(name string) (shapeInterface, error) {
    if name == "circle" {
      return (&circle{}).new(), nil
    } else if name == "square" {
      return (&square{}).new(), nil
    }

    return nil, fmt.Errorf("Unknown shape type")
}
