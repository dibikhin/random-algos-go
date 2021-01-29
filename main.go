// Refined, packaged

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		strings.Unique(
			[]string(nil),
			Names{},
		))
	fmt.Println(
		strings.Unique(
			Names{},
			Names{},
		))
	fmt.Println(
		strings.Unique(
			Names{},
			Names{"Ava", "Emma", "Olivia"},
		))
	fmt.Println(
		strings.Unique(
			Names{"Ava", "Emma", "Olivia"},
			Names{},
		))

	// should print Ava, Emma, Olivia, Sophia, any order
	fmt.Println(
		strings.Unique(
			Names{"Ava", "Emma", "Olivia"},
			Names{"Olivia", "Sophia", "Emma"},
		))
	fmt.Println(
		strings.Unique(
			Names{"Ava", "Emma", "Olivia"},
			Names{"Sophia"},
		))
}
