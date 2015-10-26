package account

import "testing"

func TestAcctNumToDigits(t *testing.T) {
	acctnum := 345882865
	expected := digits{3, 4, 5, 8, 8, 2, 8, 6, 5}
	actual := acctnumToDigits(acctnum)

	ok := true

	for i := 0; i < acctnumLength; i++ {
		if actual[i] != expected[i] {
			ok = false
			break
		}
	}

	if !ok {
		t.Error()
	}
}

func TestChecksum(t *testing.T) {
	good := 345882865
	expected := 231
	if actual := checksum(good); actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestIsValid(t *testing.T) {
	good := 345882865
	bad := 111111111
	short := 34588286
	long := 3458828650

	if !IsValid(good) {
		t.Errorf("expected %v to be valid", good)
	}

	if IsValid(bad) {
		t.Errorf("expected %v to not be valid", bad)
	}

	if IsValid(short) {
		t.Errorf("expected %v to not be valid", short)
	}

	if IsValid(long) {
		t.Errorf("expected %v to not be valid", long)
	}
}
