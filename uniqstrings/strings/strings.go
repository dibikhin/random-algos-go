// Refined, packaged

package strings

type void struct{}

type strSet map[string]void

// Uniq returns unique strings in any order
func Uniq(a, b []string) []string {
	set := make(strSet)
	insert(a, set)
	insert(b, set)
	return slice(set)
}

// private below

func insert(strs []string, set strSet) {
	for _, v := range strs {
		set[v] = void{}
	}
}

func slice(set strSet) []string {
	var res []string
	for k := range set {
		res = append(res, k)
	}
	return res
}
