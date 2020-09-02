package utils

import "testing"

func TestAdd(t *testing.T) {
	expected := 3
	actual := Add(1, 2)
	if actual != expected {
		t.Errorf("Total was incorrect! Expected: %d, Actual: %d", expected, actual)
	}
}
