# Chain of Responsibility
`Chain of Responsibility` is a behavioral design pattern that allows separated handlers to run in a kind of chain. For every request the chain execute the handlers one by one where each handler executes some action, quit or continue the chain by calling the next handler.

The implementation here is based on the [Source Making](https://sourcemaking.com/design_patterns/chain_of_responsibility) example where an ATM withdraw is handled by it's money cassettes. In this example the cassette `handlers` had repeated code intentionally to show that each `handler` can do it's own implementation.

![ATM withdraw - https://sourcemaking.com](https://sourcemaking.com/files/sm/images/patterns/Chain_of_responsibility_example.png?id=d86e4ad550281e49f387)

**The diagram**

![Chain Of Responsibility Diagram](https://raw.githubusercontent.com/marcelkohl/go-design-patterns/main/ChainOfResponsibility/diagram.png)

1. The `Handler` (here as `casseteInterface`) declares the common aspects for all concrete handlers (`casseteN`). Usually it contains a single method for handling requests, but sometimes it may also have another method for setting the next handler on the chain.
2. The `Base Handler` is an optional implementation where is set common code to all other handler implementations, as for example the method `setNext`.


## Running
```
go run .
```

## Copyright and References
Some content here is based on:

- [Chain of Responsibility Pattern](https://www.tutorialspoint.com/design_pattern/chain_of_responsibility_pattern.htm) from [Tutorials Point](https://www.tutorialspoint.com).
- [Chain of Responsibility Design Pattern in Go](https://golangbyexample.com/chain-of-responsibility-design-pattern-in-golang//) from [Golang By Example](https://golangbyexample.com).
- [Chain of Responsibility](https://sourcemaking.com/design_patterns/chain_of_responsibility) from [Source Making](https://sourcemaking.com).
- [Chain of Responsibility](https://refactoring.guru/design-patterns/chain-of-responsibility) from [Refactoring Guru](https://refactoring.guru)
- [Inheritance in GoLang](https://golangdocs.com/inheritance-in-golang) from [GODocs](https://golangdocs.com)
