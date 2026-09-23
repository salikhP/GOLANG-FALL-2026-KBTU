# Lecture 04 - Structs and Methods

## Learning Outcomes

- Declare and initialize structs in Go
- Access and modify struct fields
- Explain how structs behave when assigned and passed to functions
- Distinguish between struct values and pointers to structs
- Define methods on named types
- Explain the role of a method receiver
- Distinguish value receivers from pointer receivers
- Choose an appropriate receiver type for a method
- Understand how exported and unexported struct fields and methods work
- Use constructor-style functions to create and validate struct values

## Key Ideas

- A struct groups multiple fields into a single value
- `type User struct {...}` defines a new named type
- Struct fields can have different types
- Struct fields receive their zero values when not explicitly initialized
- Struct fields are accessed using the `.` operator
- Struct assignment copies the struct value
- Passing a struct to a function passes a copy of the struct
- Passing a pointer to a struct allows a function to modify the original value
- Go automatically dereferences struct pointers when accessing fields with `.`
- A method is a function associated with a named type through a receiver
- The receiver appears between `func` and the method name
- A value receiver receives a copy of the value
- A pointer receiver can modify the original value
- Pointer receivers are commonly used when a method modifies the receiver
- Pointer receivers can also avoid copying large structs
- Receiver choice should generally be consistent across methods of the same type
- Methods can be defined on locally defined named types, not only structs
- Go does not have classes; structs, methods, interfaces, and composition are used instead
- Identifiers starting with an uppercase letter are exported from their package
- Identifiers starting with a lowercase letter are unexported
- Go has no special constructor syntax
- Functions such as `NewUser` are normal functions commonly used as constructors
- A struct is comparable only when all of its fields are comparable
- Copying a struct does not necessarily create a deep copy of reference-like fields such as slices or maps

## Materials

- [Lecture board: Group 1](./Lecture-4-G1.pdf)
- [Lecture board: Group 2](./Lecture-4-G2.pdf)

## References

- [A Tour of Go: Structs](https://go.dev/tour/moretypes/2)
- [A Tour of Go: Struct Fields](https://go.dev/tour/moretypes/3)
- [A Tour of Go: Pointers to Structs](https://go.dev/tour/moretypes/4)
- [A Tour of Go: Struct Literals](https://go.dev/tour/moretypes/5)
- [A Tour of Go: Methods](https://go.dev/tour/methods/1)
- [A Tour of Go: Pointer Receivers](https://go.dev/tour/methods/4)
- [A Tour of Go: Methods and Pointer Indirection](https://go.dev/tour/methods/5)
- [Go Specification: Struct types](https://go.dev/ref/spec#Struct_types)
- [Go Specification: Method declarations](https://go.dev/ref/spec#Method_declarations)
- [Effective Go: Constructors and composite literals](https://go.dev/doc/effective_go#composite_literals)
- [Effective Go: Pointers vs. Values](https://go.dev/doc/effective_go#pointers_vs_values)
