package main

import "fmt"

func main() {
	fmt.Println(factorial(0)) // 1
	fmt.Println(factorial(1)) // 1
	fmt.Println(factorial(2)) // 2
	fmt.Println(factorial(3)) // 6
	fmt.Println(factorial(4)) // 24

	fmt.Println(factorial(20)) // 2432902008176640000
	fmt.Println(factorial(21)) // 14197454024290336768
	fmt.Println(factorial(22)) // 17196083355034583040

	// fmt.Println(factorial(23)) // 8128291617894825984 WARN: overflow
}

func factorial(n uint64) uint64 {
	if n > 22 {
		panic("uint64 overflow")
	}
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
