# Lecture 06 - Functions, Errors, Defer, Panic and Recover

## Learning Outcomes

- Define functions with parameters and return values
- Use multiple return values
- Understand functions as first-class values
- Pass functions as arguments and return functions from functions
- Define and use anonymous functions
- Explain how closures capture variables from their surrounding scope
- Understand the built-in `error` interface
- Explain the idea of errors as values in Go
- Create errors using `errors.New`
- Create formatted errors using `fmt.Errorf`
- Handle errors using the idiomatic `if err != nil` pattern
- Propagate errors to callers
- Add context to errors using error wrapping
- Understand the difference between `%w` and `%v` when formatting errors
- Inspect wrapped errors using `errors.Is`
- Extract typed errors using `errors.As`
- Use `defer` for cleanup operations
- Explain when deferred function calls are executed
- Understand the LIFO execution order of multiple deferred calls
- Explain when deferred function arguments are evaluated
- Understand how `panic` changes normal program execution
- Explain stack unwinding during a panic
- Understand that deferred functions still execute during stack unwinding
- Use `recover` to intercept a panic
- Explain the restrictions on where `recover` can be used effectively
- Distinguish normal error handling from `panic` and `recover`
- Choose appropriately between returning an error and causing a panic

## Key Ideas

- Functions in Go can accept zero or more parameters and return zero or more values
- Parameters of the same type can share a type declaration, for example `func add(a, b int) int`
- Multiple return values are commonly used to return both a result and an error
- A common Go function signature is `(value, error)`
- Functions are values and can be assigned to variables
- Functions can be passed as arguments to other functions
- Functions can be returned from other functions
- Anonymous functions are functions declared without a name
- A closure is a function that refers to variables from its surrounding lexical scope
- A closure can continue to access captured variables even after the surrounding function has returned
- Closures capture variables rather than simply copying their current values
- Closures are commonly used for callbacks, configuration, middleware, handlers, and maintaining private state
- `error` is a built-in interface used to represent failures
- Conceptually, the built-in error interface is `interface { Error() string }`
- Errors in Go are ordinary values
- Expected runtime failures are normally represented using returned `error` values
- A successful operation conventionally returns `nil` as its error value
- The idiomatic error-handling pattern is `if err != nil { ... }`
- Errors can be handled immediately or propagated to a caller
- An error should normally be handled at the level where the program knows what to do with it
- `errors.New("message")` creates a simple error value
- `fmt.Errorf` can create errors containing formatted contextual information
- Error wrapping allows additional context to be added while preserving the original error
- `%w` in `fmt.Errorf` wraps an error
- `%v` formats an error as text but does not wrap it
- Wrapped errors form an error chain
- `errors.Is(err, target)` checks whether `err` or an error in its wrapping chain matches a target error
- Sentinel errors are predefined error values such as `var ErrNotFound = errors.New("not found")`
- `errors.Is` should generally be preferred over direct `==` comparison when errors may be wrapped
- `errors.As` searches an error chain for an error assignable to a particular error type
- `errors.As` is useful when the caller needs information stored inside a custom error type
- Custom error types implement the `error` interface by defining an `Error() string` method
- `errors.Unwrap` can retrieve an error wrapped directly inside another error
- In normal application code, callers usually use `errors.Is` and `errors.As` rather than manually unwrapping an entire error chain
- Errors should normally contain useful context as they propagate through application layers
- The same error should usually not be repeatedly logged at every layer while also being returned upward
- `defer` schedules a function call to execute when the surrounding function returns
- A deferred call executes whether the function returns normally or because of a panic
- `defer` is commonly used for cleanup such as closing files, closing response bodies, closing database rows, and unlocking mutexes
- Placing cleanup with `defer` close to resource acquisition makes code easier to reason about
- Multiple deferred calls execute in last-in-first-out order
- Deferred function arguments are evaluated when the `defer` statement executes
- The deferred function itself executes later, when the surrounding function returns
- A deferred closure can observe variables that are modified after the `defer` statement
- `panic` stops the normal execution of the current function
- When a panic occurs, Go begins unwinding the call stack
- During stack unwinding, deferred functions are executed
- If a panic reaches the top of a goroutine without being recovered, the program normally terminates
- `panic` is not intended to replace ordinary error handling
- Invalid user input, missing files, database failures, and network failures should normally be represented as errors
- `panic` is more appropriate for situations where normal execution cannot reasonably continue
- Some programmer errors and violated internal invariants may justify a panic
- `recover` can stop a panic from continuing to unwind the stack
- `recover` is intended to be called from a deferred function during panic unwinding
- Calling `recover` during normal execution returns `nil`
- After a panic is recovered, execution does not resume from the statement after the original `panic`
- Instead, the function in which the panic occurred finishes unwinding and control continues from the caller of the function where recovery occurred
- `recover` only handles a panic occurring in the same goroutine
- A panic in another goroutine cannot be recovered from a different goroutine
- Recovery is commonly used at system boundaries such as HTTP middleware, worker execution boundaries, or framework infrastructure
- Recovery code should usually log or record enough information to diagnose the panic
- `panic` and `recover` are not Go's equivalent of using exceptions for routine control flow
- The normal Go error-handling model is explicit error values returned through function results

## Materials

- [Lecture board: Group 1](./Lecture-6-G1.pdf)
- [Lecture board: Group 2](./Lecture-6-G2.pdf)

## References

- [A Tour of Go: Functions](https://go.dev/tour/basics/4)
- [A Tour of Go: Functions continued](https://go.dev/tour/basics/5)
- [A Tour of Go: Multiple results](https://go.dev/tour/basics/6)
- [A Tour of Go: Function values](https://go.dev/tour/moretypes/24)
- [A Tour of Go: Function closures](https://go.dev/tour/moretypes/25)
- [A Tour of Go: Errors](https://go.dev/tour/methods/19)
- [A Tour of Go: Defer](https://go.dev/tour/flowcontrol/12)
- [A Tour of Go: Stacking defers](https://go.dev/tour/flowcontrol/13)
- [Go Blog: Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
- [Go Blog: Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [Go Package Documentation: errors](https://pkg.go.dev/errors)
- [Go Package Documentation: fmt](https://pkg.go.dev/fmt)
- [Go Specification: Function declarations](https://go.dev/ref/spec#Function_declarations)
- [Go Specification: Function types](https://go.dev/ref/spec#Function_types)
- [Go Specification: Function literals](https://go.dev/ref/spec#Function_literals)
- [Go Specification: Defer statements](https://go.dev/ref/spec#Defer_statements)
- [Go Specification: Handling panics](https://go.dev/ref/spec#Handling_panics)
