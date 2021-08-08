package main

type builderInterface interface {
    bringMaterial()
    prepareGarage()
    putOnPlace()
    getConstruction() construction
}
