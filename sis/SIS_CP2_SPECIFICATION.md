# SIS Project - Checkpoint 2 Specification

**Course:** Golang Application Development  
**Checkpoint:** 2 - Concurrent Processing  
**Points:** 5  
**Project type:** Individual cumulative project

## Project Overview

Checkpoint 2 extends the application developed in Checkpoint 1 by introducing **concurrent task processing**.

In Checkpoint 1, tasks were created, stored, and processed synchronously.

In Checkpoint 2, your application must be able to process multiple independent tasks concurrently using a **worker pool**.

Instead of processing each task immediately in the calling goroutine, tasks should be submitted to a bounded queue and processed by worker goroutines.

For example:

```text
                 Task Queue
                 (bounded)
                     |
                     v
               +-----------+
Task 1 ------> |           |
Task 2 ------> |   Queue   |
Task 3 ------> |           |
               +-----------+
                     |
          +----------+----------+
          |          |          |
          v          v          v
       Worker 1   Worker 2   Worker 3
          |          |          |
          v          v          v
       Process    Process    Process
          |          |          |
          v          v          v
       Results    Results    Results
```

Workers must process tasks independently while safely updating the shared application state.

The application must preserve all functionality from Checkpoint 1, including:
- task creation and validation;
- unique task identifiers;
- in-memory task storage;
- task retrieval;
- meaningful domain-specific processing;
- task state management;
- error handling;
- automated unit tests.

**You must continue developing your existing Checkpoint 1 project.**

Do not replace it with an unrelated application.

The original domain and processing logic should remain recognizable. You may refactor the implementation to support concurrency.

## 1. Goal

Checkpoint 2 focuses on implementing correct concurrent task processing in Go.

You will extend your application to:

- process tasks concurrently using multiple goroutines;
- implement a worker pool;
- use a bounded queue for pending tasks;
- communicate between goroutines using channels;
- synchronize access to shared application state;
- correctly update task states during concurrent processing;
- prevent data races and invalid concurrent operations;
- manage worker startup and termination;
- test concurrent behavior using automated tests;
- verify concurrency safety using Go's race detector.

The focus of this checkpoint is on the concurrency concepts covered in Weeks 9-12:

- goroutines and their lifecycle;
- channels and communication;
- buffered channels;
- `sync.Mutex` and `sync.RWMutex`;
- `sync.WaitGroup`;
- race conditions and critical sections;
- worker pools;
- synchronization and deadlock prevention.

You are expected to demonstrate that you understand how concurrent components interact and why your implementation is safe.

HTTP APIs, databases, and external infrastructure are not required.

Context-based cancellation, HTTP integration, and application-level graceful shutdown will be addressed in Checkpoint 3.

## 2. Freedom of Implementation

As in Checkpoint 1, there is **no required project template, package structure, function signature, or architecture**.

You may choose:

- how to organize your packages;
- how to represent workers;
- how to implement the task queue;
- how many workers to create;
- how to configure the queue capacity;
- which synchronization primitives to use;
- how to expose task submission and retrieval operations;
- how to start and stop concurrent processing.

However, your design must satisfy the observable requirements of this specification.

You are responsible for explaining your concurrency design during the individual code defense.

**Using goroutines alone does not automatically satisfy the requirements.**

Your application must demonstrate meaningful concurrent task processing with controlled resource usage and correct synchronization.

## 3. Required Functionality

Your application must support the following behavior:

1. Preserve the domain-specific task processing implemented in Checkpoint 1.
2. Create and validate tasks as before.
3. Store tasks and their results in memory.
4. Submit existing or newly created tasks for asynchronous processing.
5. Maintain a bounded queue of submitted tasks waiting for processing.
6. Start multiple worker goroutines that consume tasks from the queue.
7. Allow different workers to process independent tasks concurrently.
8. Update the state of each task according to its processing lifecycle.
9. Store processing results in the correct corresponding tasks.
10. Allow tasks to be retrieved while other tasks are being processed.
11. Safely handle concurrent access to shared task data.
12. Prevent the same task from being processed simultaneously by multiple workers.
13. Handle attempts to submit tasks in an invalid state.
14. Handle task submission when the queue is full.
15. Support controlled termination of worker goroutines.
16. Prevent new tasks from being accepted after processing has been stopped.
17. Ensure that accepted tasks and their results are not silently lost during normal worker shutdown.

Task submission must not require the caller to wait for the task's processing to finish.

For example:

