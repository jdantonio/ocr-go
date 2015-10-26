package account

const acctnumLength = 9

const maxAcctnum = 999999999

type digits []int

func newDigits() digits {
	return make([]int, acctnumLength)
}

func IsValid(acctnum int) bool {
	if acctnum <= maxAcctnum && checksum(acctnum)%11 == 0 {
		return true
	} else {
		return false
	}
}

func checksum(acctnum int) int {
	digits := acctnumToDigits(acctnum)
	checksum := 0

	for i, v := range digits {
		checksum += v * (acctnumLength - i)
	}

	return checksum
}

func acctnumToDigits(acctnum int) digits {
	current := acctnum
	dgts := newDigits()

	for i := acctnumLength - 1; i >= 0; i-- {
		digit := current % 10
		dgts[i] = digit
		current = (current - digit) / 10
	}

	return dgts
}
