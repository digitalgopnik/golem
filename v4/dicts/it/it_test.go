package it

import (
	"testing"

	"github.com/digitalgopnik/golem/v4"
)

func TestItalianUsage(t *testing.T) {
	lem, err := golem.New(New())
	if err != nil {
		t.Fatal(err)
	}
	// abbarbicare	abbarbicai
	result := "abbarbicare"
	lemma := lem.Lemma("abbarbicai")
	if lemma != result {
		t.Errorf("Expected %v, but got %v", result, lemma)
	}

}
