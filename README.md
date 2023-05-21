# notc

A dynamic langauge that puts C to shame. (jk. I built this to understand how langauges are designed from the ground up)

## Quick start

Run `./notc` to use the REPL or give it a file

```#!sh
$ git clone https://github.com/EhsaanIqbal/notc
$ go build
$ ./notc examples/basic.notc
```


## The Language

> Check out: [examples](./examples)

### Types

notc supports the following data types: `null`, `bool`, `int`, `str`, `array`,
and `fn`. The `int` type is a signed 64-bit integer, strings are
immutable arrays of bytes, arrays are grow-able arrays.

### Variable Bindings

```#!sh
>> let x = 10
```

### Math

```#!sh
>> let x = 10
>> let y = x * 2
>> (x + y) / 2 - 3
12
```

### Conditional Expressions

notc supports `if` and `else`:

```sh
>> let x = 10
>> let y = x * 2
>> if (x > y) { print("x is greater") } else { print("y is greater") }
>>
x is greater
```

### Functions and Closures

You can define named or anonymous functions, including functions inside
functions that reference outer variables (_closures_).

```sh
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

```sh
>> let x = "hello"
>> print(x + "mars")
hello mars
```

### Arrays

```sh
>> myArray := ["x", "y", 1, fn(x) { x * x }]
>> myArray[0]
x
>> myArray[4 - 2]
1
```

### Builtin functions

- `len(iterable)`
  Returns the length of the iterable (`str`, `array`).
- `print(value...)`
  Prints the `value`(s) to standard output followed by a newline.
- `push(array, value)`
  Returns a new array with `value` pushed onto the end of `array`.
