package tml

import (
	"testing"
)

func TestStrLen(t *testing.T) {
	s := "αβγ"
	if l := StrLen(s); l != 3 {
		t.Fatalf("Strlen is %v", l)
	}
	if l := len(s); l != 6 {
		t.Fatalf("len is %v", l)
	}
}

func BenchmarkStrLen(b *testing.B) {
	s := "α β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ ς σ τ υ φ χ ψ ω"
	for i := 0; i < b.N; i++ {
		_ = StrLen(s)
	}
}

func BenchmarkStrLenAlternative(b *testing.B) {
	s := "α β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ ς σ τ υ φ χ ψ ω"
	for i := 0; i < b.N; i++ {
		_ = len([]rune(s))
	}
}
