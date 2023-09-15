package main

import "testing"

func Test_equal(t *testing.T) {

	type args struct {
		a []int
		b []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"два массива равны",
			args{
				a: []int{2, 3, 4},
				b: []int{2, 3, 4},
			},
			true},
		{"два массива не равны",
			args{
				a: []int{2, 3, 5},
				b: []int{2, 9, 4},
			},
			false},
		{"у массивов разная длина ",
			args{
				a: []int{2, 3, 5, 0},
				b: []int{2, 9, 4},
			},
			false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
