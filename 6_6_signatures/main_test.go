package main

import (
	"reflect"
	"testing"
)

func Test_points(t *testing.T) {
	type args struct {
		segments []segment
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "asd",
			args: args{
				segments: []segment{
					{left: 1, right: 3},
					{left: 2, right: 5},
					{left: 3, right: 6},
				},
			},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := points(tt.args.segments); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("points() = %v, want %v", got, tt.want)
			}
		})
	}
}
