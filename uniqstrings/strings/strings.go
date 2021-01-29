// Refined, packaged

package strings

type void struct{}

// Uniq returns unique strings in any order
func Uniq(a, b []string) []string {
	m := make(map[string]void)
	insert(a, m)
	insert(b, m)
	return slice(m)
}

// private below

func insert(ss []string, m map[string]void) {
	for i := range ss {
		m[ss[i]] = void{}
	}
}

func slice(m map[string]void) []string {
	var res []string
	for k := range m {
		res = append(res, k)
	}
	return res
}
