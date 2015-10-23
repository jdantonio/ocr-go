package main

import (
	"bufio"
	//"bytes"
	"errors"
	"fmt"
	"os"
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

const AccountNumberLength = 30 // for the test file' the real file will have 27

var lcd = [...]string{" _ | ||_|", "     |  |", " _  _||_ ", " _  _| _|", "   |_|  |", " _ |_  _|", " _ |_ |_|", " _   |  |", " _ |_||_|", " _ |_|__|"}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {

	//for _, digit := range lcd {
	//fmt.Printf("%v\n", digit)
	//for _, char := range digit {
	//fmt.Printf("%v", char)
	//fmt.Println("")
	//}
	//}

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
