package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
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
		fmt.Println(scanner.Text())
	}

	if lines%4 != 0 {
		fmt.Println("Boom!")
	}
	check(scanner.Err())
}
