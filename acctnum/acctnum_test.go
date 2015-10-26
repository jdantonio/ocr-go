package acctnum

import "testing"

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
	high := maximum + 10
	low := minimum - 10

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

	if IsValid(high) {
		t.Errorf("expected %v to not be valid", high)
	}

	if IsValid(low) {
		t.Errorf("expected %v to not be valid", low)
	}
}
