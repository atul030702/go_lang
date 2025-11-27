# Arrays: Fixed-Size Data Collections in Go

In Go, an array is a collection of elements of the same data type that have a fixed size. Once you declare an array, its size can never change.

---

## 1. Declaring and Initializing an Array

You must specify both the size (number of elements) and the data type when declaring an array.

```go
var nums [4]int // Declares an array named nums that can hold 4 elements, all of which must be integers (int).

var vals [4]bool // Declares an array named vals that holds 4 boolean values.

var name [3]string // Declares an array named name that holds 3 strings.
```

## 2. Default Values (Zero Values)

When you declare an array without assigning values, Go automatically fills it with the zero value for that specific data type.

| Data Type                | Zero Value          | Example                                                    |
|--------------------------|---------------------|------------------------------------------------------------|
| `int (integers)`         | 0                   | `var nums [4]int` results in `[0 0 0 0]`                   |
| `string (Text)`          | "" (Empty String)   | `var name [3]string` results in `["" "" ""]`               |
| `bool (Booleans)`        | false               | `var vals [4]bool` results in `[false false false false]`  |

## 3. Accessing and Modifying Elements

Arrays are zero-indexed, meaning the first element is at index $\mathbf{0}$, the second at index $\mathbf{1}$, and so on.

- Assigning a Value: `arrayName[index] = value`, example: `nums[0] = 1`
- Reading a Value: `fmt.Println(arrayName[index])`

## 4. Array Initialization in a Single Line

You can declare, size, and initialize an array all at once using a composite literal:

```go
nums := [3]int{1, 2, 3} // Declares a 3-integer array and immediately fills it with 1, 2, and 3.
```

## 5. Multi-Dimensional (2D) Arrays

Arrays can contain other arrays, creating structures like tables or grids.
Syntax:

```go
nums := [2][2]int{{3, 4}, {5, 6}}
```

- `[2][2]int` This array holds 2 arrays, and each of those inner arrays holds 2 integers.
- `{{3, 4}, {5, 6}}` The first inner array is `{3, 4}`, and the second is `{5, 6}`.

## 6. Key Characteristics of Go Arrays

Arrays are generally used when you know the exact, fixed size of your data set.

- Fixed Size: This is the most crucial characteristic.
- Constant Time Access: Getting any element (e.g., `nums[0]`) takes the same amount of time, regardless of the array's size (this is known as $O(1)$ complexity).

If you need a collection whose size can grow or shrink, you should use a Slice, which is built on top of the array structure.
