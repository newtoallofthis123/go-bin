# FMT Module in Go

The `fmt` module is used for formatting and printing data to the console. It is a part of the Go standard library, so we don't need to install it.

## Printing to the Console
To print to the console, we can use the `Println` function from the `fmt` module. Here is an example:

```go
    println("Hello World")
```
## Formatting Output
We can also format the output using the `Printf` function. Here is an example:

```go
    name := "John"
    age := 25
    rating := 4.73
    fmt.Printf("%s (%d) has a rating of %.2f\n", name, age, rating)
```
The `%s` is a placeholder for a string, `%d` is a placeholder for an integer, and `%.2f` is a placeholder for a float. The `\n` is a newline character.

More info [here](https://pkg.go.dev/fmt)

## Scanning Input
We can also scan input from the console using the `Scanln` function. Here is an example:

```go
    var name string
    fmt.Print("Enter name: ")
    fmt.Scanln(&name)
    fmt.Println("Hello", name)
```

Here we are using & to get the memory address of the variable `name`. This is required by the `Scanln` function.

## Summing Up
Simple Addition with fmt
    
```go
    var num1, num2 int
    fmt.Print("Enter num1: ")
    fmt.Scanln(&num1)
    fmt.Print("Enter num2: ")
    fmt.Scanln(&num2)
    fmt.Println("Sum is", num1+num2)
```