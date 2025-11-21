# Go Variable Declaration Fundamentals

Go is statically typed, but it leans heavily on type inference. The compiler often figures out the type from the value you assign, which feels familiar if you're coming from JavaScript.

---

## 1. Standard `var` Declaration

`var` is used when you want more control over the declaration, especially when:

- Declaring variables at the package (global) level  
- Declaring a variable without an initial value (zero value)  
- Being explicit about the type (e.g., `float32` vs `float64`)

### Syntax Variants

| Syntax                   | Description                                  | Example                       |
|--------------------------|----------------------------------------------|-------------------------------|
| `var name type = value`  | Explicit type and initialization             | `var name string = "Go-Lang"` |
| `var name = value`       | Type inference                               | `var my_name = "Atul"`        |
| `var name type`          | Zero value with declared type                | `var name string`             |

### Example: Standard `var`

```go
package main

import "fmt"

func main() {
    // 1. Explicit type and initialization
    var age int = 23
    var price float32 = 50.5

    // 2. Type inference
    var my_name = "Atul"
    var isAdult = true

    // 3. Zero value
    var location string
    location = "Singapore"

    fmt.Println("Name:", my_name, "Age:", age, "Adult:", isAdult, "Price:", price, "Location:", location)
}
```

---

## 2. Short Variable Declaration (`:=`)

The shorthand `:=` is the most idiomatic way to declare and initialize variables inside functions.

Key points:

- Only works inside functions  
- Always infers type  
- Must declare a new variable  
- Concise and widely used  

### Example: Shorthand `:=`

```go
package main

import "fmt"

func main() {
    name := "go lang"
    count := 5
    rate := 50.5 // float64 by default

    fmt.Println("Name:", name, "Count:", count, "Rate:", rate)
}
```
