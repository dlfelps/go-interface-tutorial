package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PressEnterToContinue pauses execution until the user presses Enter
func PressEnterToContinue() {
	fmt.Print("\nPress Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// GetUserInput prompts the user and returns their input as a string
func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// GetUserChoice prompts the user to select from a list of options
func GetUserChoice(prompt string, options []string) (int, error) {
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	input := GetUserInput(prompt)

	// Try to convert the input to an integer
	var choice int
	_, err := fmt.Sscanf(input, "%d", &choice)
	if err != nil {
		return 0, err
	}

	// Check if the choice is valid
	if choice < 1 || choice > len(options) {
		return 0, fmt.Errorf("invalid choice: %d", choice)
	}

	return choice, nil
}
