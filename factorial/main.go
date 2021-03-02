package main

func main() {
	println(factorial(0)) // 1
	println(factorial(1)) // 1
	println(factorial(2)) // 2
	println(factorial(3)) // 6
	println(factorial(4)) // 24

	println(factorial(20)) // 2432902008176640000
	println(factorial(21)) // 14197454024290336768
	println(factorial(22)) // 17196083355034583040

	// println(factorial(23)) // 8128291617894825984 WARN: overflow
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
