// Package fizzbuzzer is for learning using packages inside modules only.
// Could have the project stucture plain as well.
package fizzbuzzer

import "errors"

// FizzBuzz contains a number or string.
// It's for keeping the numbers and indexes independent
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
	if to <= from || to < 0 || from < 0 {
		return nil, errors.New("from <= to OR from/to is < 0. But should be: from < to and both > 0")
	}

	// This optimizing is for fun only
	res := good(from, to)
	// res := wierd(to, from)
	// res := better(to, from)

	return res, nil
}

// Good but looks slow
// BenchmarkFizzBuzzRange-4   	 1470634	       814 ns/op	     368 B/op	       4 allocs/op
func good(from int, to int) []FizzBuzz {
	var res []FizzBuzz

	for i := from; i < to; i++ {
		switch {
		case i%15 == 0:
			res = append(res, FizzBuzz{i, "FizzBuzz"})
		case i%5 == 0:
			res = append(res, FizzBuzz{i, "Buzz"})
		case i%3 == 0:
			res = append(res, FizzBuzz{i, "Fizz"})
		default:
			res = append(res, FizzBuzz{i, ""})
		}
	}
	return res
}

// Fast but looks wierd
// BenchmarkFizzBuzzRange-4   	 5342506	       227 ns/op	     128 B/op	       1 allocs/op
func wierd(to int, from int) []FizzBuzz {
	s := make([]FizzBuzz, to-from)

	for i := from; i < to; i++ {
		switch {
		case i%15 == 0:
			s[i-from] = FizzBuzz{i, "FizzBuzz"}
		case i%5 == 0:
			s[i-from] = FizzBuzz{i, "Buzz"}
		case i%3 == 0:
			s[i-from] = FizzBuzz{i, "Fizz"}
		default:
			s[i-from] = FizzBuzz{i, ""}
		}
	}
	return s
}

// Fast and mainainable
// BenchmarkFizzBuzzRange-4   	 5475526	       217 ns/op	     128 B/op	       1 allocs/op
func better(to int, from int) []FizzBuzz {
	s := make([]FizzBuzz, 0, to-from) // NOTE the 0

	for i := from; i < to; i++ {
		switch {
		case i%15 == 0:
			s = append(s, FizzBuzz{i, "FizzBuzz"})
		case i%5 == 0:
			s = append(s, FizzBuzz{i, "Buzz"})
		case i%3 == 0:
			s = append(s, FizzBuzz{i, "Fizz"})
		default:
			s = append(s, FizzBuzz{i, ""})
		}
	}
	return s
}
