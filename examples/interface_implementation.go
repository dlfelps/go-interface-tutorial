package examples

import (
        "fmt"
        "strings"

        "go-interface-enum-explorer/utils"
)

// Define Writer interface for the InterfaceImplementation example
type ImplWriter interface {
        Write(data string) (int, error)
}

// Define ConsoleWriter for the InterfaceImplementation example
type ImplConsoleWriter struct {
        Prefix string
}

func (cw ImplConsoleWriter) Write(data string) (int, error) {
        formatted := cw.Prefix + data
        fmt.Println(formatted)
        return len(formatted), nil
}

// Define FileLogger for the InterfaceImplementation example
type ImplFileLogger struct {
        FileName string
}

func (fl ImplFileLogger) Write(data string) (int, error) {
        fmt.Printf("[Writing to %s]: %s\n", fl.FileName, data)
        return len(data), nil
}

// Define UppercaseWriter for the InterfaceImplementation example
type ImplUppercaseWriter struct {
        ActualWriter ImplWriter
}

func (uw ImplUppercaseWriter) Write(data string) (int, error) {
        return uw.ActualWriter.Write(strings.ToUpper(data))
}

// InterfaceImplementation demonstrates how types implement interfaces in Go
func InterfaceImplementation() {
        utils.PrintExplanation(`
INTERFACE IMPLEMENTATION IN GO
=============================

In Go, a type implements an interface by implementing its methods.
There's no explicit declaration of intent like "implements" in other languages.
This is called "implicit implementation" and is a key feature of Go's design.

Key points:
- Any type that implements all methods of an interface automatically satisfies that interface
- Interface implementation is implicit (no "implements" keyword)
- Concrete types can implement many interfaces simultaneously
- Methods must match exactly (same name, parameters, and return types)
`)

        utils.PrintCode(`
// Writer is an interface that defines a Write method
type Writer interface {
        Write(data string) (int, error)
}

// ConsoleWriter writes data to the console
type ConsoleWriter struct {
        Prefix string
}

// Write method for ConsoleWriter - implements the Writer interface
func (cw ConsoleWriter) Write(data string) (int, error) {
        formatted := cw.Prefix + data
        fmt.Println(formatted)
        return len(formatted), nil
}

// FileLogger simulates logging to a file
type FileLogger struct {
        FileName string
}

// Write method for FileLogger - also implements the Writer interface
func (fl FileLogger) Write(data string) (int, error) {
        // In a real implementation, this would write to a file
        fmt.Printf("[Writing to %s]: %s\n", fl.FileName, data)
        return len(data), nil
}

// UppercaseWriter converts text to uppercase before writing
type UppercaseWriter struct {
        ActualWriter Writer // Composition with another Writer
}

// Write method for UppercaseWriter - also implements Writer interface
func (uw UppercaseWriter) Write(data string) (int, error) {
        // Convert to uppercase and delegate to the embedded Writer
        return uw.ActualWriter.Write(strings.ToUpper(data))
}

// WriteToSomewhere is a function that uses the Writer interface
func WriteToSomewhere(writer Writer, messages []string) {
        for _, msg := range messages {
                writer.Write(msg)
        }
}

func main() {
        // Create instances of different Writers
        console := ConsoleWriter{Prefix: "LOG: "}
        file := FileLogger{FileName: "app.log"}
        
        // Create a composed writer that converts to uppercase
        uppercaseConsole := UppercaseWriter{ActualWriter: console}
        
        messages := []string{"Hello, World!", "Learning Go interfaces", "Composition is powerful"}
        
        fmt.Println("Writing to console:")
        WriteToSomewhere(console, messages)
        
        fmt.Println("\nWriting to file:")
        WriteToSomewhere(file, messages)
        
        fmt.Println("\nWriting uppercase to console:")
        WriteToSomewhere(uppercaseConsole, messages)
}
`)

        // Actual implementation
        utils.PrintOutput("Running the code...")
        
        // Function that uses the interface
        writeToSomewhere := func(writer ImplWriter, messages []string) {
                for _, msg := range messages {
                        writer.Write(msg)
                }
        }
        
        // Create instances of different Writers
        console := ImplConsoleWriter{Prefix: "LOG: "}
        file := ImplFileLogger{FileName: "app.log"}
        
        // Create a composed writer that converts to uppercase
        uppercaseConsole := ImplUppercaseWriter{ActualWriter: console}
        
        messages := []string{"Hello, World!", "Learning Go interfaces", "Composition is powerful"}
        
        fmt.Println("Writing to console:")
        writeToSomewhere(console, messages)
        
        fmt.Println("\nWriting to file:")
        writeToSomewhere(file, messages)
        
        fmt.Println("\nWriting uppercase to console:")
        writeToSomewhere(uppercaseConsole, messages)

        utils.PrintKey(`
KEY TAKEAWAYS:
- Both ConsoleWriter and FileLogger implement the Writer interface by providing a Write method
- The UppercaseWriter shows interface composition by containing another Writer
- The WriteToSomewhere function works with any type that satisfies the Writer interface
- Different implementations of the same interface allow for different behaviors
- This demonstrates the "program to an interface, not an implementation" principle
`)
}
