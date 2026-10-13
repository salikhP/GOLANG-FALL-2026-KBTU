# Lecture 07 - Memory Management

## Learning Outcomes

- Explain how memory management works in Go
- Distinguish between stack and heap memory
- Understand how goroutine stacks and function stack frames work
- Explain how Go determines where values are stored
- Understand the purpose of escape analysis
- Identify common situations that cause values to escape to the heap
- Explain why using pointers does not necessarily cause heap allocation
- Understand memory allocations and their performance implications
- Explain how slices allocate and grow their backing arrays
- Reduce unnecessary allocations through preallocation
- Use compiler flags to inspect escape analysis decisions
- Measure memory allocations using Go benchmarks
- Explain the purpose of garbage collection
- Understand how Go's mark-and-sweep garbage collector works
- Explain reachability and how garbage collection identifies unused objects
- Understand the relationship between allocations and garbage collection overhead
- Explain the purpose of `GOGC` and `GOMEMLIMIT`
- Recognize how memory leaks can occur despite automatic garbage collection
- Use basic Go tools to investigate memory usage and performance

## Study Expectations

You are **not expected to memorize or fully understand every concept** listed in this lecture.

Focus primarily on understanding stack vs. heap memory, escape analysis, memory allocations, and the basic principles of garbage collection.

More advanced topics, such as `GOGC`, `GOMEMLIMIT`, benchmarking, memory profiling, and runtime diagnostics, are included for **optional self-study**. They are useful for understanding how memory management works in real-world Go applications, but are not required at this stage.

The main goal is to understand **how Go manages memory and why it matters**, rather than memorizing implementation details or optimization techniques.

## Key Ideas

- Memory management involves allocating, using, and reclaiming memory during program execution
- Go provides automatic memory management through its compiler and runtime
- Go developers do not normally need to manually allocate and free memory
- The compiler determines which values can use stack storage and which require heap storage
- The stack is primarily used for temporary storage associated with function execution
- Each goroutine has its own stack, which can grow and shrink dynamically
- Function calls use stack frames to manage local data and execution state
- Stack storage can be reused when function frames are no longer needed
- Stack allocation is generally faster and does not require garbage collection of individual objects
- The heap stores objects that require dynamically managed lifetimes
- Heap objects can remain alive after the function that created them returns
- Heap memory is managed by the Go runtime and garbage collector
- A local variable is not guaranteed to be allocated on the stack
- Escape analysis is a compile-time process that determines whether values can safely remain on the stack
- A value may escape when its address must remain valid beyond the lifetime of its function's stack frame
- Returning a pointer to a local variable commonly causes that variable to escape to the heap
- Storing a pointer in a global variable can cause the referenced value to escape
- Passing a pointer to another function does not automatically cause heap allocation
- Sharing values with goroutines or capturing them in closures may cause heap allocation
- Large objects or dynamically sized data may require heap allocation
- The Go compiler can optimize values into CPU registers or eliminate unnecessary storage
- The command `go build -gcflags="-m=2"` displays compiler optimization and escape analysis information
- The compiler may report messages such as `moved to heap` and `does not escape`
- A memory allocation reserves storage for a value or data structure
- `new(T)` creates a zero-initialized value and returns a pointer to it
- `make()` initializes slices, maps, and channels
- Neither `new()` nor `make()` automatically guarantees heap allocation
- A slice consists of a descriptor containing a pointer, length, and capacity
- A slice's backing array stores its elements and may be allocated separately from its descriptor
- Appending beyond a slice's capacity causes a new backing array to be allocated and existing elements to be copied
- Preallocating slice capacity with `make([]T, 0, capacity)` can reduce allocations and copying
- Heap allocations introduce runtime overhead and may increase garbage collection work
- Reducing unnecessary allocations can improve application performance
- Not every allocation is a performance problem, and optimization should be based on measurements
- Go provides built-in benchmarks through the `testing` package
- The command `go test -bench=. -benchmem` measures benchmark performance and memory allocations
- `ns/op` represents the average execution time per operation
- `B/op` represents the average number of heap-allocated bytes per operation
- `allocs/op` represents the average number of heap allocations per operation
- Garbage collection automatically reclaims memory occupied by unreachable heap objects
- Go uses a concurrent tracing garbage collector based on mark-and-sweep
- Garbage collection begins by identifying reachable objects from GC roots
- GC roots include relevant global references and live references from goroutine stacks
- During the mark phase, the collector traces references and marks reachable heap objects
- During the sweep phase, memory belonging to unmarked objects becomes available for reuse
- Go performs much of its garbage collection concurrently with application execution
- Garbage collection may still require short stop-the-world pauses
- An object becomes eligible for collection when it is no longer reachable
- Unreachable objects are not necessarily reclaimed immediately
- The garbage collector does not automatically release objects that remain reachable through unnecessary references
- Growing caches, global collections, and retained references can cause logical memory leaks
- `GOGC` controls the trade-off between garbage collection frequency and memory consumption
- The default `GOGC` value is `100`
- Increasing `GOGC` generally reduces garbage collection frequency but increases memory usage
- `GOMEMLIMIT` configures a soft memory limit for the Go runtime
- The `GOMEMLIMIT` setting is not a strict limit on total operating-system process memory
- The environment variable `GODEBUG=gctrace=1` enables garbage collection diagnostics
- The `runtime.MemStats` structure provides information about heap allocation and GC activity
- Heap profiling with `pprof` helps identify allocation hotspots and retained memory
- Automatic garbage collection does not replace explicitly closing files, connections, and other external resources
- Memory optimization should focus on reducing unnecessary allocations and object retention without sacrificing code readability

## Materials

- Lecture board: Group 1 (to be uploaded)
- Lecture board: Group 2 (to be uploaded)

## References

- [Go Documentation: A Guide to the Go Garbage Collector](https://go.dev/doc/gc-guide)
- [Go Blog: Allocating on the Stack](https://go.dev/blog/allocation-optimizations)
- [Go Blog: The Green Tea Garbage Collector](https://go.dev/blog/greenteagc)
- [Go Blog: Getting to Go - The Journey of Go's Garbage Collector](https://go.dev/blog/ismmkeynote)
- [Go Blog: Profiling Go Programs](https://go.dev/blog/pprof)
- [Go Documentation: Diagnostics](https://go.dev/doc/diagnostics)
- [Go Documentation: Command go](https://pkg.go.dev/cmd/go)
- [Go Standard Library: runtime](https://pkg.go.dev/runtime)
- [Go Standard Library: runtime.MemStats](https://pkg.go.dev/runtime#MemStats)
- [Go Standard Library: runtime/pprof](https://pkg.go.dev/runtime/pprof)
- [Go Standard Library: testing - Benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)
