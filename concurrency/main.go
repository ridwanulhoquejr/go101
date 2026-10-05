package main

import (
	"fmt"
	"time"
)

// ============================================================
// TOPIC: Concurrency - Goroutines, Channels, Context
// ============================================================

// --- EXAMPLE ---

func producer(ch chan<- int) {
	ch <- 34
}

func consumer(ch <-chan int) {
	fmt.Println(<-ch)
}

func generate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		// to close the chan we have created
		defer close(out)

		// loop through the nums, then assing the values into the out chan
		for _, val := range nums {
			out <- val
		}
	}()

	return out
}

// receive only channel params and returns it
func square(in <-chan int) <-chan int {

	nums := make(chan int)
	go func() {
		defer close(nums)

		// loop through the param chan
		for n := range in {
			nums <- n * n
		}
	}()

	return nums
}

func writeOnly(c chan<- int) {}

func main() {

	done := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("my work is finshied, signalling done from newly spawned goroutine")

		done <- struct{}{}
	}()

	<-done // main blocks here untill done signal arrives

	// readChan := make(chan int)
	// writeChan := make(chan int)

	intChann := make(chan int)

	// go func() {
	// 	intChann <- 11
	// }()

	// chanValue, ok := <-intChann

	// fmt.Println(chanValue)
	// fmt.Println(ok)
	// close(intChann)

	// fmt.Println("after closing the channel")

	// chanValue, ok = <-intChann

	// fmt.Println(chanValue)
	// fmt.Println(ok)

	// go func() {
	// 	defer close(intChann)
	// 	for i := range 5 {
	// 		intChann <- i
	// 	}
	// 	// close(intChann)
	// }()

	// for j := range intChann {
	// 	fmt.Println(j)
	// }

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := range 3 {
			ch1 <- i
		}
		// time.Sleep(1 * time.Second)
	}()

	go func() {
		defer close(ch2)
		for i := range 3 {
			ch2 <- i
		}
		// time.Sleep(1 * time.Second)
	}()

	/*
		select takes one value from whichever channel is ready.
		Loop it to keep taking values. When a case reports ok == false, that channel is done — nil it so select stops offering it.
		When all are nil, the condition fails and the loop exits.
	*/

	for ch1 != nil || ch2 != nil {
		select {
		case value, ok := <-ch1:
			if !ok {
				// if ch1 is it has
				ch1 = nil
				continue
			}
			fmt.Printf("ch1 value: %d\n", value)

		case value, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Printf("ch2 value: %d\n", value)
			// default:
			// 	fmt.Println("all gorutine are asleep")
		}
	}

	return

	writeOnly(intChann)

	go func() {
		fmt.Println("Hello i am in another goroutine")
		time.Sleep(3 * time.Second)
		fmt.Println("My work is done")

		done <- struct{}{}
	}()

	value := <-done
	fmt.Println("Received signal from done channel, quiting the program", value)

	fmt.Println("Generate a numbs then pass it to the func to send it to the chan then returns it\n")

	numbers := generate(1, 2, 3, 4, 5)
	result := square(numbers)

	for val := range result {
		fmt.Println("Returned channel value:", val)
	}

	// 	var wg sync.WaitGroup

	// 	wg.Add(1)

	// 	// --- GOROUTINES + CHANNELS ---
	// 	// Channels let goroutines communicate safely
	// 	results := make(chan string, 3) // buffered channel

	// 	go fetchService("payments", results)
	// 	go fetchService("analytics", results)
	// 	go fetchService("notifications", results)

	// 	// Collect results
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println("Result:", <-results)
	// 	}

	// 	// --- CONTEXT WITH TIMEOUT ---
	// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// 	defer cancel()

	// 	ch := make(chan string, 1)
	// 	go func() {
	// 		time.Sleep(500 * time.Millisecond)
	// 		ch <- "done"
	// 	}()

	// 	select {
	// 	case msg := <-ch:
	// 		fmt.Println("Got:", msg)
	// 	case <-ctx.Done():
	// 		fmt.Println("Timeout:", ctx.Err())
	// 	}

	// 	// --- PIPELINE PATTERN ---
	// 	pipeline := make(chan int)
	// 	squares := make(chan int)

	// 	go producer(pipeline, 5)
	// 	go squareWorker(pipeline, squares)

	// 	// Consumer runs in main goroutine
	// 	for v := range squares {
	// 		fmt.Println("Square:", v)
	// 	}
	// }

	// func fetchService(name string, out chan<- string) {
	// 	delay := time.Duration(100+rand.Intn(300)) * time.Millisecond
	// 	time.Sleep(delay)
	// 	out <- fmt.Sprintf("%s responded in %s", name, delay)
	// }

	// func producer(out chan<- int, count int) {
	// 	defer close(out)
	// 	for i := 1; i <= count; i++ {
	// 		out <- i
	// 	}
	// }

	//	func squareWorker(in <-chan int, out chan<- int) {
	//		defer close(out)
	//		for n := range in {
	//			out <- n * n
	//		}
}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Launch 5 goroutines that each print their index.
//     Use sync.WaitGroup to wait for all to finish.
//
// Q2: Create a channel of ints. Send numbers 1-10 in a
//     goroutine, close the channel, then range over it
//     in main to print them.
//
// Q3: Write a fan-out pattern: one producer sends work to
//     3 worker goroutines via a shared channel.
//
// Q4: Use `select` with two channels and a timeout:
//       ch1 sends after 1s, ch2 sends after 2s,
//       timeout after 1.5s. What gets printed?
//
// Q5: What happens if you send to an unbuffered channel
//     with no receiver? (Answer: deadlock!)
//     Try: ch := make(chan int); ch <- 1
//
// Q6: Use context.WithCancel() to cancel a long-running
//     goroutine. The goroutine should check ctx.Done()
//     in a select loop.
//
// ============================================================
