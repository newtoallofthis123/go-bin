# Variables in Go

Go has various value types including strings, integers, floats, booleans, etc. Variables are simply "named storage" that we can use to store these values in. We can create a variable using the `var` keyword, like so:
## Explicit variable declaration
```go
    var name string
```
This creates a variable called `name` which we can use to store a string value. We can then assign a value to this variable using the `=` operator, like so:
```go
    name = "John"
```
We can also create a variable and assign a value in one line, like so:
```go
    var name string = "John"
```
## Implicit variable declaration
Go is smart enough to infer the type of value that we are assigning to the variable, so we don't actually need to specify the type explicitly. We can rewrite the last line as:
```go
    var name = "John"
```

## Short variable declaration
We can also use the `:=` operator to create and assign a value to a variable in one line. This is known as the short variable declaration. We can rewrite the last line as:
```go
    name := "John"
```
The `:=` operator can only be used inside a function. It cannot be used at the package level.

## Multiple variable declaration
We can also declare multiple variables in one line, like so:
```go
    var firstName, lastName string
```

## Common Variable Types
Go has various value types including strings, integers, floats, booleans, etc. Here are some common variable types and their default values:
```go
    var name string // ""
    var rating float64 // 0
    var isCool bool // false
    var size int8 // 0
    var num uint //0->positive
    var num int //0->positive and negative
```

## Zero Values
When we create a variable, it is given a default value. For example, if we create a string variable, it will be given an empty string as its default value. If we create an integer variable, it will be given the value `0`. This is known as the "zero value" for that type. Here are some common zero values:
```go
    "" // string
    0 // int, float, uint
    false // bool
    nil // pointers, functions, interfaces, slices, channels, maps
```

## Summing Up
We can then assign values to these variables in multiple ways:
```go
    firstName = "John"
    lastName = "Doe"
```
```go
    firstName, lastName = "John", "Doe"
```
```go
    var firstName, lastName = "John", "Doe"
```
```go
    firstName, lastName := "John", "Doe"
```