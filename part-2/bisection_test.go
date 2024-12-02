package main

import (
	"math"
	"testing"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func TestBisection_Sqrt(t *testing.T) {
	tests := []struct {
		name      string
		value     float64
		want      float64
		threshold float64
		precision uint
		maxiter   uint
	}{
		{
			name:      "m=2",
			value:     2,
			threshold: 0.0001,
			want:      1.4142,
			precision: 4,
			maxiter:   32,
		},
		{
			name:      "m=2",
			value:     2,
			threshold: 0.00000001,
			want:      1.41421356,
			precision: 8,
			maxiter:   32,
		},
		{
			name:      "m=79543",
			value:     79543,
			threshold: 0.00000001,
			want:      282.03368593,
			precision: 8,
			maxiter:   64,
		},
		{
			name:      "maxiter reached",
			value:     2,
			threshold: 0.00000001,
			want:      -1,
			precision: 8,
			maxiter:   8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewBisection(tt.threshold, tt.maxiter, false, nil)
			got := roundFloat(s.Calc(tt.value), tt.precision)
			if got != tt.want {
				t.Errorf("Bisection.Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSqrtBisection2(b *testing.B) {
	value := 2.0
	threshold := 0.000001
	maxiter := uint(64)
	s := NewBisection(threshold, maxiter, false, nil)

	for i := 0; i < b.N; i++ {
		s.Calc(value)
	}
}

func BenchmarkSqrtBisection16(b *testing.B) {
	value := 15.0
	threshold := 0.000001
	maxiter := uint(64)
	s := NewBisection(threshold, maxiter, false, nil)

	for i := 0; i < b.N; i++ {
		s.Calc(value)
	}
}

func BenchmarkSqrtBisection32(b *testing.B) {
	value := 32.0
	threshold := 0.000001
	maxiter := uint(64)
	s := NewBisection(threshold, maxiter, false, nil)

	for i := 0; i < b.N; i++ {
		s.Calc(value)
	}
}

func BenchmarkSqrtBisection79543(b *testing.B) {
	value := 79543.0
	threshold := 0.000001
	maxiter := uint(64)
	s := NewBisection(threshold, maxiter, false, nil)

	for i := 0; i < b.N; i++ {
		s.Calc(value)
	}
}
func BenchmarkSqrtBisection6632888162(b *testing.B) {
	value := 6632888162.0
	threshold := 0.000001
	maxiter := uint(256)
	s := NewBisection(threshold, maxiter, false, nil)

	for i := 0; i < b.N; i++ {
		s.Calc(value)
	}
}
