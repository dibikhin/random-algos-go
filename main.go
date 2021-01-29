// Refined, packaged

package main

import (
	"fmt"
	"names"
)

func main() {
	fmt.Println(
		names.Unique(
			[]string(nil),
			Names{},
		))
	fmt.Println(
		names.Unique(
			Names{},
			Names{},
		))
	fmt.Println(
		names.Unique(
			Names{},
			Names{"Ava", "Emma", "Olivia"},
		))
	fmt.Println(
		names.Unique(
			Names{"Ava", "Emma", "Olivia"},
			Names{},
		))

	// should print Ava, Emma, Olivia, Sophia, any order
	fmt.Println(
		names.Unique(
			Names{"Ava", "Emma", "Olivia"},
			Names{"Olivia", "Sophia", "Emma"},
		))
	fmt.Println(
		names.Unique(
			Names{"Ava", "Emma", "Olivia"},
			Names{"Sophia"},
		))
}
