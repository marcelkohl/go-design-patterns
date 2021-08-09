package main

import "fmt"

func getShape(name string) (shapeInterface, error) {
    if name == "circle" {
      shape := &circle{}
      return shape.new(), nil
    }

    return nil, fmt.Errorf("Unknown shape type")
}
