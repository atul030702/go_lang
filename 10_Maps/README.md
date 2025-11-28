# Maps: Key-Value Storage

A Map in Go is an unordered collection of key-value pairs. They are similar to objects in JavaScript or dictionaries in Python. Maps provide a way to look up a value by its unique key.

---

## 1. Creating a Map

You must specify the type for the key and the type for the value when creating a map.
Syntax: `map[<KeyType>]<ValueType>`

```go
    var map1 = map[string]int{"price": 40, "items": 23} // key = string, value = integer

    /* Using make() method */
    var map2 = make(map[string]string) // key = string, value = string
```

## 2. Adding and Accessing Elements

**Adding/Setting Elements**  
To add or update an element, use the bracket notation `[key]`.

`map2["name"] = "golang"`: Sets the value `"golang"` for the key `"name"`.

**Getting Elements**  
To retrieve a value, use the same bracket notation.

`fmt.Println(map2["name"])` will print `"golang"`.

**The Zero Value Issue**  
If you try to retrieve a value using a key that doesn't exist in the map, Go returns the zero value for the map's value type, not an error.

| Map Value Type           | Zero Value          | Example                                 |
|--------------------------|---------------------|-----------------------------------------|
| `int (integers)`         | `0`                 | `map1["phone"]` returns 0               |
| `string (Text)`          | `""` (Empty String) | `map2["phone"]` returns ""              |
| `bool (Booleans)`        | `false`             | `map2["phone"]` returns false           |

## 3. Checking for Key Existence (The Comma-Ok Idiom)

Since trying to access a non-existent key returns a zero value (which might be a valid value, like `0` or `""`), Go provides a special two-value return, known as the "comma-ok" idiom, to safely check if a key exists.

Syntax: `value, ok := mapName[key]`

| Component               | Purpose                                                                            |
|-------------------------|------------------------------------------------------------------------------------|
| `value`                 | The value associated with the key (or the zero value if the key doesn't exist).    |
| `ok`                    | A boolean (`true` or `false`) indicating if the key was found in the map.          |

```go
    _, ok := map3["phones"] // We ignore the value with _ and just check the boolean 'ok'
    if ok {
        // key "phones" exists!
    } else {
        // key "phones" does not exist!
    }
```

## 4. Common Map Operations

| Operation                | Function/Keyword       | Explanation                                                    |
|--------------------------|------------------------|----------------------------------------------------------------|
| Length                   | `len(map3)`            | Returns the number of key-value pairs in the map.              |
| Delete                   | `delete(m, "key")`     | Removes the pair associated with the given key.                |
| Clear                    | `clear(map3)`          | Removes all elements, making the map empty. (Go 1.21+ only)    |

## 5. Comparing Maps (`maps.Equal`)

Maps are a reference type, and you cannot use the == operator to compare them directly. You must use the `maps.Equal` function from the maps package (available since Go 1.21).

- `maps.Equal(m1, m2)` returns true only if both maps have the same number of keys, the same keys, and the corresponding values are equal.

| Code                                 | Explanation                                                     |
|--------------------------------------|-----------------------------------------------------------------|
| `m1 := {"price": 40, "phones": 3}`   | Map 1 has keys and values.                                      |
| `m2 := {"price": 40, "phones": 8}`   | Map 2 has the same keys, but different value for `"phones"`.    |
| `fmt.Println(maps.Equal(m1, m2))`    | Output: `false` (because `m1["phones"] != m2["phones"]`).       |
