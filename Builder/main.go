package main

import "fmt"

func main() {
  houseBuilt := newDirector(&campHouseBuilder{}).buildConstruction()
  edificeBuilt := newDirector(&edificeBuilder{}).buildConstruction()

  fmt.Println("Camp house info:")
  printInfo(houseBuilt)

  fmt.Println("\nEdifice info:")
  printInfo(edificeBuilt)
}

func printInfo(c construction) {
  fmt.Printf("Place: %s\n", c.place)
  fmt.Printf("Material: %s\n", c.material)
  fmt.Printf("Has garage: %t\n", c.hasGarage)
}
