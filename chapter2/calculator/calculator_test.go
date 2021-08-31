package calculator

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		nb1 int
		nb2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "add 2 zeros",
			args: args{0, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.nb1, tt.args.nb2); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	type args struct {
		nb1 int
		nb2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.args.nb1, tt.args.nb2); got != tt.want {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
