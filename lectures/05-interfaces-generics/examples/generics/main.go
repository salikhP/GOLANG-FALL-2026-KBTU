package main

import "fmt"

// go run ./lecture-05-slices/examples/slice-copy

func main() {
	sumExample()
	searchExample()
}

func sumExample() {
	floats := []float64{1.1, 2.2, 3.3}
	integers := []int64{1, 2, 3}

	fmt.Println(sum(floats))          // valid
	fmt.Println(sum[int64](integers)) // also valid

	fmt.Println(sumWithInterface(floats))

	// type approximation
	customIntegers := []CustomInt{1, 2, 3}
	fmt.Println(sumWithApproximation(customIntegers))
}

func searchExample() {
	type User struct {
		name     string
		age      int64
		jobTitle string
	}

	integers := []int64{1, 2, 3, 4, 5}

	fmt.Println("int:", search(integers, 3))

	users := []User{
		{name: "Alice", age: 30, jobTitle: "Engineer"},
		{name: "Bob", age: 25, jobTitle: "Designer"},
		{name: "Charlie", age: 35, jobTitle: "Manager"},
	}

	fmt.Println("structs:", search(users, User{
		name: "Bob", age: 25, jobTitle: "Designer",
	}))
}

/*
Type parameter T
'|' - operator union
*/
func sum[T int64 | float64](numbers []T) T {
	var sum T
	for _, number := range numbers {
		sum += number
	}

	return sum
}

/*
Type parameter T
"comparable" - operator interface
*/
func search[T comparable](slice []T, item T) bool {
	for _, value := range slice {
		if item == value {
			return true
		}
	}

	return false
}

type Number interface {
	int64 | float64
}

type NumberWithApproximation interface {
	~int64 | float64
}

func sumWithInterface[T Number](numbers []T) T {
	var sum T
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func sumWithApproximation[T NumberWithApproximation](numbers []T) T {
	var sum T
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// Numbers generic type
type Numbers[T Number] []T

type NumberApproximate interface {
	int64 | float64
}

type CustomInt int64

func (c CustomInt) IsPositive() bool {
	return c > 0
}
