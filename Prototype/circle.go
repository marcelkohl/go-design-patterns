package main

type circle struct {
  name string
  radius int
  color string
  border bool
}

func (c circle) clone() circle {
  return circle{
    name: c.name + "_clone",
    radius: c.radius,
    color: c.color,
    border: c.border,
  }
}
