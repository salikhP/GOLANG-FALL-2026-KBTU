# Lecture 09 - Synchronization and Goroutine Leaks

## Learning Outcomes

- Explain why synchronization is necessary in concurrent programs
- Understand the difference between synchronization for task completion and synchronization for shared data access
- Use `sync.WaitGroup` to wait for multiple goroutines to finish
- Understand how `Add`, `Done`, and `Wait` work together
- Explain why `WaitGroup` does not prevent data races
- Define a data race and identify situations where data races may occur
- Explain why operations such as `counter++` are not atomic
- Understand what a critical section is
- Distinguish between a data race and a race condition
- Use `sync.Mutex` to protect shared data
- Understand how `Lock` and `Unlock` work
- Use `defer` to safely release mutexes
- Understand the purpose of `sync.RWMutex`
- Distinguish between `Lock`/`Unlock` and `RLock`/`RUnlock`
- Choose appropriately between `Mutex` and `RWMutex`
- Use `sync.Once` to execute initialization code only once
- Explain what an atomic operation is
- Use the `sync/atomic` package for simple concurrent operations
- Understand the purpose of atomic `Load`, `Store`, and `Add` operations
- Explain the basic idea behind Compare-and-Swap (CAS)
- Understand why atomic operations do not automatically make multi-step logic safe
- Choose appropriately between mutexes and atomic operations
- Define a deadlock and explain how it occurs
- Recognize common deadlock situations involving mutexes and channels
- Understand how consistent lock ordering helps prevent deadlocks
- Use Go's built-in race detector to identify data races
- Run programs and tests with the `-race` flag
- Understand the limitations of race detection
- Define a goroutine leak and explain why it is dangerous
- Recognize how blocked channel operations can cause goroutine leaks
- Understand the importance of properly terminating goroutines
- Use `context.Context` to support cancellation of goroutines
- Understand how `context.WithCancel` and `ctx.Done()` work together
- Recognize common approaches to preventing goroutine leaks

## Key Ideas

### Synchronization and WaitGroup

- Concurrency allows multiple tasks to make progress independently, but their execution order is not guaranteed
- Synchronization coordinates concurrent operations to ensure correct behavior
- Synchronization can be used to wait for tasks to finish or to control access to shared resources
- A goroutine does not automatically make the function that launched it wait for completion
- When the `main` function returns, the program terminates even if other goroutines have not finished
- `sync.WaitGroup` allows a goroutine to wait for a group of tasks to complete
- A WaitGroup maintains an internal counter representing unfinished tasks
- `wg.Add(n)` increases the counter by `n`
- `wg.Done()` decreases the counter by one
- `wg.Wait()` blocks until the counter reaches zero
- The WaitGroup counter starts at zero
- When using the traditional `Add`/`Done` pattern, tasks should normally be registered before launching their goroutines
- `defer wg.Done()` is commonly used to mark a task as completed when its function returns
- A WaitGroup must not be copied after its first use
- When passed to a separate function, a WaitGroup is normally passed by pointer
- Starting with Go 1.25, `wg.Go(f)` can launch a goroutine and automatically track its completion
- `WaitGroup` does not guarantee the order in which goroutines execute
- `WaitGroup` does not protect shared variables from concurrent modification
- `WaitGroup` only ensures that registered tasks have finished before `Wait()` returns

### Data Races and Critical Sections

- Shared data is data that can be accessed by multiple goroutines
- A data race occurs when conflicting memory accesses are not properly synchronized
- A data race involves multiple goroutines accessing the same memory location, with at least one access being a write
- Concurrent reads of the same variable do not cause a data race if no conflicting writes occur
- Concurrent writes or combinations of reads and writes can cause data races
- `counter++` is not an atomic operation
- Conceptually, incrementing a counter involves reading its value, adding one, and writing the new value
- Multiple goroutines can read the same old value and overwrite each other's updates
- This problem is commonly called a lost update
- A program containing a data race may sometimes produce the expected output
- Correct output does not prove that concurrent code is free from data races
- A critical section is a part of the program that requires protection from conflicting concurrent access
- Critical sections often involve reading or modifying shared variables, maps, or other mutable data structures
- Data races and race conditions are related but different concepts
- A data race is a specific problem involving unsynchronized memory accesses
- A race condition is a broader logical problem where correctness depends on the timing or ordering of operations
- A program can contain a race condition even when all individual memory accesses are properly synchronized

