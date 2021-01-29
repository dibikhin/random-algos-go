// Refined, standalone

package main

import "fmt"

type names []string

func insert(ss []string, m map[string]struct{}) {
	for i := range ss {
		m[ss[i]] = struct{}{}
	}
}

func slice(m map[string]struct{}) []string {
	var res []string
	for k := range m {
		res = append(res, k)
	}
	return res
}

func unique(a, b []string) []string {
	m := make(map[string]struct{})
	insert(a, m)
	insert(b, m)
	return slice(m)
}

func uniqueNames(a, b names) names {
	return unique(a, b)
}

func main() {
	fmt.Println(uniqueNames(
		[]string(nil),
		names{},
	))
	fmt.Println(uniqueNames(
		names{},
		names{},
	))
	fmt.Println(uniqueNames(
		names{},
		names{"Ava", "Emma", "Olivia"},
	))
	fmt.Println(uniqueNames(
		names{"Ava", "Emma", "Olivia"},
		names{},
	))

	// should print Ava, Emma, Olivia, Sophia, any order
	fmt.Println(uniqueNames(
		names{"Ava", "Emma", "Olivia"},
		names{"Olivia", "Sophia", "Emma"},
	))
	fmt.Println(uniqueNames(
		names{"Ava", "Emma", "Olivia"},
		names{"Sophia"},
	))
}
