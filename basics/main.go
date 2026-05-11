package main

import (
	"fmt"
)

// ============================================================
// TOPIC: Go Basics - Variables, Types, Control Flow, Strings
// ============================================================

// --- EXAMPLE ---

var course = "Go 101" // package-level variable

// a function which returns true if the provided argument number is even; otherwise false
func isEven(n int) bool {
	return n%2 == 0
}

// a funct which returns a interpolated string using fmt.Sprintf
func greet(name string) string {
	return fmt.Sprintf("Hello %s! Welcome to GO 101.", name)
}

// IOTA
type Weekday int8

const (
	Sunday Weekday = iota
	_
	Tuesday
)

/// defere functions

// named return function -- return does include the defer execution
func NamedReturnDeferFunc(param int64) (result int64) {
	// defer func() { result = param + 10 }() // IIF
	defer DeferTest()
	return // would return param + 10 as sum
}

// Type return function -- return does not consider defer execution
func TypedReturnDeferFunc(param int) int {
	result := 0
	// defer func() { result = param + 10 }() // IIF
	defer DeferTest()
	return result // would return 0
}

func DeferTest() {
	fmt.Println(Sunday)
	fmt.Println(Tuesday)
	fmt.Println("defe func executes")
}

func main() {

	// Q1: Declare a constant called `maxRetries` with value 5.
	//     Print it inside main().
	// Ans:
	const maxRetries = 5
	fmt.Println(maxRetries) // 5

	// Q2: Write a function `isEven(n int) bool` that returns true
	//     if n is even. Call it from main() with a few test values.
	// Ans:
	test1 := isEven(10)
	test2 := isEven(25)
	test3 := isEven(47)

	fmt.Printf("Test case 1 for argument 10: %v \n", test1) // true
	fmt.Printf("Test case 2 for argument 25: %v \n", test2) // false
	fmt.Printf("Test case 3 for argument 47: %v \n", test3) // false

	// Q3: Write a function `greet(name string) string` that returns
	//     "Hello, <name>! Welcome to Go 101." using fmt.Sprintf.
	// Ans:
	fmt.Println(greet("Ridwanul Hoque"))

	// Q4: Use a for loop to print numbers 1 to 10. Skip even
	//     numbers using `continue`.
	// Ans:
	for i := 0; i <= 10; i++ {
		if isEven(i) {
			continue
		} else {
			fmt.Println(i)
		}
	}

	// Q5: Write a switch statement that takes a day string
	//     ("Monday", "Saturday", etc.) and prints whether it's
	//     a weekday or weekend.
	// Ans:
	day := "Monday"

	switch day {
	case "Friday", "Saturday":
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	fmt.Printf("Named return defer func return value: %d\n", NamedReturnDeferFunc(1))
	fmt.Printf("Type return defer func return value: %d\n", TypedReturnDeferFunc(1))

	// Q6: What will this print? Why?
	//     for i := 0; i < 3; i++ {
	//         defer fmt.Println(i)
	//     }
	//     (Answer: 2, 1, 0 - defer is LIFO)
	/*
		Ans and Explanations:

			It will prints: 2, 1, 0

			Because, defer statement stores the codes in a Stack for later execution.
			So deferred block is not executed immidietly rather it waits to finish other executions of the function first.

			Since Stack used LIFO (Last In First Out) mechanism, the last deferred statement prints the first.
			Thus, 2 prints first followed by 1 and 0.
	*/

	// 	fmt.Println("Welcome to", course)

	// 	// Variables & short declaration
	// 	name := "Ada"
	// 	age := 25
	// 	isStudent := true
	// 	fmt.Printf("name=%s age=%d student=%t\n", name, age, isStudent)

	// 	// String formatting
	// 	greeting := fmt.Sprintf("Hello, %s!", strings.ToUpper(name))
	// 	fmt.Println(greeting)

	// 	// Multiple return values + error handling
	// 	result, err := add(21, 21)
	// 	if err != nil {
	// 		fmt.Println("add error:", err)
	// 		return
	// 	}
	// 	fmt.Println("21 + 21 =", result)

	// 	// Switch statement (no break needed in Go)
	// 	switch {
	// 	case result > 50:
	// 		fmt.Println("Big number!")
	// 	case result == 42:
	// 		fmt.Println("The answer to everything!")
	// 	default:
	// 		fmt.Println("Just a number:", result)
	// 	}

	// 	// For loop + defer (LIFO order)
	// 	for i := 0; i < 3; i++ {
	// 		defer fmt.Println("deferred:", i)
	// 		fmt.Println("loop iteration:", i)
	// 	}
	// }

	//	func add(a, b int) (int, error) {
	//		sum := a + b
	//		if sum < 0 {
	//			return 0, fmt.Errorf("negative sum: %d", sum)
	//		}
	//		return sum, nil
}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Declare a constant called `maxRetries` with value 5.
//     Print it inside main().
//
// Q2: Write a function `isEven(n int) bool` that returns true
//     if n is even. Call it from main() with a few test values.
//
// Q3: Write a function `greet(name string) string` that returns
//     "Hello, <name>! Welcome to Go 101." using fmt.Sprintf.
//
// Q4: Use a for loop to print numbers 1 to 10. Skip even
//     numbers using `continue`.
//
// Q5: Write a switch statement that takes a day string
//     ("Monday", "Saturday", etc.) and prints whether it's
//     a weekday or weekend.
//
// Q6: What will this print? Why?
//     for i := 0; i < 3; i++ {
//         defer fmt.Println(i)
//     }
//     (Answer: 2, 1, 0 - defer is LIFO)
//
// ============================================================
