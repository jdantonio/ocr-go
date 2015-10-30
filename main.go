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
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/jdantonio/ocr-go/acctnum"
	"github.com/jdantonio/ocr-go/lcd"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func scan(scanner lcd.DisplayScanner, writer io.Writer) {
	var output string

	for scanner.Scan() {
		account, err := lcd.ScanNext(scanner)
		if err != nil {
			output = fmt.Sprintf("%09s ILL", err.Error())
		} else if acctnum.IsValid(account) {
			output = fmt.Sprintf("%09d", account)
		} else {
			output = fmt.Sprintf("%09d ERR", account)
		}
		writer.Write([]byte(output))
		writer.Write([]byte("\n"))
	}
}

func main() {

	infile, err := os.Open("data.txt")
	check(err)
	defer infile.Close()

	var outfilename string
	flag.StringVar(&outfilename, "outfile", "", "name of the output file")
	flag.Parse()

	var writer io.Writer

	if outfilename == "" {
		writer = os.Stdout
	} else {
		outfile, err := os.Create(outfilename)
		check(err)
		defer outfile.Close()
		writer = outfile
	}

	scanner := bufio.NewScanner(infile)
	scan(scanner, writer)
}
