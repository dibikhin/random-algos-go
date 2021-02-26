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
		{"empty", []string{}, [][]string{{}}},
		{"one string", []string{"a"}, [][]string{{"a"}}},

		{"two strings", []string{"a", "b"}, [][]string{{"a", "b"}, {"b", "a"}}},

		{"three strings", []string{"a", "b", "c"}, [][]string{
			{"a", "b", "c"}, {"a", "c", "b"}, {"c", "b", "a"},
			{"b", "c", "a"}, {"c", "a", "b"}, {"c", "b", "a"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permute(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
