package main

import (
	"bufio"
	//"github.com/jdantonio/ocr-go/lcd"
	//"bytes"
	"errors"
	"fmt"
	"os"
)

const AccountNumberLength = 30 // for the test file; the real file will have 27

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {

	lines := 0

	file, err := os.Open("data.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines += 1
		//fmt.Println(scanner.Text())
		buffer := scanner.Text()

		length := len(buffer)
		if length != AccountNumberLength && length != 0 {
			check(errors.New("main: incorrect LCD string length"))
		}

		fmt.Println(buffer)
	}

	if lines%4 != 0 {
		fmt.Println("Boom!")
	}
	check(scanner.Err())
}
