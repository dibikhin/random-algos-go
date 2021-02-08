package main

import fb "fizzbuzz/fizzbuzzer"

func main() {
	const from = 1
	const to = 101

	fbs, err := fb.FizzBuzzRange(from, to)
	if err != nil {
		panic(err)
	}
	for _, v := range fbs {
		if v.Text != "" {
			println(v.Text)
		} else {
			println(v.Num)
		}
	}
}
