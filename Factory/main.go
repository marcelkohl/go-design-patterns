package main

import "fmt"

func main() {
    shapeCircle, _ := getShape("circle")
    // maverick, _ := getGun("maverick")
    printDetails(shapeCircle)
    // printDetails(maverick)
}

func printDetails(s shapeInterface) {
    fmt.Printf("Type: %s\n", s.getShapeType())
    fmt.Printf("Amgles: %d\n", s.getAmountOfAngles())
}
