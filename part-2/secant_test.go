package main

import "testing"

func TestSqrtSecant(t *testing.T) {
	tests := []struct {
		name      string
		value     float64
		threshold float64
		maxiter   uint
		precision uint
		want      float64
	}{
		{
			name:      "m=2",
			value:     2.0,
			threshold: 0.00001,
			want:      1.4142,
			maxiter:   32,
			precision: 4,
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
			maxiter:   32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSecant(tt.threshold, tt.maxiter, false, nil)
			if got := roundFloat(s.Calc(tt.value), tt.precision); got != tt.want {
				t.Errorf("SqrtSecant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSqrtSecant2(b *testing.B) {
	value := 2.0
	threshold := 0.000001
	maxiter := uint(64)
	s := NewSecant(threshold, maxiter, false, nil)

	for i := 0; i < b.N; i++ {
		s.Calc(value)
	}
}

func BenchmarkSqrtSecant16(b *testing.B) {
	value := 15.0
	threshold := 0.000001
	maxiter := uint(64)
	s := NewSecant(threshold, maxiter, false, nil)

	for i := 0; i < b.N; i++ {
		s.Calc(value)
	}
}

func BenchmarkSqrtSecant32(b *testing.B) {
	value := 32.0
	threshold := 0.000001
	maxiter := uint(64)
	s := NewSecant(threshold, maxiter, false, nil)

	for i := 0; i < b.N; i++ {
		s.Calc(value)
	}
}

func BenchmarkSqrtSecant79543(b *testing.B) {
	value := 79543.0
	threshold := 0.000001
	maxiter := uint(64)
	s := NewSecant(threshold, maxiter, false, nil)

	for i := 0; i < b.N; i++ {
		s.Calc(value)
	}
}

func BenchmarkSqrtSecant6632888162(b *testing.B) {
	value := 6632888162.0
	threshold := 0.000001
	maxiter := uint(64)
	s := NewSecant(threshold, maxiter, false, nil)

	for i := 0; i < b.N; i++ {
		s.Calc(value)
	}
}
