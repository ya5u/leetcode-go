package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_powerfulIntegers(t *testing.T) {
	type args struct {
		x     int
		y     int
		bound int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				x:     2,
				y:     3,
				bound: 10,
			},
			want: []int{
				2, 3, 4, 5, 7, 9, 10,
			},
		},
		{
			name: "ex2",
			args: args{
				x:     3,
				y:     5,
				bound: 15,
			},
			want: []int{
				2, 4, 6, 8, 10, 14,
			},
		},
		{
			name: "ex3",
			args: args{
				x:     2,
				y:     1,
				bound: 10,
			},
			want: []int{
				2, 3, 5, 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := powerfulIntegers(tt.args.x, tt.args.y, tt.args.bound)
			if diff := cmp.Diff(got, tt.want, cmpopts.SortSlices(func(x, y int) bool { return x < y })); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
