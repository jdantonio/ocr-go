/*
Package lcd provides utilities for converting LCD digits to Go data types.

LCD Digit Layout

LCD digits consist of a series of pipe, underscore, and space characters
arranged in a nine-character grid. The grid is three characters wide by three
characters high (which means that each digit spans three lines when read from a
file). Numbered from 0-8, the grid is alligned like this:

  0 1 2
  3 4 5
  6 7 8

Any element which does not contain a pipe or an underscore must contain a
space character, even when the space would be the last character in a line.
Trimming any trailing whitespace that is logically a part of the grid
renders the digit invalid.

The first row represents the top row of the digit. It may only contain
underscores and only in the center. The second row represents the top half
of the digit as well as the center "bar" of the digit, when present. It may
contain both pipes and underscores. The first and last characters may only
be pipes and the center character may only be an underscore. The bottom row
of each digit represents the bottom half of each digit as well as the lower
"bar" of the digit, when present. It has the same format as the middle row.
Following this pattern, the number 4 will be represented with the following
string:

  "   |_|  |"

  With indexes:

   0    1    2    3    4    5    6    7    8
  ' ', ' ', ' ', '|', '_', '|', ' ', ' ', '|'

When rendered, the digits 0-9 appear like this:

   _       _    _         _    _    _    _    _
  | |  |   _|   _|  |_|  |_   |_     |  |_|  |_|
  |_|  |  |_    _|    |   _|  |_|    |  |_|   _|

LCD File Format

When LCD digits are read from or written to a file the format is as follows:

* A single LCD number may contain multiple LCD digits.

* An LCD number may begin with one or more zero digits.

* Every LCD digit must contain exactly nine characters.

* Removing trailing whitespace that is part of a digit is an error.

* Only one LCD digit may exist on a single logical "line" (3 lines in the file).

* LCD digits must not be separated by additional whitespace.

* A blank like must appear after each LCD sequence, including the last one in the file.

* Which means that the number of lines in the file must be evenly divisible by 4
(three lines for the digits followed by one blank line).

Following these rules, an LCD input file consisting of three numbers each with
nine digits would look like this:

   _  _  _  _  _  _  _  _  _
  |_|  | _||_||_|| || |  | _|
  |_|  | _| _||_||_||_|  ||_

   _  _  _  _  _  _        _
  |_  _|  ||_   ||_||_|  ||_
  |_||_   | _|  ||_|  |  | _|

      _  _  _     _  _  _  _
    | _||_ |_   ||_ |_ |_||_|
    ||_ |_| _|  ||_| _| _||_|

*/
package lcd

import (
	"bytes"
	"errors"
)

//allDigits = [
//  # 0    1    2    3    4    5    6    7    8
//  # 1    2    3    4    5    6    7    8    9
//  #' '  '_'  ' '  '|'  '_'  '|'  '|'  '_'  '|'
//  ---------------------------------------------
//  [' ', '_', ' ', '|', ' ', '|', '|', '_', '|'], # 0
//  [' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', '|'], # 1
//  [' ', '_', ' ', ' ', '_', '|', '|', '_', ' '], # 2
//  [' ', '_', ' ', ' ', '_', '|', ' ', '_', '|'], # 3
//  [' ', ' ', ' ', '|', '_', '|', ' ', ' ', '|'], # 4
//  [' ', '_', ' ', '|', '_', ' ', ' ', '_', '|'], # 5
//  [' ', '_', ' ', '|', '_', ' ', '|', '_', '|'], # 6
//  [' ', '_', ' ', ' ', ' ', '|', ' ', ' ', '|'], # 7
//  [' ', '_', ' ', '|', '_', '|', '|', '_', '|'], # 8
//  [' ', '_', ' ', '|', '_', ' ', '|', '_', '|'], # 9
//]

// Number of characters necessary to describe a single digit on an LCD display.
const digitLength = 9

