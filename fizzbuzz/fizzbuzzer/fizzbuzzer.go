// Refined

package fizzbuzzer

import "errors"

// FizzBuzz contains a number or string.
// It make the numbers and indexes independent
type FizzBuzz struct {
	Num  int
	Text string
}

// FizzBuzzRange returns slice of FizzBuzz as of https://wiki.c2.com/?FizzBuzzTest
// (suitable for ranging over).
//
// For example:
//
//	fbs, err := fb.FizzBuzzRange(1, 101)
//
// ... will create a slice of FizzBuzz from 1 to 101 excl. (100 pcs.)
//
func FizzBuzzRange(from, to int) ([]FizzBuzz, error) {
	if to < from || to < 0 || from < 0 {
		return nil, errors.New("from <= to, but should be: from < to and both positive")
	}
	var res []FizzBuzz

	for i := from; i <= to; i++ {
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
	return res, nil
}
