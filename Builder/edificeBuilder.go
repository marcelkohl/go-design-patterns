package main

type edificeBuilder struct {
  edifice
}

func (c *edificeBuilder) bringMaterial() {
  c.material = "Bricks"
}

func (c *edificeBuilder) prepareGarage() {
  c.hasGarage = true
}

func (c *edificeBuilder) putOnPlace() {
  c.place = "City"
}

func (c *edificeBuilder) getConstruction() construction {
  return construction{
    material: c.material,
    hasGarage: c.hasGarage,
    place: c.place,
  }
}
