package main

type square struct {
  shape
}

func (s *square) new() shapeInterface {
  s.shapeType = "square"
  s.angles = 4

  return s;
}
