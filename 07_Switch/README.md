# Go's switch Statement

The switch statement is a cleaner and often more powerful way than using long chains of if-else if-else statements. It allows you to execute different code blocks based on the value of an expression.

---

## 1. Basic switch (Value Comparison)

This is the standard form where you compare a single expression against multiple potential values (case).

Syntax:

```go
var i = 3

switch i { // switch will evaluate the i variable
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
case 3:
    fmt.Println("Three")
default:
    fmt.Println("More than Three")
}
```

Key Differences from other Languages:
No break needed: In Go, once a case is matched, the program automatically exits the switch block. You don't need to write a break keyword.

Optional default: The default case is executed if none of the case values match, but it is entirely optional.

## 2. Multiple Values in a Single case

Go allows you to list multiple values for a single case, separated by commas.

Example: In the active code, the switch checks the current day of the week.

```go
switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
    fmt.Println("It's weekend")
default:
    fmt.Println("It's workday")
}
// If time.Now().Weekday() is either Saturday or Sunday, the code runs and prints "It's weekend".
```

## 3. switch for Checking Type (type switch)

This is one of the most powerful uses of switch in Go. Instead of checking a value, you can check the data type of a variable (specifically an interface{}).

interface{} means any type

It matches the default case because 10.5 is a float64.
Output: "Other"

```go
whoAmI := func(i interface{} ) {
    switch i.(type) {
    case int:
        fmt.Println("Integer")
    case string:
        fmt.Println("String")
    case bool:
        fmt.Println("Boolean")
    default:
        fmt.Println("Other")
    }
}

whoAmI(10.5)
```
