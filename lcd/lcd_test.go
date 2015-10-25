package lcd

import "testing"

const mockScannerValue = 123

type mockScanner struct {
	TotalLines, CurrentLine int
}

func (scanner mockScanner) Text() string {
	switch scanner.CurrentLine {
	case 0:
		return "    _  _ "
	case 1:
		return "  | _| _|"
	case 2:
		return "  ||_  _|"
	default:
		return ""
	}
}

func (scanner *mockScanner) Scan() bool {
	scanner.CurrentLine++
	return scanner.CurrentLine < 5
}

func (scanner *mockScanner) ShortCircuit() {
	scanner.CurrentLine = 2
}

func TestScanNextWithValidInput(t *testing.T) {
	actual, err := ScanNext(new(mockScanner))
	if err != nil {
		t.Errorf("%v should be valid but returned error %v", actual, err)
	}
	if actual != mockScannerValue {
		t.Errorf("expected %v but got %v", mockScannerValue, actual)
	}
}

func TestScanNextWithInvalidInput(t *testing.T) {
	scanner := new(mockScanner)
	scanner.ShortCircuit()
	actual, err := ScanNext(scanner)
	if err == nil {
		t.Error("should have returned an error but didn't")
	}
	if actual != Invalid {
		t.Errorf("expected invalid but got %v", actual)
	}
}

func TestDigitToIntegerWithGoodValues(t *testing.T) {
	for expected, digit := range allDigits {
		actual, err := digit.integer()
		if err != nil {
			t.Errorf("%v should be valid but returned error %v", actual, err)
		}
		if actual != expected {
			t.Errorf("%v does not equal expected value %v", actual, expected)
		}
	}
}

func TestDigitToIntegerWithBadValues(t *testing.T) {
	invalids := []digit{
		"",              // empty digit
		"/-/-/-/-/",     // invalid characters
		" _ ",           // too short
		" _ | ||_|    ", // too long
		"_________",     // does not match
	}

	for _, digit := range invalids {
		actual, err := digit.integer()
		if err == nil {
			t.Errorf("%v should be invalid but no error was returned", actual)
		}
	}
}

func TestNumberToInteger(t *testing.T) {
	number := Number{
		allDigits[1],
		allDigits[3],
		allDigits[5],
		allDigits[7],
		allDigits[9],
		allDigits[2],
		allDigits[4],
		allDigits[6],
		allDigits[8],
		allDigits[0],
	}

	expected := 1357924680
	actual, _ := number.integer()
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	invalid := Number{
		allDigits[1], // good
		"____",       // bad
		allDigits[2], // good
	}

	actual, err := invalid.integer()
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}
}

func TestDisplayToInteger(t *testing.T) {
	dsply := display{
		"   ",
		"  |",
		"  |",
	}
	expected := 1
	actual, _ := dsply.integer()
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	dsply = display{
		"    _  _     _  _  _  _  _  _ ",
		"  | _| _||_||_ |_   ||_||_|| |",
		"  ||_  _|  | _||_|  ||_| _||_|",
	}
	expected = 1234567890
	actual, _ = dsply.integer()
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	invalid := display{
		"       _     _  _  _  _  _  _ ",
		"  |  | _||_||_ |_   ||_||_|| |",
		"  ||   _|  | _||_|  ||_| _||_|",
	}
	actual, err := invalid.integer()
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}

	invalid = display{
		"    _  _     _  _  _  _  _  _ ",
		"  | _| _||_||_ |_   ||_||_|| |",
	}
	actual, err = invalid.integer()
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}

	invalid = display{
		"    _  _     _  _  _  _  _ ",
		"  | _| _||_||_ |_   ||_||_|| |",
		"  ||_  _|  | _||_|  ||_|__|",
	}
	actual, err = invalid.integer()
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}
}
