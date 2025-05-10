package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "runtime"
        "strconv"
        "strings"

        "go-interface-enum-explorer/examples"
        "go-interface-enum-explorer/utils"
)

// Tutorial topics organized by category and complexity
var topics = map[string][]string{
        "interfaces": {
                "Basic Interfaces",
                "Interface Implementation",
                "Empty Interface",
                "Type Assertion",
                "Interface Composition",
        },
        "enums": {
                "Basic Enums",
                "Iota Enums",
                "String Enums",
                "Behavior Enums",
        },
}

// Example functions mapped to their names
var exampleFunctions = map[string]func(){
        "Basic Interfaces":        examples.BasicInterfaces,
        "Interface Implementation": examples.InterfaceImplementation,
        "Empty Interface":         examples.EmptyInterface,
        "Type Assertion":          examples.TypeAssertion,
        "Interface Composition":   examples.InterfaceComposition,
        "Basic Enums":             examples.BasicEnums,
        "Iota Enums":              examples.IotaEnums,
        "String Enums":            examples.StringEnums,
        "Behavior Enums":          examples.BehaviorEnums,
}

func main() {
        clearScreen()
        displayWelcome()

        scanner := bufio.NewScanner(os.Stdin)
        
        for {
                displayMainMenu()
                fmt.Print("\nEnter your choice (or 'q' to quit): ")
                scanner.Scan()
                choice := strings.TrimSpace(scanner.Text())

                if choice == "q" || choice == "Q" {
                        fmt.Println("Thank you for learning Go interfaces and enums. Happy coding!")
                        break
                }

                switch choice {
                case "1":
                        tutorialMode(scanner)
                case "2":
                        browseExamples(scanner)
                case "3":
                        displayHelp()
                        utils.PressEnterToContinue()
                default:
                        fmt.Println("Invalid choice. Please try again.")
                        utils.PressEnterToContinue()
                }
        }
}

func displayWelcome() {
        utils.PrintColoredTitle("Welcome to Go Interface & Enum Explorer", utils.ColorCyan)
        fmt.Println("This interactive tool will help you learn about interfaces and enums in Go.")
        fmt.Println("You'll see examples ranging from basic concepts to advanced usage patterns.")
        fmt.Println("\nEach example includes:")
        fmt.Println("- Explanation of the concept")
        fmt.Println("- Sample code with comments")
        fmt.Println("- Output of the code execution")
        fmt.Println("\nLet's begin exploring Go's powerful interface system and enum patterns!")
        utils.PressEnterToContinue()
}

func displayMainMenu() {
        clearScreen()
        utils.PrintColoredTitle("Main Menu", utils.ColorGreen)
        fmt.Println("1. Start Tutorial (guided journey)")
        fmt.Println("2. Browse Examples (pick specific topics)")
        fmt.Println("3. Help")
        fmt.Println("q. Quit")
}

func tutorialMode(scanner *bufio.Scanner) {
        var allTopics []string
        
        // Combine all topics in the right order for a progressive learning experience
        allTopics = append(allTopics, topics["interfaces"]...)
        allTopics = append(allTopics, topics["enums"]...)
        
        for i, topic := range allTopics {
                clearScreen()
                utils.PrintColoredTitle(fmt.Sprintf("Tutorial (%d/%d): %s", i+1, len(allTopics), topic), utils.ColorYellow)
                
                // Run the example for this topic
                if fn, exists := exampleFunctions[topic]; exists {
                        fn()
                } else {
                        fmt.Printf("Example for %s is not implemented yet.\n", topic)
                }
                
                // After showing an example, offer navigation options
                if i < len(allTopics)-1 {
                        fmt.Println("\nOptions:")
                        fmt.Println("n - Next example")
                        fmt.Println("m - Return to main menu")
                        fmt.Print("\nYour choice: ")
                        
                        scanner.Scan()
                        choice := strings.TrimSpace(scanner.Text())
                        
                        if choice == "m" || choice == "M" {
                                break
                        }
                        // Any other input will move to the next example
                } else {
                        fmt.Println("\nCongratulations! You've completed all the tutorials.")
                        utils.PressEnterToContinue()
                }
        }
}

