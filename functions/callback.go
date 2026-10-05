package main

/// A callback is a function you pass to another function as an argument, so that function can call it later.
// You hand over the "what to do," and the receiver decides when to run it.

/*
	The pieces:
	- fn func(int) int is the callback's type. Any function with that signature fits.
	- apply doesn't know what fn does. It only calls it.
	- You can pass a named function, a variable like double, or an inline anonymous function.
*/

// this apply function takes a slice of nums and a function to perform
// so here, in this operation we just run the func received from the param
// and return the output

// for example: we will double the output
func apply(nums []int, fn func(int) int) []int {

	out := make([]int, 0, len(nums))

	for _, num := range nums {
		out = append(out, fn(num))
	}

	return out
}
