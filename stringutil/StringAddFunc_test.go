package stringtil

import (
	"strings"
	"testing"
)

//TODO test the performance
func TestStringAddFunc(t *testing.T) {
	ts := []string{"abcd", "efg", "hij", "klm", "nop", "lmn", "ovw", "xyz", "abcd", "efg", "hij", "klm", "nop", "lmn", "ovw", "xyz"}
	if s := StringAddFunc(ts); s != "abcd efg hij klm nop lmn ovw xyz abcd efg hij klm nop lmn ovw xyz" {
		t.Errorf("StringAddFunc is not %s/n", s)
	}

}

func TestStringJoin(t *testing.T) {
	ts := []string{"abcd", "efg", "hij", "klm", "nop", "lmn", "ovw", "xyz", "abcd", "efg", "hij", "klm", "nop", "lmn", "ovw", "xyz"}
	if s := StringJoin(ts); s != "abcd efg hij klm nop lmn ovw xyz abcd efg hij klm nop lmn ovw xyz" {
		t.Errorf("StringJoin is not %s/n", s)
	}
}

func BenchmarkStringAddFunc(b *testing.B) {
	ts := []string{"abcd", "efg", "hij", "klm", "nop", "lmn", "ovw", "xyz", "abcd", "efg", "hij", "klm", "nop", "lmn", "ovw", "xyz"}
	b.N = 10000000
	for index := 0; index < b.N; index++ {
		StringAddFunc(ts)
	}
}

func BenchmarkStringJoin(b *testing.B) {
	ts := []string{"abcd", "efg", "hij", "klm", "nop", "lmn", "ovw", "xyz", "abcd", "efg", "hij", "klm", "nop", "lmn", "ovw", "xyz"}
	b.N = 10000000
	for index := 0; index < b.N; index++ {
		StringJoin(ts)
	}
}

func StringAddFunc(ts []string) string {
	s, sep := "", ""
	for _, value := range ts {
		s += sep + value
		sep = " "
	}
	return s
}

func StringJoin(ts []string) string {
	sep := " "
	return strings.Join(ts, sep)
}
