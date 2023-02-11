# Loops in Go

Loops are a fundamental part of any programming language. Go has three types of loops: for, while, and do-while. The for loop is the most common and is used in most cases. The while and do-while loops are rarely used in Go.

## For Loop
The for loop is the most common loop in Go. It is used to repeat a block of code a certain number of times.The condition is evaluated before every iteration of the loop. If the condition is true, the loop body is executed. The post statement is executed after every iteration of the loop. The loop body is executed as long as the condition is true.

The init statement is optional. The condition is required. The post statement is optional.

```go
    for init; condition; post {
        // loop body
    }
```
Example
```go
    package main

    import "fmt"

    func main() {
        for i := 0; i < 5; i++ {
            fmt.Println(i)
        }
    }
```

## While Loop
The while loop is rarely used in Go. It is used to repeat a block of code as long as a condition is true. The condition is evaluated before every iteration of the loop. If the condition is true, the loop body is executed. The loop body is executed as long as the condition is true.

```go
    for condition {
        // loop body
    }
```
Example
```go
    package main

    import "fmt"

    func main() {
        i := 0
        for i < 5 {
            fmt.Println(i)
            i++
        }
    }
```

## Do-While Loop

The do-while loop is rarely used in Go. It is used to repeat a block of code as long as a condition is true. The condition is evaluated after every iteration of the loop. If the condition is true, the loop body is executed. The loop body is executed as long as the condition is true.

```go
    for {
        // loop body
        if condition {
            break
        }
    }
```
Example
```go
    package main

    import "fmt"

    func main() {
        i := 0
        for {
            fmt.Println(i)
            i++
            if i >= 5 {
                break
            }
        }
    }
```