```text
Create Task A
    |
    v
Submit Task A
    |
    v
Task A enters queue
    |
    v
Submit returns
    |
    v
Worker eventually processes Task A
    |
    v
Task A becomes Completed
```

The exact API and implementation are your choice.

## 4. Worker Pool

Your application must implement a worker pool for processing tasks.

A worker pool consists of a fixed number of worker goroutines that receive tasks from a shared queue.

### Requirements

1. The worker pool must contain at least **2 worker goroutines**.
2. Workers must be able to process different tasks concurrently.
3. Workers must receive tasks through a channel-based queue.
4. Workers should remain available to process additional tasks rather than terminate after processing a single task.
5. The number of worker goroutines must be controlled.
6. Workers must be able to terminate when the application stops processing.
7. Worker termination must be coordinated correctly.

For example:

```text
Workers: 2

Queue:
Task 1, Task 2, Task 3, Task 4

Worker 1 -> Task 1
Worker 2 -> Task 2

After Worker 1 finishes:
Worker 1 -> Task 3

After Worker 2 finishes:
Worker 2 -> Task 4
```

The processing order is not guaranteed.

For example, Task 2 may finish before Task 1 even if Task 1 was submitted first.

Your application must not depend on the order in which concurrent tasks complete.

### Important

Creating a separate goroutine for every submitted task without controlling concurrency is **not considered a worker pool**.

For example:

```go
for _, task := range tasks {
    go process(task)
}
```

This approach alone does not satisfy the worker pool requirement.

Your implementation must reuse a controlled number of workers.

## 5. Bounded Task Queue

Your application must use a **bounded queue** to store tasks waiting for processing.

A bounded queue has a finite capacity.

For example:

```text
Queue capacity: 3

Queue:
[Task A] [Task B] [Task C]

The queue is full.
```

When workers take tasks from the queue, space becomes available for new submissions.

### Requirements

1. The queue must have a clearly defined finite capacity.
2. Tasks must enter the queue before being processed by workers.
3. Workers must consume tasks from the queue.
4. The application must correctly handle a full queue.
5. Queue capacity must not grow indefinitely as more tasks are submitted.
6. Task submission must not create an unbounded number of goroutines to bypass the queue limit.

You may choose how to handle a full queue.

For example:

**Option A - Reject submission**

```text
Queue is full

Submit Task D
    |
    v
Return an error
```

**Option B - Block submission**

```text
Queue is full

Submit Task D
    |
    v
Wait until a worker removes a task
    |
    v
Submit Task D
```

Both approaches are acceptable.

However, the chosen behavior must be clearly documented and correctly implemented.

If submission blocks, your implementation must avoid leaving callers blocked indefinitely because the worker pool has stopped.

If submission is rejected, the task must remain in a valid state and must not be incorrectly marked as processing.

## 6. Task Lifecycle

Your application must preserve the task lifecycle introduced in Checkpoint 1 and adapt it to concurrent processing.

For example:

```text
Pending
   |
   v
Queued
   |
   v
Processing
   |
   v
Completed
```

A task may have the following states:

- **Pending:** The task has been created but has not been submitted.
- **Queued:** The task has been accepted for processing and is waiting in the queue.
- **Processing:** A worker is currently processing the task.
- **Completed:** Processing has finished successfully.

You may use different state names.

A separate `Queued` state is recommended but not mandatory. For example, you may keep a task in `Pending` while it waits in the queue, provided that your application still correctly tracks submissions and prevents duplicate processing.

### Requirements

1. Task state must reflect the actual processing lifecycle.
2. State changes must be synchronized when accessed by multiple goroutines.
3. Multiple workers must not process the same task simultaneously.
4. Completed tasks must retain their correct processing results.
5. Invalid operations and state transitions must be handled correctly.

For example:

```text
Task A: Processing

Submit Task A again

Result: Error
```

Similarly:

```text
Task B: Completed

Submit Task B again

Result: Error
```

You may support retries or reprocessing if they are meaningful for your domain, but their behavior must be explicitly defined and correctly synchronized.

If processing itself can fail, the application must not incorrectly mark the task as successfully completed. You may introduce a `Failed` state or another appropriate representation.

Briefly document your task lifecycle and allowed transitions in the repository README.

## 7. Synchronization and Shared State

Multiple goroutines may access the same application data concurrently.

For example:

```text
Worker 1 -> Updates Task A

Worker 2 -> Updates Task B

Main goroutine -> Retrieves Task A

Another goroutine -> Retrieves Task B
```

