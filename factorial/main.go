package main

func main() {
	println(factorial(0)) // 1
	println(factorial(1)) // 1
	println(factorial(2)) // 2
	println(factorial(3)) // 6
	println(factorial(4)) // 24

	println(factorial(20)) // 2432902008176640000
}

func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
