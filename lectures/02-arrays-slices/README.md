# Lecture 02 - Arrays and Slices

## Learning Outcomes

- Declare and initialize arrays and slices in Go

- Use len, cap, slicing, append, and copy

- Distinguish arrays from slices

- Explain how slices behave when passed to functions

- Predict the effects of modifying shared slices

- Use slices in common backend scenarios

## Key Ideas

- Arrays store elements contiguously in memory
- Array length is part of the array type
- Array indexing is `O(1)`
- Arrays use value semantics
- A slice is a view over a backing array
- A slice conceptually contains a pointer, length, and capacity
- Multiple slices can share the same backing array
- Slicing usually does not copy data
- `append` may reuse or replace the backing array
- Slice assignment does not copy the underlying elements
- `copy` creates an independent copy when used with separate storage
- Nil and empty slices can behave similarly internally but differ semantically in some contexts

## Materials

- [Lecture board](./lecture.svg)
- [Excalidraw source](./lecture.excalidraw)
- [Code examples](./examples/)

## References

- [Go by Example: Arrays](https://gobyexample.com/arrays)
- [A Tour of Go: Arrays](https://go.dev/tour/moretypes/6)
- [Go by Example: Slices](https://gobyexample.com/slices)
- [A Tour of Go: Arrays](https://go.dev/tour/moretypes/7)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Flow control statements: for, if, else, switch and defer](https://go.dev/tour/flowcontrol/1)