Your application must ensure that shared data remains consistent.

### Requirements

1. Shared mutable data must be accessed safely.
2. Concurrent map reads and writes must be synchronized.
3. Task state updates must not introduce data races.
4. Task results must not be overwritten by unrelated workers.
5. Concurrent task retrieval must behave correctly.
6. Operations that must be atomic must not be implemented as unsafe sequences of independent reads and writes.
7. Synchronization must not introduce deadlocks.

Use appropriate synchronization mechanisms.

For example:

- `sync.Mutex`;
- `sync.RWMutex`;
- channels and ownership-based communication;
- other appropriate standard Go synchronization primitives.

You are not required to use every synchronization primitive covered in the course.

Choose mechanisms that make sense for your implementation.

### Important

Do not assume that accessing different keys in the same ordinary Go map concurrently is automatically safe.

Also, protecting a map does not necessarily protect the mutable fields of objects stored inside it.

You must understand which data is shared, which operations require synchronization, and what each synchronization mechanism protects.

Avoid holding locks during long-running task processing unless there is a justified reason.

## 8. Worker Lifecycle and Shutdown

Your application must correctly manage the lifecycle of worker goroutines.

At minimum, it must support:

1. Starting workers.
2. Allowing workers to receive and process tasks.
3. Stopping acceptance of new tasks.
4. Finishing tasks already accepted for processing.
5. Waiting for worker goroutines to terminate.

For example:

```text
Start worker pool
       |
       v
Workers wait for tasks
       |
       v
Tasks are submitted
       |
       v
Workers process tasks
       |
       v
Stop requested
       |
       v
Stop accepting new tasks
       |
       v
Finish accepted tasks
       |
       v
Workers terminate
       |
       v
Shutdown completes
```

You may use `sync.WaitGroup`, channel closure, or other appropriate mechanisms to coordinate worker termination.

### Requirements

- Workers must not remain running indefinitely after a normal shutdown.
- The application must wait for worker goroutines to terminate.
- Tasks accepted before shutdown must not be silently discarded.
- Task submission and shutdown must be safe even when they occur concurrently.
- A shutdown operation must not cause a panic from sending to a closed channel.
- Calling shutdown more than once must be handled safely, either as a no-op or through an appropriate error.

You must understand which part of your application owns each channel and which component is responsible for closing it.

**Context cancellation, deadlines, and forced interruption of running tasks are not required for Checkpoint 2.**

The goal here is basic worker lifecycle management, not the complete application-level graceful shutdown required in Checkpoint 3.

## 9. Automated Tests

You must extend your existing automated tests using Go's standard `testing` package.

All tests from Checkpoint 1 must continue to pass.

Your new tests must cover, at minimum:

1. Successful worker pool startup.
2. Successful submission of tasks for asynchronous processing.
3. Successful processing of multiple tasks.
4. Evidence that at least two independent tasks can be processed concurrently.
5. Correct processing results for individual tasks.
6. Correct task states after processing.
7. Correct handling of invalid or duplicate task submissions.
8. Correct handling of a full queue.
9. Safe concurrent access to shared task data.
10. Correct worker termination after shutdown.
11. Rejection of new submissions after shutdown.

Your tests must verify expected behavior, not merely execute concurrent code.

For example, a test that simply launches several goroutines and finishes without checking the results is insufficient.

### Testing Concurrent Behavior

Concurrent execution can be nondeterministic.

Tests should not depend on a particular order in which workers finish their tasks.

For example, this assumption is incorrect:

```text
Task A was submitted first.

Therefore, Task A must finish first.
```

Instead, verify that:

- all accepted tasks are processed;
- every task receives the correct result;
- task states are correct;
- the worker pool operates within the specified concurrency limits;
- workers terminate correctly.

Prefer synchronization mechanisms such as channels and `sync.WaitGroup` to coordinate concurrency tests.

Avoid relying only on arbitrary `time.Sleep` calls to make tests pass.

At least one test should demonstrate actual overlapping task processing, rather than merely proving that multiple tasks can be processed sequentially.

### Race Detection

Your application must pass:

```bash
go test ./...
go test -race ./...
```

The race detector must not report data races when your concurrency tests execute.

**Important:** Passing the race detector does not automatically prove that your implementation is free from all concurrency problems.

For example, a program may be free from data races but still contain:
- deadlocks;
- incorrect state transitions;
- lost tasks;
- goroutine leaks;
- logical concurrency errors.

