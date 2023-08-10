# notc

A dynamic language that puts C to shame. (jk, I built this to understand how langauges are designed from the ground up)

## Quick Start

To utilize the REPL, run `./notc`, or provide it a file.

```shell
$ git clone https://github.com/EhsaanIqbal/notc
$ go build
$ ./notc examples/basic.notc
```

## The Language

Explore: [examples](./examples)

### Types

notc supports the following data types: `null`, `bool`, `int`, `str`, `array`,
and `fn`. The `int` type signifies a signed 64-bit integer, while strings are
immutable arrays of bytes, and arrays can dynamically grow.

### Variable Bindings

```shell
>> let x = 10
```

### Math

```shell
>> let x = 10
>> let y = x * 2
>> (x + y) / 2 - 3
12
```

### Conditional Expressions

notc supports `if` and `else`:

```shell
>> let x = 10
>> let y = x * 2
>> if (x > y) { print("x is greater") } else { print("y is greater") }
x is greater
```

### Functions and Closures

You can define named or anonymous functions, including functions within
functions that reference outer variables (known as _closures_).

```shell
>> multiply := fn(x, y) { x * y }
>> multiply(50 / 2, 1 * 2)
50
>> fn(x) { x + 10 }(10)
20
>> newAdder := fn(x) { fn(y) { x + y } }
>> addTwo := newAdder(2)
>> addTwo(3)
5
>> sub := fn(a, b) { a - b }
>> applyFunc := fn(a, b, func) { func(a, b) }
>> applyFunc(10, 2, sub)
8
```

### Strings

```shell
>> let x = "hello"
>> print(x + "mars")
hello mars
```

### Arrays

```shell
>> myArray := ["x", "y", 1, fn(x) { x * x }]
>> myArray[0]
x
>> myArray[4 - 2]
1
```

### Builtin Functions

- `len(iterable)`
  Returns the length of the iterable (`str`, `array`).
- `print(value...)`
  Outputs the `value`(s) to standard output followed by a newline.
- `push(array, value)`
  Returns a new array with `value` added to the end of `array`.
