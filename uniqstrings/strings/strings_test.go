package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	name string
	a    []string
	b    []string
	want []string
}{
	{
		"first is nil",
		[]string(nil),
		[]string{},
		[]string{},
	},
	{
		"both empty",
		[]string{},
		[]string{},
		[]string{},
	},
	{
		"first empty",
		[]string{},
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Ava", "Emma", "Olivia"},
	},
	{
		"second empty",
		[]string{"Ava", "Emma", "Olivia"},
		[]string{},
		[]string{"Ava", "Emma", "Olivia"},
	},
	{
		"equal len",
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"},
		[]string{"Ava", "Emma", "Olivia", "Sophia"},
	},
	{
		"diff len",
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Sophia"},
		[]string{"Ava", "Emma", "Olivia", "Sophia"},
	},
}

func TestStringsUniq(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uniq(tt.a, tt.b)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
