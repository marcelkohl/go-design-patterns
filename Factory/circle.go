package main

type circle struct {
  shape
}

func (c *circle) new() shapeInterface {
  c.shapeType = "circle"
  c.angles = 0

  return c;
}
