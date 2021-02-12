// Refined

package fizzbuzzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	from int
	to   int
}

var dummy []FizzBuzz

var tests = []struct {
	name    string
	args    args
	want    []FizzBuzz
	wantErr bool
}{
	{"from 1 to 0", args{1, 0}, []FizzBuzz(nil), true},
	{"from 1 to 1", args{1, 1}, []FizzBuzz(nil), true},
	{"from 1 to 2", args{1, 2}, []FizzBuzz{{1, ""}}, false},
	{"from 1 to 4", args{1, 4}, []FizzBuzz{{1, ""}, {2, ""}, {3, "Fizz"}}, false},
	{"from 1 to 6", args{1, 6}, []FizzBuzz{{1, ""}, {2, ""}, {3, "Fizz"}, {4, ""}, {5, "Buzz"}}, false},
	{
		"from 11 to 16",
		args{11, 16},
		[]FizzBuzz{{11, ""}, {12, "Fizz"}, {13, ""}, {14, ""}, {15, "FizzBuzz"}},
		false,
	},
}

func TestFizzBuzzRange(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FizzBuzzRange(tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("FizzBuzzRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("FizzBuzzRange() = %v %T, want %v %T", got, tt.want, got, tt.want)
			}
		})
	}
}

// Just for fun too
func BenchmarkFizzBuzzRange(b *testing.B) {
	var r []FizzBuzz
	tt := tests[len(tests)-1]
	for i := 0; i < b.N; i++ {
		r, _ = FizzBuzzRange(tt.args.from, tt.args.to)
	}
	dummy = r // prevent inlining
}