### Mutex and RWMutex

- A mutex provides mutual exclusion, allowing only one goroutine to hold the lock at a time
- `sync.Mutex` is used to protect critical sections that access shared data
- `mu.Lock()` acquires a mutex or blocks until it becomes available
- `mu.Unlock()` releases the mutex
- When one goroutine holds a mutex, other goroutines trying to acquire the same mutex must wait
- A mutex should protect the entire operation that needs to be performed without conflicting concurrent access
- `defer mu.Unlock()` is commonly used immediately after acquiring a mutex
- Using `defer` helps ensure that the mutex is released when the surrounding function returns
- Holding a mutex longer than necessary can reduce concurrency and increase waiting time
- Go mutexes are not reentrant, so attempting to acquire the same mutex again while holding it can cause a deadlock
- A mutex must not be copied after its first use
- `sync.RWMutex` is a read-write mutex that distinguishes between read access and write access
- `RLock()` and `RUnlock()` are used for read-only operations
- Multiple goroutines can hold read locks simultaneously
- `Lock()` and `Unlock()` provide exclusive access for modifications
- A writer holding an `RWMutex` prevents other readers and writers from acquiring that lock
- `RWMutex` is useful for shared data that is read frequently and modified less often
- `RWMutex` is not automatically faster than `Mutex`
- The choice between `Mutex` and `RWMutex` depends on the workload and the simplicity of the implementation

### Other Synchronization Primitives

- The `sync` package provides additional synchronization primitives for different coordination problems
- `sync.Once` ensures that a function is executed only once, even when called concurrently
- `once.Do(f)` executes `f` during the first call and prevents subsequent executions
- Concurrent callers of `once.Do(f)` wait for the first execution to complete
- `sync.Once` is commonly used for one-time initialization of shared resources
- The `sync` package also provides more specialized tools such as `sync.Map`, `sync.Cond`, and `sync.Pool`

### Atomic Operations

- An atomic operation is performed as one indivisible operation relative to competing atomic operations
- Go provides atomic operations through the `sync/atomic` package
- Atomic operations can be useful for simple shared counters, flags, and state transitions
- `atomic.Int64` provides an integer type with atomic methods
- `Add(delta)` atomically adds a value to an atomic integer
- `Load()` atomically reads the current value
- `Store(value)` atomically replaces the current value
- `Swap(value)` atomically replaces a value and returns the previous value
- `CompareAndSwap(old, new)` changes a value only if its current value matches `old`
- Compare-and-Swap is commonly abbreviated as CAS
- CAS is useful for implementing conditional state transitions
- Atomic operations can avoid explicit mutex locking for some simple operations
- Atomic operations prevent data races only when all conflicting accesses use appropriate synchronization
- Making individual operations atomic does not automatically make a sequence of operations atomic
- A check followed by an update may still contain a logical race condition even when both operations are atomic
- Mutexes are generally easier to use when several variables or multiple operations must be protected together
- Atomic operations should be preferred only when they provide a clear benefit without making the code unnecessarily complicated

### Deadlocks

- A deadlock occurs when concurrent operations are blocked waiting for conditions that cannot be satisfied because of their dependencies
- Deadlocks can occur when goroutines acquire mutexes in conflicting orders
- Attempting to lock the same non-reentrant mutex twice without unlocking it can cause a deadlock
- Forgetting to unlock a mutex can cause other goroutines to block indefinitely
- A circular wait can occur when two goroutines each hold a lock needed by the other
- Consistent lock ordering is an important technique for preventing deadlocks involving multiple mutexes
- Keeping critical sections small and avoiding unnecessary nested locks can reduce deadlock risks
- Deadlocks can also occur when using channels
- Sending to an unbuffered channel blocks until a corresponding receiver is ready
- Receiving from an unbuffered channel blocks until a sender is ready, unless the channel is closed
- An operation waiting for a channel communication that can never happen may block forever
- The Go runtime can detect some situations where all goroutines are blocked and no progress is possible
- Not every deadlock is automatically detected by the Go runtime

### Race Detection

