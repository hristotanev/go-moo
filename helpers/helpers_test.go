package helpers

import "testing"

func TestItReturnsTheMinOfTwoNumbers(t *testing.T) {
	if Min(1, 2) != 1 {
		t.Error("The returned min of two numbers is wrong.")
	}
}

func TestGenerateNumberOfLength(t *testing.T) {
	number := GenerateNumber(4)
	if len(number) != 4 {
		t.Errorf("Expected the length of the generated number to be 4, got %d", len(number))
	}
}
