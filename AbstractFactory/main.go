package main

import "fmt"

func main() {
  fmt.Println("City Ab info:")
  cityAb, _ := getCityFactory("Ab")
  printDemographics(cityAb.makeDemographics())
  printExpenses(cityAb.makeExpenses())
}

func printDemographics(d iDemographics) {
  fmt.Printf("Population: %d\n", d.getPopulation())
  fmt.Printf("Majority gender: %s\n", d.getMajorityGender())
}

func printExpenses(e iExpenses) {
  fmt.Printf("Forecast: %.2f\n", e.getForecast())
  fmt.Printf("Spent: %.2f\n", e.getSpent())
}
