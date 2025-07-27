# 📘 Golang DSA Cheat Sheet

A handy cheat sheet for mastering Go (Golang) fundamentals and data structures

---

## 🔢 Operators
```go
import (
	"math"
)

a + b      // addition
a - b      // subtraction
a * b      // multiplication
a / b      // division (int if both ints)
a % b      // modulo
// exponents
math.Pow(base, exponent) //  two float64 arguments
math.Pow10(4) // for base 10

a == b     // equal
a != b     // not equal
a > b, a < b, a >= b, a <= b  // comparisons

a && b     // logical AND
a || b     // logical OR
!a         // logical NOT

a & b      // bitwise AND
a | b      // bitwise OR
a ^ b      // bitwise XOR
a << 1     // left shift
a >> 1     // right shift

```

---

## 📜 Strings
```go
s := "hello"
len(s)        // 5
s[0]          // byte value of 'h'
s2 := s[:]    // copy of s, but with different address
```

- Strings are **immutable** and **byte-indexed**.
- For slicing or character access:
```go
s := "hello"
sub := s[1:4] // "ell"
```

### ✂️ Splitting
```go
import "strings"

strings.Split("a,b,c", ",")      // ["a", "b", "c"] split on any operator
strings.Fields("hi there")        // ["hi", "there"] split on any whitespace
```

### 🧵 Formatting
```go
fmt.Sprintf("%s is %d years old", "Bob", 25)  // "Bob is 25 years old"
```
- `%v`: default format
- `%+v`: struct with field names
- `%#v`: Go syntax representation
- `%.2f`: float precision
- `%t`: boolean
- `%q`: quoted string
---

## 🧱 Arrays
```go
var a [3]int       // [0, 0, 0]
a := [3]int{1, 2, 3}
```
- Fixed-size, not often used directly

---

## 🧩 Slices (like Python lists)
```go
arr := []int{1, 2, 3}
arr = append(arr, 4)
copyArr := make([]int, len(arr))
copy(copyArr, arr)
```

- `append`: add elements
- `copy`: shallow copy
- `make([]T, len)`: create slice with default values
- `sort.Ints(arr)` to sort (from `sort` package)

---

## 🔐 Maps (like dict)
```go
m := map[string]int{"a": 1, "b": 2} // declare a map
m2 := make(map[string]int) // initialize an empty map
// Set key/value pairs using typical name[key] = val syntax.
m2["k1"] = 7
m2["k2"] = 13
val := m["a"]
_, ok := m["c"]  // check existence
m["d"] = 4

delete(m, "b") // delete key
clear(m) //clears the entire map
fmt.Println("map:", m) //appear in the form map[k:v k:v] when printed with fmt.Println.
```

- `len(m)` gives number of key-value pairs
- No guaranteed order

---

## 🧬 Structs
```go
type Person struct {
  Name string
  Age  int
}

p := Person{"Alice", 30}
p.Age = 31

//anonymous struct type, common for table driven tests
dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }

```

- Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.
- Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it
- Structs are mutable.

---

## 📦 Tuples? (Go-style)
Go does **not have native tuples**, but we can return multiple values in a function:
```go
func coords() (int, int) {
  return 3, 4
}
x, y := coords()
```
Or, use `struct` for named tuple-like behavior.

---

## 🔁 Loops
```go
for i := 0; i < 10; i++ {}

for i, v := range []int{10, 20} {
  fmt.Println(i, v)
}

for condition {
  // while-style loop
}
```

---

## ✅ Conditionals
```go
if x > 0 {
} else if x == 0 {
} else {
}
```

---

## 🧠 Functions
```go
func add(a, b int) int {
  return a + b
}

func swap(a, b string) (string, string) {
  return b, a
}

func sum(nums ...int) int {
  total := 0
  
  for _, num := range nums {
    total += nums
  }

  return total
}

sum(1, 2, 3)
nums := []int{1, 2, 3, 4}
sum(nums...)
```

- Functions can return multiple values
- Named return values are optional
- Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
- If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.

Closures
```go
// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.
	nextInt := intSeq()

	// See the effect of the closure by calling `nextInt`
	// a few times.
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	newInts := intSeq() 
	fmt.Println(newInts()) // 1
}

```

---

## 🔃 Higher-Order Functions
```go
func apply(fn func(int) int, val int) int {
  return fn(val)
}
```
- Functions are first-class citizens in Go

---

## 🧷 Pointers
```go
x := 10
p := &x
*p = 20 // x is now 20
```

- Use `&` to get address, `*` to dereference

---

## 🔁 Recursion
```go
func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n-1)
}
```

- Similar to Python, but no max recursion depth by default

---

## ✅ Boolean & Truthy-ness
- Only `true` and `false`
- No implicit truthy/falsy values
- Conditionals must use a `bool` (not like Python’s `if mySlice:`)

---
