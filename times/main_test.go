package main

import "testing"

func Test_fn_times(t *testing.T) {
	c := 0
	tests := []struct {
		name string
		f    fn
		want int
	}{
		{"zero", func() { c++ }, 0},
		{"five", func() { c++ }, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.times(tt.want)
			if c != tt.want {
				t.Errorf("Run %v times, want %v times", c, tt.want)
			}
		})
	}
}
