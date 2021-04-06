package permute

import (
	"fmt"
)

func Permute(ss []string) [][]string {
	if len(ss) == 2 {
		return [][]string{ss, swap(ss)}
	}
	if len(ss) == 3 {
		a := [][]string{}
		for i := range ss {
			print(i) // TODO remove
			fmt.Println(a)

			xs := remove(append([]string{}, ss...), i)
			fmt.Println(xs)

			sss := Permute(xs)
			fmt.Println(sss)

			a = append(
				a,
				append([]string{ss[i]}, sss[0]...))
			a = append(
				a,
				append([]string{ss[i]}, sss[1]...))
		}
		return a
	}
	if len(ss) > 3 {
		return append(
			[][]string{{ss[0]}}, Permute(ss[1:])...)
	}
	return [][]string{ss}
}

func swap(ss []string) []string {
	return []string{ss[1], ss[0]}
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
