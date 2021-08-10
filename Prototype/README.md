# Prototype
Prototype pattern delegates the cloning process to the actual objects that are being cloned instead of the application being responsible to "copy" all the relevant attributes (public or private ones).

The prototype also may avoid the problem when an application "hard wires" the class of object to creation in each "new" expression.

![Alt text](https://raw.githubusercontent.com/marcelkohl/go-design-patterns/main/Prototype/diagram.png)

1. The `Prototype interface` declares the cloning methods
2. In the implementation, in addition to copying the original object’s data to the clone, this method may also cover some specific edge cases of the cloning process like: linked objects, untangling recursive dependencies, etc.

## Running
```
go run .
```

## Copyright and References
Some content here is based on:

- [Prototype Pattern in Go](https://golangbyexample.com/prototype-pattern-go/) from [Golang By Example](https://golangbyexample.com).
- [Prototype](https://refactoring.guru/design-patterns/prototype) from [Refactoring Guru](https://refactoring.guru).
- [Prototype Design Pattern](https://sourcemaking.com/design_patterns/prototype) from [Source Making](https://sourcemaking.com).
