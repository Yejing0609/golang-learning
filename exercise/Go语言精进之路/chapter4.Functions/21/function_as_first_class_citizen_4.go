package main

import "fmt"

func times(x, y int) int {
	return x * y
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}
func main() {
	timesTwo := partialTimes(2)
	fmt.Println(timesTwo(5))
}