// Return value for integer conversion when an error is raised.
const Invalid = -1

// A single LCD digit.
type digit string

// A series of LCD digits representing a number.
type Number []digit

// A series of LCD digits ordered for display or file output.
type display []string

type DisplayScanner interface {
	Scan() bool
	Text() string
}

// String representations of all digits, 0-9, as read from an LCD display.
var allDigits = []digit{
	" _ | ||_|", // 0
	"     |  |", // 1
	" _  _||_ ", // 2
	" _  _| _|", // 3
	"   |_|  |", // 4
	" _ |_  _|", // 5
	" _ |_ |_|", // 6
	" _   |  |", // 7
	" _ |_||_|", // 8
	" _ |_| _|", // 9
}

func newDisplay() display {
	return make([]string, 3)
}

func ScanNext(scanner DisplayScanner) (int, error) {
	buffer := newDisplay()

	for i := 0; i < 3; i++ {
		buffer[i] = scanner.Text()
		if !scanner.Scan() {
			return Invalid, errors.New("insufficient lines")
		}
	}

	return buffer.integer()
}

// integer converts a series of LCD digits in display order into the
// corresponding integer value. The layout of each digit is explained in "LCD"
// Digit Layout" and "LCD File Format" above. The input is an array of exactly
// three strings where each string represents one "row" of a series of digits.
// Each individual digit is arranged across all three strings as it would be if
// the data were read from a file (sans the final blank line). Each string in
// the input array must be the same length as the other two as each represents
// different portions of the same digits.
func (dsply display) integer() (int, error) {
	if !dsply.isValid() {
		return Invalid, errors.New("invalid LCD display")
	}

	number := dsply.toDigits()

	return number.integer()
}

// integer converts a series of LCD digits into the integer value represented
// by those digits. The input is an array of strings where each string
// represents a single digit. The format of each string in the array is
// described in the digit.integer function.
func (number Number) integer() (int, error) {
	value := 0
	multiplier := 1

	for i := len(number) - 1; i >= 0; i-- {
		integer, err := digit(number[i]).integer()
		if err != nil {
			return Invalid, err
		} else {
			value += integer * multiplier
			multiplier *= 10
		}
	}

	return value, nil
}

// integer converts a string of characters which represents a single LCD digit
// into the corresponding integer value. Each character in the string
// represents a single element in the LCD grid described in "LCD Digit Layout"
// above. The index of each character within the string corresponds to the grid
// position in the 0-8 format. Thus, the input string must be exactly nine
// characters.
func (dgt digit) integer() (int, error) {
	if len(dgt) < digitLength {
		return Invalid, errors.New("too short")
	} else if len(dgt) > digitLength {
		return Invalid, errors.New("too long")
	}

	for integer, str := range allDigits {
		if dgt == str {
			return integer, nil
		}
	}

	return Invalid, errors.New("invalid characters")
}

// isValid checks a multi-digit LCD array, as described in display.integer, for
// a valid structure. It does not check any of the characters within the
// contained strings. Validating individual digits is left to the conversion
// functions.
func (dsply display) isValid() bool {
	return len(dsply) == 3 &&
		len(dsply[0]) == len(dsply[1]) &&
		len(dsply[1]) == len(dsply[2])
}

// toDigits converts an array of strings which represent LCD digits in "display
// order" (as described in display.integer) into an array of strings in
// conversion order (as described in Number.integer).
func (dsply display) toDigits() Number {
	digits := len(dsply[0]) / 3
	buffer := make([]bytes.Buffer, digits)

	// reorder
	for row := 0; row < 3; row++ {
		for digit := 0; digit < digits; digit++ {
			buffer[digit].WriteString(dsply[row][digit*3 : digit*3+3])
		}
	}

	// stringify
	number := make(Number, digits)
	for i := 0; i < digits; i++ {
		number[i] = digit(buffer[i].String())
	}

	return number
}