Your tests should also check these behaviors where applicable.

There is no mandatory code coverage percentage for Checkpoint 2.

## 10. Technical Requirements

All technical requirements from Checkpoint 1 remain applicable unless explicitly modified in this specification.

Your repository must:

- continue using the Go module from Checkpoint 1;
- preserve the existing domain-specific processing functionality;
- implement concurrent task processing with a worker pool;
- use a bounded channel-based task queue;
- use appropriate synchronization for shared state;
- correctly manage worker startup and termination;
- handle expected application errors through normal Go error handling;
- compile successfully;
- be formatted with `gofmt`;
- pass `go vet ./...`;
- pass `go test ./...`;
- pass `go test -race ./...`.

Do not use `panic` for expected application errors.

Do not rely on arbitrary delays or sleeps to make the application appear concurrent.

You may use standard library packages such as:

```go
sync
sync/atomic
time
errors
fmt
```

You are free to use other appropriate standard library packages.

External concurrency libraries are not necessary.

### Not Required

For Checkpoint 2, you do not need:

- an HTTP API;
- a database;
- message brokers;
- external task queues;
- cloud infrastructure;
- context-based cancellation;
- HTTP request timeouts;
- application-level graceful HTTP server shutdown;
- benchmarks or profiling.

These topics are outside the required scope of this checkpoint.

Do not add unnecessary infrastructure merely to make the project appear more complex.

## 11. README Documentation

Update your repository README to describe the concurrent version of your application.

At minimum, explain:

1. The purpose of your application and its task-processing logic.
2. How tasks are created and submitted.
3. How the worker pool operates.
4. How many workers are used.
5. The capacity of the task queue.
6. What happens when the queue is full.
7. How task states change during processing.
8. How shared data is synchronized.
9. How workers are stopped.
10. How to run the application and automated tests.

A short explanation of your concurrency architecture is sufficient.

You may include a simple diagram illustrating the relationship between the task queue, workers, and task storage.

You are not required to produce a formal technical report.

## 12. Submission and Individual Code Defense

Continue using your individual GitHub repository from Checkpoint 1.

Submit the repository link and the corresponding commit link or tag through Microsoft Teams before the announced deadline.

The submitted version is the version evaluated during the defense. Changes made after the deadline may not be considered.

Each checkpoint includes a short individual code defense.

The defense normally lasts approximately **2 minutes**.

There is no presentation and no slides.

During the defense, you may be asked to:

- explain how your worker pool works;
- explain how tasks move through the queue;
- justify your use of channels and synchronization primitives;
- identify which data is shared between goroutines;
- explain why concurrent task access is safe;
- describe what happens when the queue is full;
- explain how worker goroutines terminate;
- identify possible race conditions or deadlocks;
- explain what would happen if a selected part of your synchronization code were removed or changed.

Questions will be based primarily on your submitted implementation.

You must understand and be able to explain all parts of your solution.

### Assessment

Maximum: **5 points**

- Implementation and functional correctness: up to **2.5 points**
- Individual code defense: up to **2.5 points**

The individual code defense is mandatory.

If you submit the implementation but do not complete the defense, you receive **0/5** for the checkpoint.

Generative AI tools may be used during SIS project development, according to the course AI policy. However, you remain responsible for reviewing, verifying, understanding, and defending all submitted code.

---

## 13. Final Checklist

Before submission, verify that:

- [ ] My Checkpoint 1 functionality still works.
- [ ] My application uses at least 2 worker goroutines.
- [ ] My workers receive tasks through a bounded channel-based queue.
- [ ] My workers can process independent tasks concurrently.
- [ ] My application handles a full queue correctly.
- [ ] Task states are updated correctly.
- [ ] Duplicate or invalid task processing is prevented.
- [ ] Shared mutable data is synchronized.
- [ ] Worker goroutines terminate correctly.
- [ ] New task submissions are rejected after shutdown.
- [ ] Accepted tasks are not silently lost during normal shutdown.
- [ ] My automated tests cover the required concurrent behavior.
- [ ] My README explains the concurrency design.
- [ ] I can explain my implementation during the individual code defense.

Run the following commands:

```bash
gofmt -l .
go vet ./...
go test ./...
go test -race ./...
```

`gofmt -l .` should produce no output.

All other commands must complete successfully.

**The main objective of Checkpoint 2 is not to use as many goroutines as possible. It is to demonstrate that you can design, implement, and reason about a correct concurrent Go application.**
