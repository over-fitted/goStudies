package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var first = -1
	var second = 0
	var temp int
	return func() int {
		if (first == -1) {
			first = 0
			return 0
		} 
		if (second == 0) {
			second = 1
			return second
		}
		temp = first + second
		first = second
		second = temp
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}