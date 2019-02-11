package main

import (
	"testing"
)

func Test_add(t *testing.T) {
	type args struct {
		input string
	}
	type TestParams struct {
		name string
		args args
		want int
		err  string
	}
	tests := []TestParams{
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
		{
			name: "given //;\n1;2",
			args: args{"//;\n1;2"},
			want: 3,
		},
		{
			name: "given nevative number",
			args: args{"1,4,-1"},
			want: 0,
			err:  "negatives not allowed: -1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := add(tt.args.input)

			if got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}

			if err != nil {
				if err.Error() != tt.err {
					t.Errorf("add() error = %v, want %v", err.Error(), tt.err)
				}
			}
		})
	}
}

func Test_processDelimiter(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name           string
		args           args
		wantCleanInput string
		wantDelim      rune
	}{
		{
			name:           "given no delimiter in string",
			args:           args{"1,2,3"},
			wantCleanInput: "1,2,3",
			wantDelim:      ',',
		},
		{
			name:           "given delimiter in string",
			args:           args{"//:\n1:2:3"},
			wantCleanInput: "1:2:3",
			wantDelim:      ':',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := processDelimiter(tt.args.input)
			if got != tt.wantCleanInput {
				t.Errorf("processDelimiter() gotCleanInput = %v, want %v", got, tt.wantCleanInput)
			}
			if got1 != tt.wantDelim {
				t.Errorf("processDelimiter() gotDelim = %v, want %v", got1, tt.wantDelim)
			}
		})
	}
}
