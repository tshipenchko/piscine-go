package piscine

import "testing"

func TestBasicJoin(t *testing.T) {
	answer := "Hello! How are you?"
	elems := []string{"Hello!", " How", " are", " you?"}
	got := BasicJoin(elems)

	if got != answer {
		t.Error(
			"Got:", got, "\n",
			"Want:", answer,
		)
	}
}
