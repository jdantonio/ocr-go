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
		"____",       // bad
		allDigits[2], // good
		allDigits[3], // good
	}
	expected_error := "1??23"

	actual, err := invalid.integer()
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}
	if err.Error() != expected_error {
		t.Errorf("expected error to be %v but got %v", expected_error, err.Error())
	}
}

func TestAlternates(t *testing.T) {
	given := 1234567890
	expected := []int{
		7234567890,
		1934567890,
		1294567890,
		1234667890,
		1234967890,
		1234557890,
		1234587890,
		1234561890,
		1234567090,
		1234567690,
		1234567990,
		1234567820,
		1234567830,
		1234567850,
		1234567880,
		1234567898,
	}

	actual := Alternates(given)

	for _, exp := range expected {
		ok := false
		for _, alt := range actual {
			if exp == alt {
				ok = true
				break
			}
		}
		if !ok {
			t.Errorf("expected %v but was not found in list of alternates", exp)
			break
		}
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
