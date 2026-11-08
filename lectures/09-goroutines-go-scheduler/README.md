# Lecture 09 - Goroutines and the Go Scheduler

## Learning Outcomes

- Explain what CPU cores, processes, and operating system threads are
- Understand how processes and threads relate to each other
- Explain how multiple threads share CPU execution time
- Understand what a context switch is
- Distinguish between concurrency and parallelism
- Explain how concurrency can exist without parallelism
- Distinguish between CPU-bound and I/O-bound tasks
- Understand why concurrency is useful for I/O-bound applications
- Explain what goroutines are and why Go uses them
- Understand the differences between goroutines and OS threads
- Create goroutines using the `go` keyword
- Understand that goroutines execute independently of their callers
- Explain why the execution order of goroutines is not guaranteed
- Understand what happens to other goroutines when `main()` returns
- Use `sync.WaitGroup` to wait for multiple goroutines to finish
- Understand the role of the Go runtime scheduler
- Explain the basic responsibilities of G, M, and P in Go's scheduler
- Understand how goroutines are scheduled onto OS threads
- Explain the purpose of `GOMAXPROCS`
- Understand how the scheduler handles goroutines waiting for I/O
- Understand the basic lifecycle of a goroutine
- Distinguish between runnable, running, waiting, and completed goroutines
- Explain what a goroutine leak is
- Identify common situations that cause goroutine leaks
- Understand why goroutine leaks can consume memory and other resources
- Understand how cancellation and timeouts help prevent goroutine leaks
- Recognize the role of `context.Context` in stopping background work
- Use `runtime.NumGoroutine()` to inspect the number of goroutines
- Understand why creating more goroutines does not always improve performance

## Key Ideas

- A CPU executes machine instructions
- Modern CPUs can contain multiple physical cores and support multiple logical CPUs
- A process is a running instance of a program
- A process has its own virtual address space and operating system resources
- A thread is an independently schedulable sequence of instructions within a process
- A process can contain multiple OS threads
- Threads within the same process share memory and other process resources
- Each OS thread has its own execution state and stack
- The operating system schedules threads onto available logical CPUs
- A CPU can switch between multiple runnable threads, allowing them to make progress over time
- A context switch involves saving and restoring execution state when switching between threads
- Context switching introduces overhead
- Creating large numbers of OS threads can consume significant system resources
- Concurrency means structuring a program so that multiple tasks can make progress during overlapping periods
- Parallelism means executing multiple computations simultaneously
- Concurrency does not necessarily require multiple CPU cores
- Parallelism requires multiple execution resources
- Concurrent programs are not automatically parallel
- CPU-bound tasks spend most of their execution time performing computations
- I/O-bound tasks spend much of their time waiting for operations such as network requests, database queries, or file access
- Concurrency can improve resource utilization by allowing other tasks to execute while one task waits for I/O
- Concurrency does not guarantee better performance in every situation
- A goroutine is a function executing concurrently with other goroutines in the same Go program
- Goroutines are lightweight units of concurrent execution managed by the Go runtime
- Goroutines are not the same as operating system threads
- Goroutines are multiplexed onto OS threads by the Go scheduler
- Many goroutines can share a smaller number of OS threads
- Goroutines start with small stacks that grow and shrink as needed
- Goroutines are generally cheaper to create than OS threads
- The `go` keyword starts a function call in a new goroutine
- For example, `go sendEmail()` starts `sendEmail` concurrently with the calling goroutine
- The calling goroutine does not wait for a newly created goroutine to finish
- Calling a function without `go` executes it in the current goroutine
- The order in which goroutines execute is not guaranteed
- A newly created goroutine is not guaranteed to execute immediately
- The `main()` function executes in the main goroutine
- When `main()` returns, the program terminates, even if other goroutines have not completed
- Using `time.Sleep()` to wait for goroutines is unreliable because it does not provide synchronization
- `sync.WaitGroup` allows a goroutine to wait for a collection of tasks to finish
- `wg.Add(n)` increases the number of tasks being tracked
- `wg.Done()` decreases the task counter by one
- `wg.Wait()` blocks until the task counter reaches zero
- `defer wg.Done()` is a common pattern for marking a goroutine's work as completed
- `WaitGroup` coordinates task completion but does not prevent concurrent access to shared data
- The Go scheduler determines which runnable goroutines execute on available OS threads
- Go scheduling and OS thread scheduling operate at different levels
- The operating system schedules OS threads, while the Go runtime schedules goroutines
- Go's runtime scheduler is commonly explained using the G-M-P model
- G stands for Goroutine and represents a unit of work
- M stands for Machine and represents an operating system thread
- P stands for Processor and represents runtime resources required to execute user Go code
- P is a logical runtime resource, not a physical CPU core
- An OS thread generally needs a P to execute user Go code
- The number of Ps is controlled by `GOMAXPROCS`
- `GOMAXPROCS` limits how many OS threads can execute user Go code simultaneously
- `GOMAXPROCS` does not limit the total number of goroutines
- `GOMAXPROCS` does not directly limit the total number of OS threads in the process
- Setting `GOMAXPROCS(1)` allows concurrency but prevents simultaneous execution of user Go code on multiple OS threads
- The default `GOMAXPROCS` value generally reflects available logical CPUs and may account for container CPU limits
- When a goroutine waits for supported network I/O, the Go runtime can park it and schedule another runnable goroutine
- Go uses runtime mechanisms such as a network poller to efficiently manage many waiting network operations
- When an OS thread blocks in a system call, the runtime can allow another thread to execute Go code using an available P
- The scheduler maintains queues of runnable goroutines
- Work stealing allows an idle P to obtain runnable work from another P
- Work stealing helps distribute runnable goroutines across available execution resources
- Preemption allows the runtime to interrupt a running goroutine so other goroutines can make progress
- `runtime.Gosched()` voluntarily yields execution and gives other runnable goroutines an opportunity to execute
- Goroutine scheduling is nondeterministic, so programs should not depend on a particular execution order
- A goroutine can be runnable, running, waiting, or finished
- A runnable goroutine is ready to execute but is not currently running
- A running goroutine is actively executing instructions on an OS thread
- A waiting goroutine is blocked until an operation or event allows it to continue
- Goroutines can wait for operations such as network I/O, timers, locks, and channel communication
- When a waiting goroutine becomes unblocked, it becomes eligible for scheduling again
- A goroutine finishes when its function returns
- A goroutine leak occurs when a goroutine remains alive unnecessarily and cannot finish as intended
- A goroutine waiting forever for a channel operation is a common example of a leak
- Goroutine leaks can also result from infinite loops without shutdown conditions, blocking operations, or missing cancellation
- Blocked goroutines still consume memory and runtime resources
- Goroutines may retain references to memory that cannot be garbage-collected while they remain alive
- Goroutines are not automatically terminated by the garbage collector simply because they are blocked
- Long-running applications can accumulate leaked goroutines over time
- Every goroutine should have a clear purpose and an appropriate termination condition
- Some goroutines intentionally run for the lifetime of the application and are not considered leaks
- `context.Context` can carry cancellation signals and deadlines across operations
- `context.WithCancel` creates a context that can be explicitly canceled
- `ctx.Done()` provides a channel that becomes ready when the context is canceled
- Canceling a context does not forcibly terminate a goroutine
- Goroutines must cooperate with cancellation by checking the context or using operations that support it
- Timeouts can prevent operations from waiting indefinitely
- `runtime.NumGoroutine()` returns the current number of goroutines
- Unexpectedly increasing goroutine counts can indicate potential leaks
- A growing goroutine count does not automatically mean a leak exists
- Goroutine profiles and execution traces can help investigate blocked or long-running goroutines
- Goroutines are lightweight, but creating an unlimited number of them can exhaust application resources
- A large number of goroutines does not mean that all of them execute simultaneously
- More goroutines do not necessarily make CPU-bound workloads faster
- Good concurrent programs manage task lifetimes, resource usage, and synchronization explicitly

