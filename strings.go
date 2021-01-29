// Refined, packaged

package strings

func Unique(a, b []string) []string {
	m := make(map[string]struct{})
	insert(a, m)
	insert(b, m)
	return slice(m)
}

// private below

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
