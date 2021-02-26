package permute

import (
	"fmt"
)

func Permute(ss []string) [][]string {
	if len(ss) == 2 {
		return [][]string{ss, swap(ss)}
	}
	if len(ss) > 2 {
		p := Permute(ss[1:])
		fmt.Println(p)
		a := append(
			[][]string{ss},
			[][]string{
				append([]string{ss[0]}, p[1]...)}...,
		)
		return a
	}
	return [][]string{ss}
}

func swap(ss []string) []string {
	return []string{ss[1], ss[0]}
}
