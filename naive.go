// Naive, standalone

package main

import "fmt"

func uniqueNames(a, b []string) []string {
	m := make(map[string]struct{})

	for i := range a {
		m[a[i]] = struct{}{}
	}
	for i := range b {
		m[b[i]] = struct{}{}
	}
	var result []string
	for k := range m {
		result = append(result, k)
	}
	return result
}

func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Sophia"}))
}
