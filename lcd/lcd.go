package lcd

import (
	"errors"
)

//LCD = [
//  # 0    1    2    3    4    5    6    7    8
//  # 1    2    3    4    5    6    7    8    9
//  #' '  '_'  ' '  '|'  '_'  '|'  '|'  '_'  '|'
//  [' ', '_', ' ', '|', ' ', '|', '|', '_', '|'], # 0
//  [' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', '|'], # 1
//  [' ', '_', ' ', ' ', '_', '|', '|', '_', ' '], # 2
//  [' ', '_', ' ', ' ', '_', '|', ' ', '_', '|'], # 3
//  [' ', ' ', ' ', '|', '_', '|', ' ', ' ', '|'], # 4
//  [' ', '_', ' ', '|', '_', ' ', ' ', '_', '|'], # 5
//  [' ', '_', ' ', '|', '_', ' ', '|', '_', '|'], # 6
//  [' ', '_', ' ', ' ', ' ', '|', ' ', ' ', '|'], # 7
//  [' ', '_', ' ', '|', '_', '|', '|', '_', '|'], # 8
//  [' ', '_', ' ', '|', '_', '|', '_', '_', '|'], # 9
//]

const LcdDigitLength = 9

var LCD = []string{
	" _ | ||_|", // 0
	"     |  |", // 1
	" _  _||_ ", // 2
	" _  _| _|", // 3
	"   |_|  |", // 4
	" _ |_  _|", // 5
	" _ |_ |_|", // 6
	" _   |  |", // 7
	" _ |_||_|", // 8
	" _ |_|__|", // 9
}

func LcdToDigit(str string) (int, error) {
	const invalid = -1

	if len(str) < LcdDigitLength {
		return invalid, errors.New("too short")
	} else if len(str) > LcdDigitLength {
		return invalid, errors.New("too long")
	}

	for digit, lcd := range LCD {
		if str == lcd {
			return digit, nil
		}
	}

	return invalid, errors.New("invalid characters")
}

func LcdToInt(slice []string) (int, error) {
	const invalid = -1

	value := 0
	multiplier := 1

	for i := len(slice) - 1; i >= 0; i-- {
		digit, err := LcdToDigit(slice[i])
		if err != nil {
			return invalid, err
		} else {
			value += digit * multiplier
			multiplier *= 10
		}
	}

	return value, nil
}
