/*
Package acctnum provides utilities for working with account numbers.
*/
package acctnum

import (
	"math"
)

// Number of digits in an account number.
const length = 9

// Maximum value for a valid account number.
var maximum = int(math.Pow10(length)) - 1

// Minimum value for a valid account number.
var minimum = int(math.Pow10(length - 1))

// IsValid assesses the validity of the given account number based on the
// length and checksum.
func IsValid(acctnum int) bool {
	if acctnum <= maximum &&
		acctnum >= minimum &&
		checksum(acctnum)%11 == 0 {
		return true
	} else {
		return false
	}
}

// checksum calculates the checksum of the given account number.
func checksum(acctnum int) int {
	current := acctnum
	checksum := 0

	for i := length - 1; i >= 0; i-- {
		digit := current % 10
		checksum += digit * (length - i)
		current = (current - digit) / 10
	}

	return checksum
}
