package main

type fn func()

func (f fn) times(n int) {
	// It's a variant of `[N]struct{}{}` when N isn't `const`.
	// Can use https://pkg.go.dev/github.com/bradfitz/iter otherwise :)
	for range make([]struct{}, n) {
		f()
	}
}

func main() {
	c := 0
	var f fn = func() {
		println(c)
		c++
	}
	f.times(5)
}
