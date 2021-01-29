// Naive, standalone

package main

import "fmt"

func uniqNames(a, b []string) []string {
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
	fmt.Println(uniqNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
	fmt.Println(uniqNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Sophia"}))
}
