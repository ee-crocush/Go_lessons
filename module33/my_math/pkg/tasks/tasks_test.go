package tasks

import (
	"math"
	"sort"
	"testing"
)

func TestCountLetters(t *testing.T) {
	type args struct {
		s string
		l rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Пример строки с английскими буквами",
			args: args{
				s: "AbgberrwAdfggdCAA",
				l: 'A',
			},
			want: 4,
		},
		{
			name: "Пример строки с русскими буквами",
			args: args{
				s: "ЫпавблоалваопМавпдМываПМ",
				l: 'М',
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := CountLetters(tt.args.s, tt.args.l); got != tt.want {
					t.Errorf("CountLetters() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SomeMath(float64(i))
	}
}

func BenchmarkInts(b *testing.B) {
	nums := []int{1, 5, 0, -5, 6, 2, 11, 7, 9, 3}
	for i := 0; i <= b.N; i++ {
		sort.Ints(nums)
	}
}

func TestSqrt(t *testing.T) {
	got := math.Sqrt(9)
	want := 3.0
	if got != want {
		t.Errorf("math.Sqrt(9) = %v, want %v", got, want)
	}
}
