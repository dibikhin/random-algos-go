// Refined

package fizzbuzzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzRange(t *testing.T) {
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name    string
		args    args
		want    []FizzBuzz
		wantErr bool
	}{
		{"from 1 to 0", args{1, 0}, []FizzBuzz(nil), true},
		{"from 1 to 1", args{1, 1}, []FizzBuzz{{1, ""}}, false},
		{"from 1 to 2", args{1, 2}, []FizzBuzz{{1, ""}, {2, ""}}, false},
		{"from 1 to 3", args{1, 3}, []FizzBuzz{{1, ""}, {2, ""}, {3, "Fizz"}}, false},
		{"from 1 to 5", args{1, 5}, []FizzBuzz{{1, ""}, {2, ""}, {3, "Fizz"}, {4, ""}, {5, "Buzz"}}, false},
		{"from 11 to 15", args{11, 15}, []FizzBuzz{{11, ""}, {12, "Fizz"}, {13, ""}, {14, ""}, {15, "FizzBuzz"}}, false},
	}
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
