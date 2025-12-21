# Understanding Pointers in Go

In Go, when you pass a variable to a function, Go usually creates a **copy** of that variable. If you want a function to modify the original variable, you need to use **Pointers**.

---

## 1. What is a Pointer?

A pointer is a variable that stores the **memory address** of another value. Instead of holding a value like `10` or `"Hello"`, it holds a coordinate like `0xc0000120b0` which points to where that value lives in your computer's RAM.

### Key Symbols

| Symbol | Name                     | What it does                                                        |
|--------|--------------------------|---------------------------------------------------------------------|
| `&`    | **Address Operator**     | Gets the memory address of a variable (e.g., `&num`).               |
| `*`    | **Dereference Operator** | Accesses the actual value stored at a memory address (e.g., `*num`).|

---

## 2. Pass by Value vs. Pass by Reference

### Pass by Value (The Default)

```go
func changeNum(num int) {
    num = 5
}
```

If you pass num here, Go makes a photocopy. If you draw on the photocopy, the original paper remains unchanged.

## 3. Pass by Reference (Using Pointers)

```go
func changeNum(num *int) { // Accepts a pointer to an int (*int)
    *num = 5 // "Follow the pointer" and change the value at that address to 5
}
```

Here, we pass the memory address. The function goes to that specific address and changes the value stored there.

## 4. The Main Execution

```go
func main() {
    num := 1

    // Use & to pass the memory address of 'num'
    changeNum(&num) 

    // This prints the address (e.g., 0xc0000120b8)
    fmt.Println("Memory Address", &num)

    // This prints 5, because changeNum modified the original variable
    fmt.Println("After changeNum in main", num)
}
```

---

## Summary

- Use `&` when you want to send the "map coordinates" (address) of a variable to a function.
- Use `*` in the function's arguments to say "I am expecting a memory address."
- Use `*` inside the function body to "reach inside" that address and change the original value.
