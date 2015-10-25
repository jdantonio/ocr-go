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
	for expected, digit := range AllDigits {
		actual, err := digit.Integer()
		if err != nil {
			t.Errorf("%v should be valid but returned error %v", actual, err)
		}
		if actual != expected {
			t.Errorf("%v does not equal expected value %v", actual, expected)
		}
	}
}

func TestDigitToIntegerWithBadValues(t *testing.T) {
	invalids := []Digit{
		"",              // empty Digit
		"/-/-/-/-/",     // invalid characters
		" _ ",           // too short
		" _ | ||_|    ", // too long
		"_________",     // does not match
	}

	for _, digit := range invalids {
		actual, err := digit.Integer()
		if err == nil {
			t.Errorf("%v should be invalid but no error was returned", actual)
		}
	}
}

func TestNumberToInteger(t *testing.T) {
	number := Number{
		AllDigits[1],
		AllDigits[3],
		AllDigits[5],
		AllDigits[7],
		AllDigits[9],
		AllDigits[2],
		AllDigits[4],
		AllDigits[6],
		AllDigits[8],
		AllDigits[0],
	}

	expected := 1357924680
	actual, _ := number.Integer()
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	invalid := Number{
		AllDigits[1], // good
		"____",       // bad
		AllDigits[2], // good
	}

	actual, err := invalid.Integer()
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}
}

func TestDisplayToInteger(t *testing.T) {
	display := Display{
		"   ",
		"  |",
		"  |",
	}
	expected := 1
	actual, _ := display.Integer()
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	display = Display{
		"    _  _     _  _  _  _  _  _ ",
		"  | _| _||_||_ |_   ||_||_|| |",
		"  ||_  _|  | _||_|  ||_| _||_|",
	}
	expected = 1234567890
	actual, _ = display.Integer()
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	invalid := Display{
		"       _     _  _  _  _  _  _ ",
		"  |  | _||_||_ |_   ||_||_|| |",
		"  ||   _|  | _||_|  ||_| _||_|",
	}
	actual, err := invalid.Integer()
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}

	invalid = Display{
		"    _  _     _  _  _  _  _  _ ",
		"  | _| _||_||_ |_   ||_||_|| |",
	}
	actual, err = invalid.Integer()
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}

	invalid = Display{
		"    _  _     _  _  _  _  _ ",
		"  | _| _||_||_ |_   ||_||_|| |",
		"  ||_  _|  | _||_|  ||_|__|",
	}
	actual, err = invalid.Integer()
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}
}
