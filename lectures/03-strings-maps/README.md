~~# Lecture 03 - Strings and Maps

## Learning Outcomes

- Work with strings, bytes, and runes in Go
- Explain how strings are represented as byte sequences
- Distinguish bytes from Unicode code points
- Iterate over UTF-8 encoded strings correctly
- Declare, initialize, read, update, and delete map entries
- Check whether a key exists using the comma-ok idiom
- Explain map semantics when assigning and passing maps to functions

## Key Ideas

- A Go string is an immutable sequence of bytes
- `byte` is an alias for `uint8`
- `rune` is an alias for `int32` and represents a Unicode code point
- UTF-8 represents Unicode code points using one or more bytes
- `len(string)` returns the number of bytes, not the number of runes
- Indexing a string with `s[i]` returns a byte
- `range` over a string decodes UTF-8 and produces runes
- The index returned by `range` over a string is a byte offset
- Converting a string to `[]byte` gives its raw bytes
- Converting a string to `[]rune` gives its Unicode code points
- A map stores key-value pairs and provides efficient lookup by key
- Map keys must be comparable
- Reading a missing map key returns the zero value of the value type
- The comma-ok idiom distinguishes a missing key from a stored zero value
- Map iteration order is unspecified
- Map assignment does not copy the underlying map data
- Maps behave reference-like, but Go still passes map values by value
- A nil map can be read from but cannot be written to

## Materials

- [Lecture board: Group 1](./Lecture-3-G1.pdf)
- [Lecture board: Group 2](./Lecture-3-G2.pdf)

## References

- [Go Blog: Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)
- [A Tour of Go: Range continued](https://go.dev/tour/moretypes/17)
- [unicode/utf8 package](https://pkg.go.dev/unicode/utf8)
- [strings package](https://pkg.go.dev/strings)
- [A Tour of Go: Maps](https://go.dev/tour/moretypes/19)
- [A Tour of Go: Map literals](https://go.dev/tour/moretypes/20)
- [A Tour of Go: Mutating Maps](https://go.dev/tour/moretypes/22)
- [Go by Example: Maps](https://gobyexample.com/maps)
- [Go Specification: Map types](https://go.dev/ref/spec#Map_types)~~
