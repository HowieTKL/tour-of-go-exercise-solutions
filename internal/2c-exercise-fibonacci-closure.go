package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		first, second = second, first+second
		return second - first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 11; i++ {
		fmt.Println(i, f())
	}
}
