package main

import "fmt"

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from panic", r)
	// 	}
	// }()

	// fmt.Println("Starting the program")
	// panic("Something went worng")
	// fmt.Println("this line will not execute")

	fmt.Println(safeDivision(5, 2))
	fmt.Println(safeDivision(1, 0))
	fmt.Printf("program continues")
}

func safeDivision(a, b int64) (result int64) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recoveredd form panic")
			result = 0.0

			switch v := r.(type) {
			case string:
				fmt.Println("string panic", v)
			case error:
				fmt.Println("error panic", v)
			default:
				fmt.Println("unknown panic", v)
			}
		}
	}()

	return (a / b)
}
