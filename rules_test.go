package xssvalidator

import (
	"errors"
	"testing"
)

func TestBracketRule_Check(t *testing.T) {
	b := BracketRule{}
	err := b.Check("[A] this (is) a {song} [D] [Am] i am [Adim] Am [f] [Dmin] ()")
	if !errors.Is(err, ErrBracketRule) {
		t.Errorf("Error should be %s, instead of %s", ErrBracketRule, err)
	}

	err = b.Check("TestData")
	if err != nil {
		t.Errorf("Error should be nil, instead of %s", err)
	}
}

func BenchmarkBracketRule_Check(b *testing.B) {
	br := BracketRule{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		br.Check("[A] this (is) a {song} [D] [Am] i am [Adim] Am [f] [Dmin] ()")
	}
}
