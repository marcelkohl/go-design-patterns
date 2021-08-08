# Builder Pattern
Builder pattern builds a complex object using simple objects and using a step by step approach.
The pattern allows you to produce different types and representations of an object using the same construction code

![Alt text](https://raw.githubusercontent.com/marcelkohl/go-design-patterns/main/Builder/diagram.png)

1. **Builder interface** declares product construction steps that are common to all types of builders.
2. **Concrete Builders** provide different implementations of the construction steps.
3. **Products** are the results from the builders. Different product's builders don’t have to belong to the same class hierarchy or interface.
4. **Director** defines the sequence in which to call builder steps.

## Running
```
go run .
```

## Copyright
Some content here is based on:

- [Builder Pattern in Go](https://golangbyexample.com/builder-pattern-golang/) from [Golang By Example](https://golangbyexample.com).
- [Abstract Factory](https://refactoring.guru/design-patterns/builder) from [Refactoring Guru](https://refactoring.guru).
