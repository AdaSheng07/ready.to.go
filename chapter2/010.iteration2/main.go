package main

import "fmt"

func main() {
	// guess a number from 1 to 1000
	guessNumber(0, 1000)
}
func guessNumber(left, right uint) {
	var guessed uint
	guessed = (left + right) / 2
	// give out the guess and tell it if it is too large/too small/correct
	var input string
	fmt.Printf("Is it %v?\n"+
		"If it is too large, input h; "+
		"if it is too small, input l; "+
		"if it is correct, input anything you want.\n", guessed)
	_, _ = fmt.Scanln(&input)
	// if left = right, guessed must be the answer!
	if ifLeftEqualsRight(left, right) {
		fmt.Printf("Your guess must be %v!\n", guessed)
		return
	}
	// if guessed is too large, guess another number from left to (guessed-1)
	if input == "h" {
		guessNumber(left, guessed-1)
	} else if input == "l" {
		// if guessed is too small, guess another number from (guessed+1) to right
		guessNumber(guessed+1, right)
	} else {
		fmt.Printf("Bingo! Your guess is %v!\n", guessed)
		return
	}
}
func ifLeftEqualsRight(a, b uint) bool {
	if a == b {
		return true
	} else {
		return false
	}
}
