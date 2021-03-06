package main

import (
	fb "fizzbuzz/fizzbuzzer"
	"fmt"
	"os"
)

func main() {
	const from = 1
	const to = 101

	fbs, err := fb.FizzBuzzRange(from, to)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, v := range fbs {
		if v.Text != "" {
			fmt.Println(v.Text)
		} else {
			fmt.Println(v.Num)
		}
	}
}
