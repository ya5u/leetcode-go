package main

import "testing"

func Test_leastBricks(t *testing.T) {
	type args struct {
		wall [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				wall: [][]int{
					{1, 2, 2, 1}, // 1, 3, 5, 6
					{3, 1, 2},    // 3, 4, 6
					{1, 3, 2},    // 1, 4, 6
					{2, 4},       // 2, 6
					{3, 1, 2},    // 3, 4. 6
					{1, 3, 1, 1}, // 1, 4, 5, 6
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastBricks(tt.args.wall); got != tt.want {
				t.Errorf("leastBricks() = %v, want %v", got, tt.want)
			}
		})
	}
}
