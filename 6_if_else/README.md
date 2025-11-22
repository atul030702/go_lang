# Go's Conditional Statements (if, else if, else)

In Go, you use conditional statements to execute different blocks of code based on whether a condition is true or false.

---

## 1. Basic if and else

The code inside the if block runs if the condition is true; otherwise, the code inside the else block runs.
Syntax: You don't need parentheses around the condition.

```go
if age >= 18 {
    fmt.Println("Person is an adult")
} else {
    fmt.Println("Person is minor.")
}
```

## 3. Declaring Variables Inside if (The Idiomatic Go Way)

The final, active code block shows a unique and recommended Go pattern: declaring a variable right inside the if statement.
Semicolon (;): Separates the variable declaration from the condition.

Scope: The variable (age in this case) is local and only exists within the if block, and any associated else if or else blocks. It cannot be accessed outside of this entire conditional structure. This keeps your variables clean and contained!

```go
if age := 15; age >= 18 {
    fmt.Println("Person is an adult", age)
} else if age >= 12 {
    fmt.Println("Person is teenager", age)
} else {
    fmt.Println("Person is a child", age)
}
```

## 4. No Ternary Operator

Go does not have a ternary operator (? :). You must use the standard if-else structure to achieve conditional assignments.
