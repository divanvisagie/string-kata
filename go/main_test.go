package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given empty string",
			args: args{""},
			want: 0,
		},
		{
			name: "given 1",
			args: args{"1"},
			want: 1,
		},
		{
			name: "given 1,2",
			args: args{"1,2"},
			want: 3,
		},
		{
			name: "given 1,2,3",
			args: args{"1,2,3"},
			want: 6,
		},
		{
			name: "given 1\n2,3",
			args: args{"1\n2,3"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.input); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
