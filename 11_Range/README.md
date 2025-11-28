# Go's `range` Keyword: Simplified Iteration

The range keyword is Go's dedicated, clean way to iterate over data structures like slices, arrays, maps, and strings. It automatically handles indexing and access, which simplifies loop setup compared to manual index-based loops (for i = 0; i < len; i++).

The key feature of range is that it always returns two values (except when iterating over a channel):

- The Index/Key: The position (for slices/arrays) or the unique identifier (for maps).
- The Value: The element stored at that position/key.

---

## 1. Iterating Over Slices and Arrays (Index and Value)

When used on a sequential collection (like an array or a slice), range returns the index and the element value for each iteration.

```go
    nums := []int{6, 7, 8}

    for i, num := range nums {
        // i gets the 0-based index. num gets the element's value. Use _ to ignore the index
        fmt.Println(num, i) 
    }
```

## 2. Iterating Over Maps (Key and Value)

When used on a map, range returns the key and the value for each entry.

- **Order is Unspecified**: Remember that Go maps are unordered. The order in which key-value pairs are returned is not guaranteed to be the same between runs.

```go
    m := map[string]string{"first_name": "John", "last_name": "Doe"}

    for k, val := range m {
        // k is key, eg: "first_name", "last_name"
        fmt.Println(k) 

        // val is value for the key, eg: "John", "Doe"
        fmt.Println(val)
    }
```

## 3. Iterating Over Strings (Dealing with Unicode)

When iterating over a string, `range` intelligently handles Unicode characters (Runes).

| Returned Values          | What it Represents?                                               | Key Difference                          |
|--------------------------|-------------------------------------------------------------------|-----------------------------------------|
| `i` (Index)              | The starting byte index of the character.                         | Because some characters (like emojis or non-Latin letters) take up more than one byte, this index does not always increment by 1.              |
| `c` (Rune)               | The Unicode code point (integer representation of the character). | You must convert this integer value back to a string using $\mathbf{string(c)}$ to print the actual character.              |

```go
    for i, c := range "golang" {
        fmt.Println(i, c, string(c))
    }
```
