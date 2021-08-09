package main

type shape struct {
  shapeType string
  angles    int
}

func (s *shape) getShapeType() string {
  return s.shapeType
}

func (s *shape) getAmountOfAngles() int {
  return s.angles
}
