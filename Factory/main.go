package main

import "fmt"

func main() {
    shapeCircle, _ := getShape("circle")
    shapeSquare, _ := getShape("square")
    printDetails(shapeCircle)
    printDetails(shapeSquare)
}

func printDetails(s shapeInterface) {
    fmt.Printf("Type: %s\n", s.getShapeType())
    fmt.Printf("Amgles: %d\n", s.getAmountOfAngles())
}
