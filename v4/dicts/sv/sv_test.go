package sv

import (
	"strings"
	"testing"

	"github.com/linga-io/golem/v4"
)

func TestSwedishUsage(t *testing.T) {
	l, err := golem.New(New())
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		in  string
		out string
	}{
		{"Avtalet", "avtal"},
		{"avtalets", "avtal"},
		{"avtalens", "avtal"},
		{"Avtaletsadlkj", "Avtaletsadlkj"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := l.Lemma(tt.in)
			if got != tt.out {
				t.Errorf("Lemmatizer.Lemma() = %v, want %v", got, tt.out)
			}
			got = l.LemmaLower(strings.ToLower(tt.in))
			if got != strings.ToLower(tt.out) {
				t.Errorf("Lemmatizer.LemmaLower() = %v, want %v", got, tt.out)
			}
		})
	}
}
