package main

import (
	"fmt"
	"time"
)

func Test() {
	start := time.Now()

	fmt.Println("this is defer test line no 1: A")
	fmt.Println("this is defer test line no 2: B")
	fmt.Println("this is defer test line no 3: C")

	// defer fmt.Println("this is defer test line no 4 with defer keyword: D")
	// defer fmt.Println("this is defer test line no 4 with defer keyword: X")
	// defer fmt.Println("this is defer test line no 4 with defer keyword: Y")

	defer func() {
		fmt.Printf("operation took %v\n", time.Since(start))
	}()

	fmt.Println("this is defer test line no 5: E")
	fmt.Println("this is defer test line no 6: F")

}
