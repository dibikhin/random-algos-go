// Naive

package main

import "fmt"

func main() {
	fmt.Println("Started.")

	const N = 100
	for n := 1; n < N; n++ {
		if n%3 == 0 && n%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if n%5 == 0 {
			fmt.Println("Buzz")
		} else if n%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(n)
		}
	}
	fmt.Println("Done.")
}
