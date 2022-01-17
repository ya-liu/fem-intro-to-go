package utils

import "testing"

func TestAdd(t *testing.T) {
	expected := 6
	actual := Add(1, 2, 3)
	if actual != expected {
		t.Errorf("Total was incorrect! Expected: %d, Actual: %d", expected, actual)
	}
}
