// Naive

package main

func main() {
	println("Started.")

	const N = 100
	for n := 1; n < N; n++ {
		if n%3 == 0 && n%5 == 0 {
			println("FizzBuzz")
		} else if n%5 == 0 {
			println("Buzz")
		} else if n%3 == 0 {
			println("Fizz")
		} else {
			println(n)
		}
	}
	println("Done.")
}