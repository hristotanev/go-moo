package helpers

import "testing"

func TestMin(t *testing.T) {
	if Min(1, 1) != 1 {
		t.Fail()
	}

	if Min(2, 1) != 1 {
		t.Fail()
	}
}

func TestGenerateNumberOfLength(t *testing.T) {
	number := GenerateNumber(4)
	if len(number) != 4 {
		t.Fail()
	}
}
