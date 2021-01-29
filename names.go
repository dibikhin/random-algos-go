// Refined, packaged

package names

type Names []string

func Unique(a, b Names) Names {
	return unique(a, b)
}

// private below

func unique(a, b []string) []string {
	m := make(map[string]struct{})
	insert(a, m)
	insert(b, m)
	return slice(m)
}

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