- Go provides a built-in race detector for identifying data races during program execution
- The race detector can be enabled using the `-race` flag
- `go run -race main.go` runs a program with race detection enabled
- `go test -race ./...` runs tests with race detection enabled
- `go build -race` builds a race-instrumented executable
- When a data race is detected, Go reports conflicting memory accesses and their stack traces
- The race detector only detects races that occur along executed code paths
- A successful run without race warnings does not guarantee that the program is completely race-free
- The race detector does not detect every logical race condition
- The race detector is not a general-purpose deadlock or goroutine leak detector
- Race detection introduces runtime and memory overhead
- Running tests with `-race` is a useful practice when developing concurrent Go applications

### Goroutine Leaks and Cancellation

- A goroutine leak occurs when a goroutine remains alive unnecessarily and cannot complete its intended work
- Goroutines consume memory and runtime resources even when they are blocked
- Leaked goroutines can accumulate over time in long-running applications
- A goroutine blocked on a channel send without a future receiver may never terminate
- A goroutine blocked on a channel receive without a future sender or closure may never terminate
- Infinite loops without an appropriate termination condition can also cause goroutine leaks
- Goroutines are not automatically terminated when the function that launched them returns
- Garbage collection does not automatically terminate blocked goroutines
- Every long-running goroutine should have a clearly defined lifecycle and termination condition
- The `context` package provides mechanisms for cancellation and deadlines
- `context.Context` can carry cancellation signals across function calls and goroutines
- `context.WithCancel(parent)` creates a derived context and returns a cancellation function
- Calling the cancellation function signals that operations using the context should stop
- `ctx.Done()` returns a channel that is closed when the context is canceled
- A goroutine can use `select` to wait for either a channel operation or a cancellation signal
- Context cancellation is cooperative and does not forcibly terminate goroutines
- Goroutines must check cancellation signals and return when appropriate
- `context.WithTimeout` can be used to establish a maximum duration for an operation
- Cancellation-aware channel operations help prevent goroutines from waiting indefinitely
- `runtime.NumGoroutine()` can be used to inspect the current number of goroutines
- A continuously increasing goroutine count under a stable workload can indicate a possible leak
- Not every blocked goroutine is leaked; some goroutines intentionally wait for work

## Further Study (Optional)

The following topics are useful for deeper understanding but are not required to master this lecture:

- The formal Go memory model and the happens-before relationship
- Memory visibility guarantees provided by synchronization primitives
- Performance differences between `Mutex`, `RWMutex`, and atomic operations
- Implementing retry loops with Compare-and-Swap
- Lock contention and its effects on concurrent performance
- Advanced synchronization primitives such as `sync.Cond` and `sync.Map`
- Profiling mutex contention and blocked goroutines
- Detecting goroutine leaks using runtime profiling tools

Students should prioritize understanding data races, WaitGroups, mutexes, atomic operations, deadlocks, race detection, and goroutine cancellation.

## Materials

- [Lecture board: Group 1](./Lecture-10-G1.pdf)
- [Lecture board: Group 2](./Lecture-10-G2.pdf)

## References

- [Go Package Documentation: sync](https://pkg.go.dev/sync)
- [Go Package Documentation: sync/atomic](https://pkg.go.dev/sync/atomic)
- [Go Package Documentation: context](https://pkg.go.dev/context)
- [Go Package Documentation: runtime](https://pkg.go.dev/runtime)
- [Go Package Documentation: runtime/pprof](https://pkg.go.dev/runtime/pprof)
- [Go Memory Model](https://go.dev/ref/mem)
- [Go Documentation: Data Race Detector](https://go.dev/doc/articles/race_detector)
- [A Tour of Go: Goroutines](https://go.dev/tour/concurrency/1)
- [A Tour of Go: Mutex](https://go.dev/tour/concurrency/9)
- [Go Blog: Share Memory By Communicating](https://go.dev/blog/codelab-share)
- [Go Blog: Go Concurrency Patterns — Pipelines and Cancellation](https://go.dev/blog/pipelines)
- [Go Blog: Go Concurrency Patterns — Context](https://go.dev/blog/context)
- [Go Blog: Goroutine Leak Profiles](https://go.dev/blog/goroutine-leak-profiles)
