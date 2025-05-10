package utils

import (
	"fmt"
)

// ANSI color codes
const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
)

// PrintColoredTitle prints a title in the specified color
func PrintColoredTitle(title string, color string) {
	fmt.Println(color + "===================================" + ColorReset)
	fmt.Println(color + title + ColorReset)
	fmt.Println(color + "===================================" + ColorReset)
}

// PrintExplanation prints an explanation block
func PrintExplanation(text string) {
	fmt.Println(ColorBlue + "--- EXPLANATION ---" + ColorReset)
	fmt.Println(text)
	fmt.Println()
}

// PrintCode prints a code example
func PrintCode(code string) {
	fmt.Println(ColorGreen + "--- CODE EXAMPLE ---" + ColorReset)
	fmt.Println(code)
	fmt.Println()
}

// PrintOutput prints the output of running the code
func PrintOutput(text string) {
	fmt.Println(ColorYellow + "--- OUTPUT ---" + ColorReset)
	fmt.Println(text)
	fmt.Println()
}

// PrintKey prints key takeaways
func PrintKey(text string) {
	fmt.Println(ColorMagenta + "--- KEY TAKEAWAYS ---" + ColorReset)
	fmt.Println(text)
}
