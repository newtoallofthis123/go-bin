# Loops in Go

Loops are a fundamental part of any programming language. Go has three types of loops: for, while, and do-while. The for loop is the most common and is used in most cases. The while and do-while loops are rarely used in Go.

## For Loop
The for loop is the most common loop in Go. It is used to repeat a block of code a certain number of times.The condition is evaluated before every iteration of the loop. If the condition is true, the loop body is executed. The post statement is executed after every iteration of the loop. The loop body is executed as long as the condition is true.

The init statement is optional. The condition is required. The post statement is optional.

```go
    for(init; condition; post) {
        // loop body
    }
```

