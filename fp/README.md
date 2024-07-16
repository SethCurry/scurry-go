# fp

`fp` is a micro-library that leverages generics to provide some functional
programming features in Go.

## Examples

### Map

```go
items := []int{1, 2, 3}

squared := fp.Map(func(i int) { return i*i}, items)

// squared == [1, 4, 9]
```

### Filter

```go
items := []int{10, 50, 100}

filtered := fp.Filter(func(i int) { return i > 20 }, items)

// filtered == [50, 100]
```

### Reduce

```go
items := []int{1, 3, 4}

sum := fp.Reduce(func(a int, b int) { return a + b }, items)

// sum == 8
```
