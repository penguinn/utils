package pmath_test

import (
	"github.com/penguinn/utils/pmath"
	"testing"
)

func TestRound(t *testing.T) {
	type args struct {
		f float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "测试1",
			args: args{
				f: 54.48341896,
				n: 2,
			},
			want: 54.48,
		},
		{name: "测试2",
			args: args{
				f: 54.48841896,
				n: 2,
			},
			want: 54.49,
		},
		{name: "测试3",
			args: args{
				f: 54.47841896,
				n: 2,
			},
			want: 54.48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pmath.Round(tt.args.f, tt.args.n); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound2(t *testing.T) {
	type args struct {
		f float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "测试1",
			args: args{
				f: 54.48341896,
				n: 2,
			},
			want: 54.48,
		},
		{name: "测试2",
			args: args{
				f: 54.48841896,
				n: 2,
			},
			want: 54.49,
		},
		{name: "测试3",
			args: args{
				f: 54.47841896,
				n: 2,
			},
			want: 54.48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pmath.Round2(tt.args.f, tt.args.n); got != tt.want {
				t.Errorf("Round2() = %v, want %v", got, tt.want)
			}
		})
	}
}
