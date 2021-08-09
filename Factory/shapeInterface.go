package main

type shapeInterface interface {
  new() shapeInterface
  getShapeType() string
  getAmountOfAngles() int
}
