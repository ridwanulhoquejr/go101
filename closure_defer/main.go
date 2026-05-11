package main

import (
	"fmt"
	"time"
)

// with closure we can create stateful func wihtout having a global var
func makeCounter() func() int {
	// func level global scope
	counter := 0

	return func() int {
		counter++ // increases and goes to heap
		return counter
	}
}

// defer
func slowOperation() {
	start := time.Now()
	fmt.Println("timer started!")

	defer func() {
		fmt.Printf("Deffered function are called %v\n", time.Since(start))
	}()

	fmt.Println("delayed 3 seconds for defer statement")
	time.Sleep(time.Second * 3)
}

// defer block wihtout named return
func deferWithOutNamedReturn(x int) int {
	result := 0

	defer func() {
		result += x
	}()

	// but when returns it sent wihtout added value of x+result; sent only declared value
	return result
}

// defer block wiht named return
func deferWithNamedReturn(x int) (result int) {
	result = 0

	defer func() {
		result += x
	}()

	// but when returns it sent with added value of x+result;
	return
}

// closure exampels
func makeAccumulator() func(int) int {
	sum := 0

	return func(i int) int {
		sum += i
		return sum
	}
}

// defer exmaples
func timed(name string, fn func()) {
	start := time.Now()
	defer func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}()
	fn()
}

func main() {

	// clousures
	c1 := makeCounter()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())

	c2 := makeCounter()
	fmt.Println(c2())
	fmt.Printf("c1=%d\n c2=%d\n c2=%d", c1(), c2(), c2())
	slowOperation()

	// call the defferefunc
	result1 := deferWithOutNamedReturn(15)
	fmt.Printf("Deffered retun value from no named reutrn is : %d\n", result1)

	result2 := deferWithNamedReturn(15)
	fmt.Printf("Deffered retun value is : %d", result2)

	// exercises
	ac := makeAccumulator()
	fmt.Println(ac(5))
	fmt.Println(ac(5))

	timed("Slow operartion", func() {
		time.Sleep(100 * time.Millisecond)
	})

}
