# Command
`Command` is a pattern that encapsulates a request as an object, thereby letting us parameterize other objects with different requests, queue or log requests, and support undoable operations.

![Command Diagram](https://raw.githubusercontent.com/marcelkohl/go-design-patterns/main/Command/diagram.png)

1. `Request` is wrapped under an object as a `command` and passed to `invoker` object.
2. `Invoker` object looks for the appropriate object which can handle the `command` and passes the `command` to the corresponding object which executes the command.

## Running
```
go run .
```

## Copyright and References
Some content here is based on:

- [Command Pattern](https://www.geeksforgeeks.org/command-pattern/) from [GeeksforGeeks](https://www.geeksforgeeks.org)


- [Chain of Responsibility Pattern](https://www.tutorialspoint.com/design_pattern/chain_of_responsibility_pattern.htm) from [Tutorials Point](https://www.tutorialspoint.com).
- [Chain of Responsibility Design Pattern in Go](https://golangbyexample.com/chain-of-responsibility-design-pattern-in-golang//) from [Golang By Example](https://golangbyexample.com).
- [Chain of Responsibility](https://sourcemaking.com/design_patterns/chain_of_responsibility) from [Source Making](https://sourcemaking.com).
- [Chain of Responsibility](https://refactoring.guru/design-patterns/chain-of-responsibility) from [Refactoring Guru](https://refactoring.guru)
- [Inheritance in GoLang](https://golangdocs.com/inheritance-in-golang) from [GODocs](https://golangdocs.com)
