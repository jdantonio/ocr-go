package lcd

import "testing"

func TestLcdToDigitWithGoodValues(t *testing.T) {
	for expected, lcd := range LCD {
		actual, err := LcdToDigit(lcd)
		if err != nil {
			t.Errorf("%v should be valid but returned error %v", actual, err)
		}
		if actual != expected {
			t.Errorf("%v does not equal expected value %v", actual, expected)
		}
	}
}

func TestLcdToDigitWithBadValues(t *testing.T) {
	invalids := []string{
		"",              // empty string
		"/-/-/-/-/",     // invalid characters
		" _ ",           // too short
		" _ | ||_|    ", // too long
		"_________",     // does not match
	}

	for _, lcd := range invalids {
		actual, err := LcdToDigit(lcd)
		if err == nil {
			t.Errorf("%v should be invalid but no error was returned", actual)
		}
	}
}

func TestLcdToInt(t *testing.T) {
	number := []string{
		LCD[1],
		LCD[3],
		LCD[5],
		LCD[7],
		LCD[9],
		LCD[2],
		LCD[4],
		LCD[6],
		LCD[8],
		LCD[0],
	}

	expected := 1357924680
	actual, _ := LcdToInt(number)
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	invalid := []string{
		LCD[1], // good
		"____", // bad
		LCD[2], // good
	}

	actual, err := LcdToInt(invalid)
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}
}

func TestDisplayToInt(t *testing.T) {
	number := []string{
		"   ",
		"  |",
		"  |",
	}
	expected := 1
	actual, _ := DisplayToInt(number)
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	number = []string{
		"    _  _     _  _  _  _  _  _ ",
		"  | _| _||_||_ |_   ||_||_|| |",
		"  ||_  _|  | _||_|  ||_|__||_|",
	}
	expected = 1234567890
	actual, _ = DisplayToInt(number)
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	invalid := []string{
		"       _     _  _  _  _  _  _ ",
		"  |  | _||_||_ |_   ||_||_|| |",
		"  ||   _|  | _||_|  ||_| _||_|",
	}
	actual, err := DisplayToInt(invalid)
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}

	invalid = []string{
		"    _  _     _  _  _  _  _  _ ",
		"  | _| _||_||_ |_   ||_||_|| |",
	}
	actual, err = DisplayToInt(invalid)
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}

	invalid = []string{
		"    _  _     _  _  _  _  _ ",
		"  | _| _||_||_ |_   ||_||_|| |",
		"  ||_  _|  | _||_|  ||_|__|",
	}
	actual, err = DisplayToInt(invalid)
	if err == nil {
		t.Errorf("expected an error but got %v", actual)
	}
}
