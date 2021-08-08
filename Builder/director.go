package main

type director struct {
  builder builderInterface
}

func newDirector(b builderInterface) *director {
  return &director{
    builder: b,
  }
}

func (d *director) buildConstruction() construction {
  d.builder.bringMaterial()
  d.builder.prepareGarage()
  d.builder.putOnPlace()
  
  return d.builder.getConstruction()
}