func browseExamples(scanner *bufio.Scanner) {
        for {
                clearScreen()
                utils.PrintColoredTitle("Browse Examples", utils.ColorBlue)
                
                fmt.Println("Categories:")
                fmt.Println("1. Interfaces")
                fmt.Println("2. Enums")
                fmt.Println("b. Back to Main Menu")
                
                fmt.Print("\nSelect a category (or 'b' to go back): ")
                scanner.Scan()
                category := strings.TrimSpace(scanner.Text())
                
                if category == "b" || category == "B" {
                        break
                }
                
                var topicList []string
                var categoryTitle string
                
                switch category {
                case "1":
                        topicList = topics["interfaces"]
                        categoryTitle = "Interfaces"
                case "2":
                        topicList = topics["enums"]
                        categoryTitle = "Enums"
                default:
                        fmt.Println("Invalid category. Please try again.")
                        utils.PressEnterToContinue()
                        continue
                }
                
                for {
                        clearScreen()
                        utils.PrintColoredTitle(fmt.Sprintf("%s Examples", categoryTitle), utils.ColorBlue)
                        
                        for i, topic := range topicList {
                                fmt.Printf("%d. %s\n", i+1, topic)
                        }
                        fmt.Println("b. Back to Categories")
                        
                        fmt.Print("\nSelect an example (or 'b' to go back): ")
                        scanner.Scan()
                        topicChoice := strings.TrimSpace(scanner.Text())
                        
                        if topicChoice == "b" || topicChoice == "B" {
                                break
                        }
                        
                        topicIndex, err := strconv.Atoi(topicChoice)
                        if err != nil || topicIndex < 1 || topicIndex > len(topicList) {
                                fmt.Println("Invalid selection. Please try again.")
                                utils.PressEnterToContinue()
                                continue
                        }
                        
                        selectedTopic := topicList[topicIndex-1]
                        clearScreen()
                        utils.PrintColoredTitle(selectedTopic, utils.ColorYellow)
                        
                        if fn, exists := exampleFunctions[selectedTopic]; exists {
                                fn()
                        } else {
                                fmt.Printf("Example for %s is not implemented yet.\n", selectedTopic)
                        }
                        
                        utils.PressEnterToContinue()
                }
        }
}

func displayHelp() {
        clearScreen()
        utils.PrintColoredTitle("Help", utils.ColorMagenta)
        
        fmt.Println("How to use this tool:")
        fmt.Println("1. Tutorial Mode: Guides you through all examples in a logical order.")
        fmt.Println("2. Browse Examples: Pick specific topics you're interested in.")
        
        fmt.Println("\nAbout Go Interfaces:")
        fmt.Println("- Interfaces in Go define behavior, not structure.")
        fmt.Println("- Types implement interfaces implicitly (no 'implements' keyword).")
        fmt.Println("- Interfaces can be composed of other interfaces.")
        fmt.Println("- The empty interface (interface{}) can hold values of any type.")
        
        fmt.Println("\nAbout Go Enums:")
        fmt.Println("- Go doesn't have built-in enums, but provides patterns to implement them.")
        fmt.Println("- The iota identifier is used to create incrementing constants.")
        fmt.Println("- Type safety can be achieved with custom types and constants.")
        fmt.Println("- String representations can be added with methods.")
        
        fmt.Println("\nTip: Running the examples and reviewing the code is the best way to learn!")
}

func clearScreen() {
        var cmd *exec.Cmd
        if runtime.GOOS == "windows" {
                cmd = exec.Command("cmd", "/c", "cls")
        } else {
                cmd = exec.Command("clear")
        }
        cmd.Stdout = os.Stdout
        _ = cmd.Run()
}
