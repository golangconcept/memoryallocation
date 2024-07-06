Go has multiple ways of memory allocation and value initialization:

&T{...}, &someLocalVar, new, make

Allocation can also happen when creating composite literals.

new can be used to allocate values such as integers, &int is illegal:

```go
new(Point)
&Point{}      // OK
&Point{2, 3}  // Combines allocation and initialization

new(int)
&int          // Illegal

// Works, but it is less convenient to write than new(int)
var i int
&i
```
The difference between new and make can be seen by looking at the following example:

```go
p := new(chan int)   // p has type: *chan int
c := make(chan int)  // c has type: chan int
Suppose Go does not have new and make, but it has the built-in function NEW. Then the example code would look like this:

p := NEW(*chan int)  // * is mandatory
c := NEW(chan int)
The * would be mandatory, so:

new(int)        -->  NEW(*int)
new(Point)      -->  NEW(*Point)
new(chan int)   -->  NEW(*chan int)
make([]int, 10) -->  NEW([]int, 10)

make(Point)  // Illegal
make(int)    // Illegal
```

Yes, merging new and make into a single built-in function is possible. However, it is probable that a single built-in function would lead to more confusion among new Go programmers than having two built-in functions.

Considering all of the above points, it appears more appropriate for new and make to remain separate.

> new(T) - Allocates memory, and sets it to the zero value for type T..
..that is 0 for int, "" for string and nil for referenced types (slice, map, chan)

- Note that referenced types are just pointers to some underlying data structures, which won't be created by new(T)
Example: in case of slice, the underlying array won't be created, thus new([]int) returns a pointer to nothing


- make(T) - Allocates memory for referenced data types (**slice**, **map**, **chan**), plus initializes their underlying data structures

### Example:
In case of slice, the underlying array will be created with the specified length and capacity
Bear in mind that, unlike C, an array is a primitive type in Go!

That being said:
- make(T) behaves like composite-literal syntax
- new(T) behaves like var (when the variable is not initialized)
```go 
func main() {
    fmt.Println("-- MAKE --")
    a := make([]int, 0)
    aPtr := &a
    fmt.Println("pointer == nil :", *aPtr == nil)
    fmt.Printf("pointer value: %p\n\n", *aPtr)

    fmt.Println("-- COMPOSITE LITERAL --")
    b := []int{}
    bPtr := &b
    fmt.Println("pointer == nil :", *bPtr == nil)
    fmt.Printf("pointer value: %p\n\n", *bPtr)

    fmt.Println("-- NEW --")
    cPtr := new([]int)
    fmt.Println("pointer == nil :", *cPtr == nil)
    fmt.Printf("pointer value: %p\n\n", *cPtr)

    fmt.Println("-- VAR (not initialized) --")
    var d []int
    dPtr := &d
    fmt.Println("pointer == nil :", *dPtr == nil)
    fmt.Printf("pointer value: %p\n", *dPtr)
}
Run the program

-- MAKE --
pointer == nil : false
pointer value: 0x118eff0  # address to underlying array

-- COMPOSITE LITERAL --
pointer == nil : false
pointer value: 0x118eff0  # address to underlying array

-- NEW --
pointer == nil : true
pointer value: 0x0

-- VAR (not initialized) --
pointer == nil : true
pointer value: 0x0

```

> Note: <span style="color: red;">Problem with new() arises when it needs to handle three other composite types - chan, slice and map</span>. 