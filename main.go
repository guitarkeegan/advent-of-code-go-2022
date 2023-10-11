package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counter := 0
	// open the file
	f, err := os.Open("rock-paper-scissors")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	// close after program ends
	defer f.Close()

	// create new scanner to read file
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		line := scanner.Text()

		// run my functions
		result := getWinner(&line)
		choice := getChoice(&line)

		counter += result
		counter += choice

	}
	fmt.Println(counter)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getWinner(s *string) int {
	elfChoice := (*s)[0:1]
	myChoice := (*s)[2:3]

	if elfChoice == "A" && myChoice == "Z" ||
		elfChoice == "B" && myChoice == "X" ||
		elfChoice == "C" && myChoice == "Y" {
		return 0
	} else if elfChoice == "A" && myChoice == "Y" ||
		elfChoice == "B" && myChoice == "Z" ||
		elfChoice == "C" && myChoice == "X" {
		return 6
	} else {
		return 3
	}
}

func getChoice(s *string) int {
	choice := (*s)[2:]
	if choice == "X" {
		return 1
	} else if choice == "Y" {
		return 2
	} else {
		return 3
	}
}
