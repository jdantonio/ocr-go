package lcd

import "testing"

func TestLcdToDigitWithGoodValues(t *testing.T) {
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

func TestLcdToDigitWithBadValues(t *testing.T) {
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

func TestLcdToNumber(t *testing.T) {
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

func TestDisplayInteger(t *testing.T) {
	number := Display{
		"   ",
		"  |",
		"  |",
	}
	expected := 1
	actual, _ := number.Integer()
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	number = Display{
		"    _  _     _  _  _  _  _  _ ",
		"  | _| _||_||_ |_   ||_||_|| |",
		"  ||_  _|  | _||_|  ||_|__||_|",
	}
	expected = 1234567890
	actual, _ = number.Integer()
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
