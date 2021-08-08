package main

type campHouseBuilder struct {
  campHouse
}

func (c *campHouseBuilder) bringMaterial() {
  c.material = "Wood"
}

func (c *campHouseBuilder) prepareGarage() {
  c.hasGarage = false
}

func (c *campHouseBuilder) putOnPlace() {
  c.place = "Forest"
}

func (c *campHouseBuilder) getConstruction() construction {
  return construction{
    material: c.material,
    hasGarage: c.hasGarage,
    place: c.place,
  }
}
