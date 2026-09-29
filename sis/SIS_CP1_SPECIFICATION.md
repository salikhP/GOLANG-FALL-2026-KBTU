# SIS Project - Checkpoint 1 Specification

**Course:** Golang Application Development  
**Checkpoint:** 1 - Core Go  
**Points:** 5  
**Project type:** Individual cumulative project

## Project Overview

The SIS project is an individual Go application that will be developed incrementally throughout the semester.

The application is built around the concept of a **task**. 
In this project, a task means a **unit of work processed by the application.**
It does **not** necessarily mean a to-do item or a task-management application.

For example, a task may represent:

- analyzing a piece of text;
- performing a calculation;
- validating structured input;
- transforming some data;
- parsing and analyzing a URL;
- calculating statistics from a set of values;
- processing another small domain-specific request of your choice.

Each task should contain some domain-specific input and have a lifecycle represented by its state.

**For example:**

Text analyzer:
```
Input:
"Hello Go world"

Processing (the task):
analyze the text

Result:
- words: 3
- characters: 14
```

Statistics Calculator
```
Input:
[10, 20, 30]

Processing:
Calculate basic statistics

Result:
Min: 10
Max: 30
Average: 20
```

Each task should therefore have:
1. Input
2. Processing logic
3. Result (Output)

In addition, the application must track the current state of each task.
The exact domain, internal architecture, package structure, type names, and implementation details are your choice, as long as the requirements of each checkpoint are satisfied.

_If you are unsure whether your project idea fits these requirements, discuss it with the instructor._

## 1. Goal

Checkpoint 1 establishes the foundation of the SIS project.

You will build a small synchronous, in-memory Go application that:
- models domain-specific tasks;
- validates task input;
- processes tasks using meaningful business logic;
- produces and stores processing results;
- tracks task state;
- handles errors correctly;
- stores multiple tasks in memory;
- contains automated unit tests.

The focus of this checkpoint is not on frameworks or infrastructure. 
The goal is to demonstrate correct use of the Go concepts covered in Weeks 1-7: 
- structs
- methods
- interfaces
- slices/maps
- value and pointer semantics
- error handling
- basic code organization

The same project will be extended in later checkpoints with concurrency and, eventually, an HTTP interface. 
Design the code so that it can evolve, but **do not implement future-checkpoint requirements yet.**

## 2. Freedom of implementation

There is **no required project template, package structure, function signature, or architecture** for Checkpoint 1.

You may organize the repository in any reasonable way. 
You are responsible for being able to explain your design decisions during the individual code defense.

Your implementation must satisfy the observable requirements in this specification.

## 3. Required functionality

Your application must have a clear domain model for a **Task**.

Each task must contain, at minimum:

- a unique identifier;
- meaningful domain-specific input or payload;
- a current state/status;
- a result produced by processing the task.

The identifier format is your choice.

For example, it may be:
- an integer;
- a string;
- a generated identifier;
- another reasonable identifier type.

The application must support the following behavior:
1. Create a new task with domain-specific input.
2. Validate the task input.
3. Reject invalid task input.
4. Assign each created task a unique identifier.
5. Store created tasks in memory.
6. Retrieve an existing task by its identifier.
7. Correctly handle a request for a task that does not exist.
8. Process a task using meaningful domain-specific logic.
9. Produce a result from the task input.
10. Associate the produced result with the corresponding task.
11. Update the task state according to its processing lifecycle.
12. Reject invalid task operations or invalid state transitions.
13. Correctly manage multiple independent tasks without mixing or overwriting their data or results.

## 4. Task Processing
A task must perform some meaningful operation on its input and produce a result.
The processing logic depends on your chosen domain.
It does not need to be complex, but it must represent actual application behavior.
Simply returning or copying the original input without meaningful processing is not sufficient.

Examples of reasonable processing logic include:
- calculating word, rune, or character statistics from text;
- calculating mathematical or statistical values;
- validating structured data against several rules;
- parsing input and extracting useful information;
- transforming data from one representation into another;
- calculating some domain-specific result.

The processing logic should be deterministic enough to be tested using automated unit tests.
Avoid choosing a domain that requires unnecessary external infrastructure for Checkpoint 1.
For example, **you DO NOT need:**
- external APIs;
- databases;
- message brokers;
- cloud services;
- external storage.

## 5. Input Validation
Your application must define at least one meaningful validation rule for task input based on the chosen domain.

Invalid input must be rejected using normal Go error handling.

For example:
- Text Analyzer: empty or whitespace-only text is invalid
- Statistics Calculator: an empty list of values is invalid
- URL Analyzer: an empty or malformed URL is invalid

Validation rules should make sense for your domain.
Do not introduce arbitrary validation rules only to satisfy this requirement.

## 6. Task Lifecycle

Each task must have a small lifecycle represented by its state/status.

At minimum, the lifecycle must distinguish:
1. a task that has been created but has not yet been processed;
2. a task that is currently being processed;
3. a task that has successfully completed processing.

For example: Pending -> Processing -> Completed

You may choose different names and representations.

State changes must correspond to the real task-processing lifecycle.

Invalid state transitions or invalid operations must be handled appropriately.

For example, processing an already completed task again may be considered an invalid operation unless your domain explicitly supports it.

Briefly document the task lifecycle and allowed transitions in your repository README.

## 7. Implementation Requirements

Your implementation must demonstrate appropriate use of:

- structs and methods;
- value and/or pointer receivers;
- at least one meaningful interface;
- slices and/or maps for in-memory data management;
- explicit error handling.

Expected application errors must be returned as errors. Do not use `panic` for normal application-level errors.

The application must store all tasks and results in memory and correctly handle multiple independent tasks.

A database or other external storage is not required.

---

## 8. Automated Tests

You must write automated tests using Go's standard `testing` package.

Your tests must cover, at minimum:

1. successful task creation;
2. invalid task input;
3. successful task retrieval;
4. retrieval of a nonexistent task;
5. successful task processing;
6. correctness of the produced result;
7. correct task state after processing;
8. at least one invalid operation or state transition;
9. multiple independent tasks.

Use at least one table-driven test where appropriate.

Tests must verify expected behavior and results, not merely execute code.

There is no mandatory code coverage percentage for Checkpoint 1.

---

## 9. Technical Requirements

Use **Go 1.27**.

Your repository must:

- be a valid Go module with a committed `go.mod`;
- compile successfully;
- be formatted with `gofmt`;
- pass `go vet ./...`;
- pass `go test ./...`.

Before submission, verify:

```bash
gofmt -l .
go vet ./...
go test ./...
