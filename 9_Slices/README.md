# Go Slices: Dynamic and Flexible

- A slice is a flexible, dynamic view over an array.
- You can think of it as a “dynamic array”:
- It has a length (number of elements currently stored).
- It has a capacity (how many elements can fit before Go needs to grow it internally).

Slices are used much more often than arrays in real Go code.

---

## 1. Uninitialized Slices (The nil Slice)

When you declare a slice without initializing it, it has a special value called nil.

| Code                               | Explanation                                                  |
|------------------------------------|--------------------------------------------------------------|
| `var num []int`                    | Declares a slice of integers.                                |
| `fmt.Println(num == nil)`          | Output: true. A newly declared, uninitialized slice is nil.  |
| `fmt.Println(len(num))`            | Output: 0. The slice is empty.                               |

## 2. Creating Slices with `make()` (Length and Capacity)

The built-in `make()` function is the standard way to create a slice, allowing you to specify its initial length and capacity.

Length (`len()`): The number of elements currently in the slice.
Capacity (`cap()`): The maximum number of elements the slice can hold without reallocating the underlying array.

- Example: `var num = make([]int, 0, 5)`
  - Initial Length: 0 (It starts empty).
  - Capacity: 5 (It can hold up to 5 elements before needing a larger underlying array).

## 3. Adding Elements with `append()`

Since slices are dynamic, you add elements using the built-in `append()` function.

- **Capacity Growth**: If you append an element and the current capacity is exceeded, Go automatically creates a new, larger underlying array (usually double the size of the previous capacity) and copies all the old elements into it. If you append 6 elements to a capacity-5 slice, the capacity will jump to 10 (or sometimes 12, depending on the Go version and size).

| Code                          | Explanation                                   |
|-------------------------------|-----------------------------------------------|
| `num = append(num, 1)`        | Adds 1 to the slice. The length becomes 1.    |
| `fmt.Println(cap(num))`       | If capacity was 5, it remains 5.              |

## 4. Copying Slices with `copy()`

The `copy()` function is used to copy elements from a source slice to a destination slice.

```go
    copy(destination, source)

    var num1 = make([]int, 0, 5) // Initial length is 0, capacity is 5
    num1 = append(num, 2)

    var num2 = make([]int, len(num))

    copy(num2, num)

    fmt.Println(num, num2)
```

- Copies elements from source slice to destination slice.
- Copies up to min(`len(source)`, `len(destination)`).
- destination must already have enough length.

Important: It only copies as many elements as the shortest of the two slices can hold.

## 5. Slice Operator (Slicing an Array or Slice)

You can create a new slice from a portion of an existing array or slice using the slice operator. This is similar to indexing but uses a range.

```go
    num[start_index:end_index]

    var num = []int{1, 2, 3, 4, 5}
    num[0:2] // output: [1, 2], elements at index 0 and 1
    num[:2] // output: [1, 2], elements at index 0 and 1 (missing start_index defaults to 0)
    num[1:] // output: [2, 3, 4, 5], elements from index 1 to end (missing end_index defaults to length of slice = 5)
```

The `start_index` is inclusive.
The `end_index` is exclusive (the element at this index is not included).

## 6. Multi-Dimensional Slices (2D Slices)

You can create slices of slices, which are used to represent structures like grids or tables.

```go
    var num = [][]int{ {<row1>}, {<row2>} }

    var num = [][]int{ {1, 2, 3}, {4, 5, 6} }
```

## 7. What is the slices Package?

The slices package, introduced in Go 1.21, provides a set of generic, efficient functions for working with Go slices.

Before this package existed, many common operations on slices (like checking if two slices are equal or searching for an element) required writing custom loops every time. The slices package standardizes these operations.

```go
    var num1 = []int{1, 2, 3}
    var num2 = []int{1, 2, 4}
    fmt.Println(slices.Equal(num1, num2)) // Output: true
```

Some key functions of slices are:

- `slices.Equal(slice1, slice2)` (equality check)
- `slices.Index(slice, value)`, Returns index of value in slice else -1 if not found
- `slices.Reverse(s)` Reverses the order of elements in slice s. (`s = {1, 2, 3}` becomes `s = {3, 2, 1}`)
- `slices.Sort(s)`, Sorts the elements of slice s in ascending order. (`s = {3, 1, 2}` becomes `s = {1, 2, 3}`)
- `slices.Clone(s)`, Returns a brand new slice with the same elements as s. (This is important to avoid accidentally modifying the original slice.)
