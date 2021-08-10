# Object Pool or Resource Pools
Object pool may be used for performance boost where the cost of initializing an object is high.

An object pool also helps to manage available resources in a better way by handling an specific amount of objects. When an object is taken from the pool, it is not available in the pool until it is put back.

Some implementations may also add a facility to clean up unused objects periodically.

![Alt text](https://raw.githubusercontent.com/marcelkohl/go-design-patterns/main/ObjectPool/diagram.png)

1. **Client** requests and gives back resources the object pool using specific methods like `acquire` and `release`;
2. **Object Pool** controls the list of available and maximum limit of objects in the pool. The object pool may instantiate new objects as they are required.

## Running
```
go run .
```

## Copyright and References
Some content here is based on:

- [Object Pool Design Pattern in Go](https://golangbyexample.com/golang-object-pool/) from [Golang By Example](https://golangbyexample.com).
- [Object Pool Design Pattern](https://sourcemaking.com/design_patterns/object_pool) from [Source Making](https://sourcemaking.com).
- [Object Pool Pattern](https://www.javatpoint.com/object-pool-pattern) from [Java T point](https://www.javatpoint.com).
