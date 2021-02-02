// Refined

package fizzbuzzer_refined

// FizzBuzz contains a number or string
type FizzBuzz struct {
	Num  int
	Text string
}

// FizzBuzzTill returns slice of fizz/buzz as of https://wiki.c2.com/?FizzBuzzTest
func FizzBuzzTill(n int) []FizzBuzz {
	// TODO check n

	// TOOO make ahead
	var res []FizzBuzz
	for i := 1; i < n; i++ {
		if i%15 == 0 {
			res = append(res, FizzBuzz{i, "FizzBuzz"})
		} else if i%5 == 0 {
			res = append(res, FizzBuzz{i, "Buzz"})
		} else if i%3 == 0 {
			res = append(res, FizzBuzz{i, "Fizz"})
		} else {
			res = append(res, FizzBuzz{i, ""})
		}
	}
	return res
}
