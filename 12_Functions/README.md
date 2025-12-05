# Functions: Definitions and Basic Use

Functions are the building blocks of Go programs. They package code to perform a specific task.

---

## 1. Simple Function Declaration

To define a function, you use the func keyword, specify parameter names and types, and define the return type.

- Syntax: `func <name>(<param1> <type1>, <param2> <type2>) <returnType> { ... }`

| Code                              | Explanation                                                                                                                   |
|-----------------------------------|-------------------------------------------------------------------------------------------------------------------------------|
| `func add(a, b int) int { ... }`  | Declares a function named `add`. It takes two parameters, `a` and `b`, both of type `int`. It returns a single value of type int. |
| `return a + b`                    | Calculates the sum and sends the result back to where the function was called.                                                |

- Notice that if consecutive parameters share the same type (like a and b being int), you only need to declare the type once at the end.

## 2. Multiple Return Values (A Go Specialty)

Unlike many other languages, Go functions can easily return multiple values. This is commonly used for returning a result alongside an error status.

- Syntax: You must define the types of all returned values inside parentheses: `func <name>() (<type1>, <type2>, <type3>) { ... }`

```go
func getLanguages() (string, string, bool) {
    return "golang", "javascript", true
}

// first string = lang1, second string = _(discarded for no use), third boolean = lang3
lang1, _, lang3 := getLanguages()
```

## 3. First-Class Functions

In Go, functions are considered "first-class citizens." This means you can treat functions like any other variable or data type.

- **Storing Functions in Variables**:  
You can define an anonymous function (a function without a name) and store it directly in a variable.
- **Passing Functions as Parameters**:  
A function can accept another function as an argument.
- **Returning Functions from Functions**:  
A function can also return another function as its result.
