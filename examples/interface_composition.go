package examples

import (
	"fmt"
	"time"

	"go-interface-enum-explorer/utils"
)

// Interfaces for the InterfaceComposition example
type CompReader interface {
	Read(p []byte) (n int, err error)
}

type CompWriter interface {
	Write(p []byte) (n int, err error)
}

type CompCloser interface {
	Close() error
}

// Composed interfaces
type CompReadWriter interface {
	CompReader
	CompWriter
}

type CompReadWriteCloser interface {
	CompReader
	CompWriter
	CompCloser
}

// Logger interface hierarchy
type CompLogger interface {
	Log(message string)
}

type CompTimestampLogger interface {
	CompLogger
	SetTimeFormat(format string)
}

// Implementation for composition examples
type CompFileHandler struct {
	filename   string
	timeFormat string
	isOpen     bool
}

func (f *CompFileHandler) Read(p []byte) (n int, err error) {
	if !f.isOpen {
		return 0, fmt.Errorf("file is not open")
	}
	fmt.Println("Reading from", f.filename)
	// Simulate reading data
	message := "Hello from file"
	n = copy(p, []byte(message))
	return n, nil
}

func (f *CompFileHandler) Write(p []byte) (n int, err error) {
	if !f.isOpen {
		return 0, fmt.Errorf("file is not open")
	}
	fmt.Printf("Writing to %s: %s\n", f.filename, string(p))
	return len(p), nil
}

func (f *CompFileHandler) Close() error {
	fmt.Println("Closing", f.filename)
	f.isOpen = false
	return nil
}

func (f *CompFileHandler) Open() error {
	fmt.Println("Opening", f.filename)
	f.isOpen = true
	return nil
}

func (f *CompFileHandler) Log(message string) {
	timestamp := time.Now().Format(f.timeFormat)
	fmt.Printf("[%s] %s\n", timestamp, message)
}

func (f *CompFileHandler) SetTimeFormat(format string) {
	f.timeFormat = format
}

// InterfaceComposition demonstrates how interfaces can be composed of other interfaces
func InterfaceComposition() {
	utils.PrintExplanation(`
INTERFACE COMPOSITION IN GO
=========================

Go supports interface composition - building larger interfaces by combining smaller ones.
This promotes the "interface segregation principle" - clients shouldn't depend on
methods they don't use. Small, focused interfaces are more reusable and composable.

Key points:
- Interfaces can embed other interfaces to create a new interface
- Types must implement all methods from all embedded interfaces
- Composition enables the creation of specific interfaces tailored to needs
- Small interfaces are more reusable and maintainable
- The standard library uses this pattern extensively (e.g., io.ReadWriter)
`)

	utils.PrintCode(`
// Small, focused interfaces
type Reader interface {
        Read(p []byte) (n int, err error)
}

type Writer interface {
        Write(p []byte) (n int, err error)
}

type Closer interface {
        Close() error
}

// Composed interfaces
type ReadWriter interface {
        Reader
        Writer
}

type ReadWriteCloser interface {
        Reader
        Writer
        Closer
}

// Logger interface hierarchy
type Logger interface {
        Log(message string)
}

type TimestampLogger interface {
        Logger
        SetTimeFormat(format string)
}

// Implementation of the interfaces
type FileHandler struct {
        filename   string
        timeFormat string
        isOpen     bool
}

func (f *FileHandler) Read(p []byte) (n int, err error) {
        if !f.isOpen {
                return 0, fmt.Errorf("file is not open")
        }
        fmt.Println("Reading from", f.filename)
        // Simulate reading data
        message := "Hello from file"
        n = copy(p, []byte(message))
        return n, nil
}

func (f *FileHandler) Write(p []byte) (n int, err error) {
        if !f.isOpen {
                return 0, fmt.Errorf("file is not open")
        }
        fmt.Printf("Writing to %s: %s\n", f.filename, string(p))
        return len(p), nil
}

func (f *FileHandler) Close() error {
        fmt.Println("Closing", f.filename)
        f.isOpen = false
        return nil
}

func (f *FileHandler) Open() error {
        fmt.Println("Opening", f.filename)
        f.isOpen = true
        return nil
}

func (f *FileHandler) Log(message string) {
        timestamp := time.Now().Format(f.timeFormat)
        fmt.Printf("[%s] %s\n", timestamp, message)
}

func (f *FileHandler) SetTimeFormat(format string) {
        f.timeFormat = format
}

func main() {
        // Create a FileHandler
        handler := &FileHandler{
                filename:   "data.txt",
                timeFormat: time.RFC3339,
        }
        handler.Open()
        
        // Use it as a ReadWriter
        var rw ReadWriter = handler
        writeData := []byte("Sample data")
        rw.Write(writeData)
        
        readData := make([]byte, 100)
        n, _ := rw.Read(readData)
        fmt.Printf("Read %d bytes: %s\n", n, string(readData[:n]))
        
        // Use it as a ReadWriteCloser
        var rwc ReadWriteCloser = handler
        rwc.Close()
        
        // Use it as a TimestampLogger
        var logger TimestampLogger = handler
        handler.Open() // Re-open for demonstration
        logger.Log("This is a log message")
        logger.SetTimeFormat("2006-01-02 15:04:05")
        logger.Log("This is a log message with new format")
}
`)

	// Actual implementation
	utils.PrintOutput("Running the code...")

	// Create a FileHandler
	handler := &CompFileHandler{
		filename:   "data.txt",
		timeFormat: time.RFC3339,
	}
	handler.Open()

	// Use it as a ReadWriter
	var rw CompReadWriter = handler
	writeData := []byte("Sample data")
	rw.Write(writeData)

	readData := make([]byte, 100)
	n, _ := rw.Read(readData)
	fmt.Printf("Read %d bytes: %s\n", n, string(readData[:n]))

	// Use it as a ReadWriteCloser
	var rwc CompReadWriteCloser = handler
	rwc.Close()

	// Use it as a TimestampLogger
	var logger CompTimestampLogger = handler
	handler.Open() // Re-open for demonstration
	logger.Log("This is a log message")
	logger.SetTimeFormat("2006-01-02 15:04:05")
	logger.Log("This is a log message with new format")

	utils.PrintKey(`
KEY TAKEAWAYS:
- Interface composition allows building larger interfaces from smaller ones
- A single type can implement multiple interfaces
- Clients can depend only on the interfaces with methods they need
- The standard library uses this pattern (e.g., io.ReadWriter, io.ReadWriteCloser)
- Small, focused interfaces lead to more flexible and reusable code
- Interface composition encourages the interface segregation principle
`)
}
