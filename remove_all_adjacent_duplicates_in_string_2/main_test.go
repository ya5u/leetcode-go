package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ex1",
			args: args{
				s: "abcd",
				k: 2,
			},
			want: "abcd",
		},
		{
			name: "ex2",
			args: args{
				s: "deeedbbcccbdaa",
				k: 3,
			},
			want: "aa",
		},
		{
			name: "ex3",
			args: args{
				s: "pbbcggttciiippooaais",
				k: 2,
			},
			want: "ps",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