## Important Notes

This lecture focuses on understanding how concurrency works in Go rather than memorizing every detail of the runtime implementation.

Students should be comfortable with:

- Explaining concurrency vs. parallelism
- Understanding the relationship between CPU cores, OS threads, and goroutines
- Creating goroutines using the `go` keyword
- Understanding why goroutine execution order is not guaranteed
- Waiting for goroutines using `sync.WaitGroup`
- Explaining the basic purpose of the Go scheduler
- Recognizing common goroutine leak scenarios
- Understanding why cancellation and proper cleanup matter

The following topics are provided for deeper understanding and are **not required to be memorized in detail**:

- Internal implementation of the G-M-P scheduler
- Local and global run queues
- Work stealing algorithms
- Goroutine preemption mechanisms
- Network poller internals
- Scheduler tracing and profiling tools
- Detailed `GOMAXPROCS` configuration and tuning

It is more important to understand **why Go uses a scheduler and how goroutines behave** than to memorize how every scheduling decision is implemented.

Detailed synchronization mechanisms, such as mutexes, channels, and race-condition prevention, will be covered separately.

## Materials

- [Lecture board: Group 1](./Lecture-9-G1.pdf)
- [Lecture board: Group 2](./Lecture-9-G2.pdf)

## References

- [A Tour of Go: Goroutines](https://go.dev/tour/concurrency/1)
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Effective Go: Goroutines](https://go.dev/doc/effective_go#goroutines)
- [Go Blog: Concurrency is not Parallelism](https://go.dev/blog/waza-talk)
- [Rob Pike: Concurrency is not Parallelism (Slides)](https://go.dev/talks/2012/waza.slide)
- [Go Runtime: Scheduler Structures (G, M, P)](https://go.dev/src/runtime/HACKING)
- [Go Blog: Container-aware GOMAXPROCS](https://go.dev/blog/container-aware-gomaxprocs)
- [Go Blog: Go Concurrency Patterns — Pipelines and Cancellation](https://go.dev/blog/pipelines)
- [Go Blog: Go Concurrency Patterns — Context](https://go.dev/blog/context)
- [Go Package Documentation: runtime](https://pkg.go.dev/runtime)
- [Go Package Documentation: sync](https://pkg.go.dev/sync)
- [Go Package Documentation: context](https://pkg.go.dev/context)
- [Go Diagnostics: Profiling and Tracing](https://go.dev/doc/diagnostics)
- [Go Specification: Go statements](https://go.dev/ref/spec#Go_statements)
- [Go 1.14 Release Notes: Goroutine Preemption](https://go.dev/doc/go1.14)
