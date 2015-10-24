package main

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
