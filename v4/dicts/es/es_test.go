package es

import (
	"fmt"
	"testing"

	"github.com/digitalgopnik/golem/v4"
)

func TestSpanishUsage(t *testing.T) {
	l, err := golem.New(New())
	if err != nil {
		fmt.Println(err)
	}
	_ = l
	word := l.Lemma("Buenas")
	fmt.Println(word)
	result := "bueno"
	if word != result {
		t.Errorf("Wanted %s, got %s.", result, word)
	}
}
