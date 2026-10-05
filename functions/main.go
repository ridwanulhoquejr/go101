package main

import "fmt"

// ============================================================
// TOPIC: Functions - Multiple Returns, Variadic, Closures
// ============================================================

// --- EXAMPLE ---

func main() {

	// callback
	double := func(n int) int { return n * 2 }
	result := apply([]int{1, 2, 3}, double)

	fmt.Println("call back double function resutl", result)
	return
	// Basic function call
	sum := add(10, 32)
	fmt.Println("sum:", sum)

	// Multiple return values + error handling
	quotient, remainder, err := divide(100, 7)
	if err != nil {
		fmt.Println("divide error:", err)
		return
	}
	fmt.Printf("100 / 7 = %d remainder %d\n", quotient, remainder)

	// Variadic function
	total := variadicSum(1, 2, 3, 4, 5)
	fmt.Println("variadic total:", total)

	// Closures - function that captures outer variable
	counter := makeCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	// Anonymous function
	// double := func(n int) int { return n * 2 }
	fmt.Println("double 5:", double(5))
}

func add(a, b int) int {
	return a + b
}

func divide(dividend, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, fmt.Errorf("cannot divide by zero")
	}
	return dividend / divisor, dividend % divisor, nil
}

func variadicSum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Write a function `multiply(a, b int) int` and call it.
//
// Q2: Write a function `swap(a, b string) (string, string)`
//     that returns the two strings in reverse order.
//
// Q3: Write a function `max(nums ...int) int` that returns
//     the largest number from a variadic list.
//
// Q4: Write a closure `makeAdder(x int) func(int) int` that
//     returns a function which adds x to its argument.
//     Example: addFive := makeAdder(5); addFive(3) -> 8
//
// Q5: Write a function `apply(a, b int, op func(int,int) int) int`
//     that takes two ints and a function, and applies the
//     function. Pass in add, multiply, etc.
//
// Q6: Rewrite `divide` using named return values:
//     func divide(a, b int) (q int, r int, err error)
//
// ============================================================
