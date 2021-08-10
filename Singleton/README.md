# Singleton
`Singleton` is a pattern that ensures that an instantiated object has only one instance, and provide a common point of access to it.

`Singleton` can be encapsulated "just-in-time initialization" or "initialization on first use".

![Alt text](https://raw.githubusercontent.com/marcelkohl/go-design-patterns/main/Singleton/diagram.png)

1. The object of the single instance is responsible to give access and "initialize on first use".
2. The single instantiated content is a private static attribute. The method to access the instance is a public static method.
3. The `Singleton` class declares the static method `getInstance` that returns the same instance of its own class.

## Running
```
go run .
```

## Copyright and References
Some content here is based on:

- [Singleton Design Pattern in Go](https://golangbyexample.com/singleton-design-pattern-go/) from [Golang By Example](https://golangbyexample.com).
- [Singleton Design Pattern](https://sourcemaking.com/design_patterns/singleton) from [Source Making](https://sourcemaking.com).
