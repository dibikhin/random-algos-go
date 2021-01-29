// Refined, packaged

package main

import (
	"fmt"

	"./strings"
)

func main() {
	// should print [Ava Emma Olivia Sophia] any order
	fmt.Println(
		strings.Uniq(
			[]string{"Ava", "Emma", "Olivia"},
			[]string{"Sophia"},
		))
}
