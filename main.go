// The Bank OCR kata from Coding Dojo.
//
// The problem description from the original source:
//
//   You work for a bank, which has recently purchased an ingenious machine to
//   assist in reading letters and faxes sent in by branch offices. The machine
//   scans the paper documents, and produces a file with a number of entries
//   which each look like this:
//
//       _  _     _  _  _  _  _
//     | _| _||_||_ |_   ||_||_|
//     ||_  _|  | _||_|  ||_| _|
//
//   Each entry is 4 lines long, and each line has 27 characters. The first 3
//   lines of each entry contain an account number written using pipes and
//   underscores, and the fourth line is blank. Each account number should have
//   9 digits, all of which should be in the range 0-9. A normal file contains
//   around 500 entries.
//
// More information, including the complete description of all user stories,
// can be found at http://www.codingdojo.org/cgi-bin/index.pl?KataBankOCR.
package main

import (
	"bufio"
	"fmt"
	"github.com/jdantonio/ocr-go/acctnum"
	"github.com/jdantonio/ocr-go/lcd"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {

	file, err := os.Open("data.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		account, err := lcd.ScanNext(scanner)
		check(err)
		if acctnum.IsValid(account) {
			fmt.Println(account)
		} else {
			fmt.Printf("INVALID (%v)\n", account)
		}
	}
}
