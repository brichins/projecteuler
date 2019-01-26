// Multiples of 3 and 5
//
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
package main

import "testing"

func TestSumMultiples(t *testing.T) {
	type args struct {
		multiples []int
		max       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"<10", args{[]int{3, 5}, 10}, 23},
		{"<1000", args{[]int{3, 5}, 1000}, 233168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumMultiples(tt.args.multiples, tt.args.max); got != tt.want {
				t.Errorf("SumMultiples() = %v, want %v", got, tt.want)
			}
		})
	}
}
