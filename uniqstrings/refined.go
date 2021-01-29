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

func uniq(a, b []string) []string {
	m := make(map[string]struct{})
	insert(a, m)
	insert(b, m)
	return slice(m)
}

func uniqNames(a, b names) names {
	return uniq(a, b)
}

func main() {
	fmt.Println(uniqNames(
		[]string(nil),
		names{},
	))
	fmt.Println(uniqNames(
		names{},
		names{},
	))
	fmt.Println(uniqNames(
		names{},
		names{"Ava", "Emma", "Olivia"},
	))
	fmt.Println(uniqNames(
		names{"Ava", "Emma", "Olivia"},
		names{},
	))

	// should print Ava, Emma, Olivia, Sophia, any order
	fmt.Println(uniqNames(
		names{"Ava", "Emma", "Olivia"},
		names{"Olivia", "Sophia", "Emma"},
	))
	fmt.Println(uniqNames(
		names{"Ava", "Emma", "Olivia"},
		names{"Sophia"},
	))
}
