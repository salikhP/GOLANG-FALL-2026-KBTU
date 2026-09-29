# Lecture 05 - Interfaces and Generics

## Learning Outcomes

- Define and use interfaces in Go
- Explain implicit interface implementation
- Explain what dynamic type and dynamic value mean for an interface
- Understand how method sets affect interface implementation
- Distinguish between value and pointer types when satisfying interfaces
- Use type assertions to retrieve values from interfaces
- Use the comma-ok idiom for safe type assertions
- Use type switches to handle different dynamic types
- Explain the difference between a nil interface and an interface containing a nil pointer
- Understand the purpose of `any` and `interface{}`
- Define generic functions and generic types
- Use type parameters and constraints
- Understand type inference in generic function calls
- Use `comparable` as a constraint
- Define custom type-set constraints using `|`
- Explain the meaning of `~` in generic constraints
- Choose appropriately between interfaces and generics

## Key Ideas

- An interface defines a set of methods that describe behavior
- A type implements an interface implicitly by implementing all methods required by the interface
- Go does not use an `implements` keyword
- Interfaces allow code to depend on behavior rather than on a specific concrete type
- Interface values conceptually contain a dynamic type and a dynamic value
- The static type of an interface variable can differ from the concrete type stored inside it
- A type's method set determines which interfaces it satisfies
- Methods with value receivers belong to the method sets of both `T` and `*T`
- Methods with pointer receivers belong to the method set of `*T`, but not `T`
- Small interfaces are generally preferred because they are easier to implement and compose
- Interfaces can embed other interfaces to combine behaviors
- `any` is an alias for `interface{}`
- The empty interface can hold a value of any type because it requires no methods
- Type assertions use the syntax `value.(T)` to access a value of a specific dynamic type
- A failed single-value type assertion causes a runtime panic
- The comma-ok form `value, ok := x.(T)` safely checks whether an assertion succeeds
- Type assertions can check both concrete types and other interfaces
- Type switches use `value.(type)` to branch based on the dynamic type stored in an interface
- `value.(type)` can only be used inside a type switch
- A nil interface has neither a dynamic type nor a dynamic value
- An interface containing a typed nil pointer is not itself nil
- Returning a typed nil pointer as an `error` can produce a non-nil error interface
- When there is no error, functions returning `error` should normally return `nil`
- Generics allow functions and data structures to work with multiple types while preserving compile-time type safety
- A type parameter is declared inside square brackets, for example `[T any]`
- A constraint specifies which types and operations are allowed for a type parameter
- `any` allows any type to be used as a type argument
- Go can often infer type arguments from function arguments
- `comparable` allows types that support `==` and `!=`
- Type sets can be defined using unions such as `int | int64 | float64`
- `~T` includes types whose underlying type is `T`
- Generics are useful for reusable algorithms, collections, and type-safe utility functions
- Interfaces are useful when code depends on behavior and different implementations may be used
- Interfaces primarily provide runtime polymorphism, while generics provide compile-time type abstraction
- Generics should not automatically replace interfaces
- Interfaces should not be introduced when a concrete type already expresses the requirement clearly

## Materials

- [Lecture board: Group 1](./Lecture-5-G1.pdf)
- [Lecture board: Group 2](./Lecture-5-G2.pdf)

## References

- [A Tour of Go: Interfaces](https://go.dev/tour/methods/9)
- [A Tour of Go: Interfaces are implemented implicitly](https://go.dev/tour/methods/10)
- [A Tour of Go: Interface values](https://go.dev/tour/methods/11)
- [A Tour of Go: Interface values with nil underlying values](https://go.dev/tour/methods/12)
- [A Tour of Go: Nil interface values](https://go.dev/tour/methods/13)
- [A Tour of Go: Type assertions](https://go.dev/tour/methods/15)
- [A Tour of Go: Type switches](https://go.dev/tour/methods/16)
- [A Tour of Go: Generics](https://go.dev/tour/generics/1)
- [A Tour of Go: Type parameters](https://go.dev/tour/generics/2)
- [Go Specification: Interface types](https://go.dev/ref/spec#Interface_types)
- [Go Specification: Method sets](https://go.dev/ref/spec#Method_sets)
- [Go Specification: Type assertions](https://go.dev/ref/spec#Type_assertions)
- [Go Specification: Type switches](https://go.dev/ref/spec#Type_switches)
- [Go Specification: Type parameter declarations](https://go.dev/ref/spec#Type_parameter_declarations)
- [Go Specification: General interfaces](https://go.dev/ref/spec#General_interfaces)
- [Go Blog: An Introduction To Generics](https://go.dev/blog/intro-generics)
