# Factory Pattern
Factory pattern creates objects without exposing the creation to the client. The result refers to the created object using a common interface.

![Alt text](https://raw.githubusercontent.com/marcelkohl/go-design-patterns/main/Factory/diagram.png)

1. The client only interacts with a factory asking for the kind of objects needed.
2. The factory interacts with the corresponding concrete creator and returns back the correct object.


## Running
```
go run .
```

## Copyright and References
Some content here is based on:

- [Factory Design Pattern in Go](https://golangbyexample.com/golang-factory-design-pattern/) from [Golang By Example](https://golangbyexample.com).
- [Design Pattern - Factory Pattern](https://www.tutorialspoint.com/design_pattern/factory_pattern.htm) from [tutorialspoint](https://www.tutorialspoint.com).
- [Factory Method](https://refactoring.guru/design-patterns/factory-method) from [Refactoring Guru](https://refactoring.guru).
