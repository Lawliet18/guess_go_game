package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	min, max := 1, 100
	victoryVal := rand.Intn(max)

	fmt.Println("Enter exit to end the game!!")
	fmt.Printf("Guess a number (%d-%d): ", min, max)

	// Read input line by line
	for scanner.Scan() {
		text := scanner.Text() // Get the current line of text
		if text == "exit" {
			fmt.Printf("Victory value is %d", victoryVal)
			break // Exit loop if an empty line is entered
		}
		val, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			fmt.Println("Incorrect input!!")
			fmt.Printf("Guess a number (%d-%d): ", min, max)
			continue
		}

		if val > int64(max) || val < int64(min) {
			fmt.Println("Not in range")
			continue
		}

		if val == int64(victoryVal) {
			fmt.Println("Victory!! Your number is ", victoryVal)
			break
		} else if val > int64(victoryVal) {
			max = int(val)
			fmt.Printf("Guess a number (%d-%d): ", min, max)
		} else {
			min = int(val)
			fmt.Printf("Guess a number (%d-%d): ", min, max)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
