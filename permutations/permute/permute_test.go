package permute

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want [][]string
	}{
		// {"empty", []string{}, [][]string{{}}},
		// {"one string", []string{"a"}, [][]string{{"a"}}},

		// {"two strings", []string{"a", "b"}, [][]string{{"a", "b"}, {"b", "a"}}},

		{"three strings", []string{"a", "b", "c"}, [][]string{
			{"a", "b", "c"}, {"a", "c", "b"}, {"b", "a", "c"},
			{"b", "c", "a"}, {"c", "a", "b"}, {"c", "b", "a"},
		}},
		// {"four strings", []string{"a", "b", "c", "d"}, [][]string{
		// 	{"a", "b", "c", "d"}, {"a", "c", "b", "d"}, {"d", "c", "a", "b"},
		// 	{"a", "b", "d", "c"}, {"a", "c", "d", "b"}, {"d", "c", "b", "a"},
		// 	{"b", "a", "c", "d"}, {"c", "d", "a", "b"}, {"d", "b", "a", "c"},
		// 	{"b", "a", "d", "c"}, {"c", "d", "b", "a"}, {"d", "b", "c", "a"},

		// 	{"b", "c", "d", "a"}, {"c", "a", "b", "d"}, {"b", "d", "c", "a"},
		// 	{"b", "c", "a", "d"}, {"c", "a", "d", "b"}, {"b", "d", "a", "c"},
		// 	{"c", "b", "d", "a"}, {"d", "a", "b", "c"}, {"a", "d", "c", "b"},
		// 	{"c", "b", "a", "d"}, {"d", "a", "c", "b"}, {"a", "d", "b", "c"},
		// }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permute(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
