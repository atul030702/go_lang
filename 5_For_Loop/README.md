# Go `for` Loop Basics

Go has only **one** looping keyword: `for`.  
With it, you can write while-loops, infinite loops, classic loops, and range loops.

---

## 1. Using `for` like a while-loop

You skip the init and post statements.  
The loop runs as long as the condition is true.

```go
i := 1
for i <= 3 {
    fmt.Println(i)
    i++
}
```

## 2. Infinite loop

Leave everything empty.
Useful for servers or long-running processes.

```go
for {
    println("1")
}
```

## 3. Classic C-style loop

The familiar init; condition; update pattern.
continue skips the current iteration and jumps to next i++ iteration.

```go
for i := 0; i <= 3; i++ {
    if i == 2 {
        continue
    }
    fmt.Println(i)
}
```

## 4. range with an integer (Go 1.22+)

This style is a simplified way to loop a specific number of times, available in Go modules.

The loop runs from 0 up to, but not including, the specified 3.

```go
for i := range 3 {
    fmt.Println(i)
}
// Output: 0, 1, 2
```
