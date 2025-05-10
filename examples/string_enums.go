package examples

import (
	"fmt"

	"go-interface-enum-explorer/utils"
)

// Define Direction for the StringEnums example
type StrDirection int

const (
	StrNorth StrDirection = iota
	StrEast
	StrSouth
	StrWest
)

func (d StrDirection) String() string {
	switch d {
	case StrNorth:
		return "North"
	case StrEast:
		return "East"
	case StrSouth:
		return "South"
	case StrWest:
		return "West"
	default:
		return fmt.Sprintf("Unknown Direction(%d)", d)
	}
}

// Define HttpStatus for the StringEnums example
type StrHttpStatus int

const (
	StrStatusOK           StrHttpStatus = 200
	StrStatusCreated      StrHttpStatus = 201
	StrStatusBadRequest   StrHttpStatus = 400
	StrStatusUnauthorized StrHttpStatus = 401
	StrStatusForbidden    StrHttpStatus = 403
	StrStatusNotFound     StrHttpStatus = 404
	StrStatusServerError  StrHttpStatus = 500
)

func (s StrHttpStatus) String() string {
	statusText := map[StrHttpStatus]string{
		StrStatusOK:           "200 OK",
		StrStatusCreated:      "201 Created",
		StrStatusBadRequest:   "400 Bad Request",
		StrStatusUnauthorized: "401 Unauthorized",
		StrStatusForbidden:    "403 Forbidden",
		StrStatusNotFound:     "404 Not Found",
		StrStatusServerError:  "500 Internal Server Error",
	}

	if text, exists := statusText[s]; exists {
		return text
	}
	return fmt.Sprintf("Unknown Status(%d)", s)
}

func (s StrHttpStatus) IsSuccess() bool {
	return s >= 200 && s < 300
}

func (s StrHttpStatus) IsError() bool {
	return s >= 400
}

func (s StrHttpStatus) IsClientError() bool {
	return s >= 400 && s < 500
}

func (s StrHttpStatus) IsServerError() bool {
	return s >= 500
}

// StringEnums demonstrates adding string representation to enum values
func StringEnums() {
	utils.PrintExplanation(`
STRING ENUMS IN GO
================

A common enhancement to enums is adding string representation for debugging,
serialization, and user display. Go's type system enables this through custom String()
methods and defined types.

Key points:
- Define a custom type for the enum
- Implement the String() method for automatic string conversion
- Use a map or switch statement to convert enum values to strings
- Values remain type-safe but gain string representation
- Useful for logging, debugging, and user interfaces
`)

	utils.PrintCode(`
// Direction represents a cardinal direction
type Direction int

// Define the direction constants
const (
        North Direction = iota
        East
        South
        West
)

// String returns the string representation of a Direction
// This implements the fmt.Stringer interface
func (d Direction) String() string {
        switch d {
        case North:
                return "North"
        case East:
                return "East"
        case South:
                return "South"
        case West:
                return "West"
        default:
                return fmt.Sprintf("Unknown Direction(%d)", d)
        }
}

// HttpStatus represents HTTP response status codes
type HttpStatus int

// Define some common HTTP status codes
const (
        StatusOK           HttpStatus = 200
        StatusCreated      HttpStatus = 201
        StatusBadRequest   HttpStatus = 400
        StatusUnauthorized HttpStatus = 401
        StatusForbidden    HttpStatus = 403
        StatusNotFound     HttpStatus = 404
        StatusServerError  HttpStatus = 500
)

// String returns the string representation of an HttpStatus
func (s HttpStatus) String() string {
        // Using a map for status code to message mapping
        statusText := map[HttpStatus]string{
                StatusOK:           "200 OK",
                StatusCreated:      "201 Created",
                StatusBadRequest:   "400 Bad Request",
                StatusUnauthorized: "401 Unauthorized",
                StatusForbidden:    "403 Forbidden",
                StatusNotFound:     "404 Not Found",
                StatusServerError:  "500 Internal Server Error",
        }
        
        if text, exists := statusText[s]; exists {
                return text
        }
        return fmt.Sprintf("Unknown Status(%d)", s)
}

// IsSuccess returns true if this status code represents a successful response
func (s HttpStatus) IsSuccess() bool {
        return s >= 200 && s < 300
}

// IsError returns true if this status code represents an error response
func (s HttpStatus) IsError() bool {
        return s >= 400
}

// IsClientError returns true if this status code represents a client error
func (s HttpStatus) IsClientError() bool {
        return s >= 400 && s < 500
}

// IsServerError returns true if this status code represents a server error
func (s HttpStatus) IsServerError() bool {
        return s >= 500
}

func main() {
        // Using Direction enum with string representation
        fmt.Println("Direction examples:")
        
        // Creating and displaying enum values
        heading := North
        fmt.Printf("Current heading: %v\n", heading)
        
        // The String() method is automatically called when printing
        fmt.Printf("Turning right to %v\n", East)
        fmt.Printf("Continuing to %v\n", South)
        fmt.Printf("Finally turning to %v\n", West)
        
        // Using HttpStatus with string representation and behavior
        fmt.Println("\nHTTP Status examples:")
        
        // Simulating some responses
        responses := []HttpStatus{
                StatusOK,
                StatusCreated,
                StatusBadRequest,
                StatusNotFound,
                StatusServerError,
        }
        
        for _, status := range responses {
                // The String() method is called automatically in formatting
                fmt.Printf("Response status: %v\n", status)
                fmt.Printf("  Is success: %v\n", status.IsSuccess())
                fmt.Printf("  Is client error: %v\n", status.IsClientError())
                fmt.Printf("  Is server error: %v\n", status.IsServerError())
        }
}
`)

	// Actual implementation
	utils.PrintOutput("Running the code...")

	// Using Direction enum with string representation
	fmt.Println("Direction examples:")

	// Creating and displaying enum values
	heading := StrNorth
	fmt.Printf("Current heading: %v\n", heading)

	// The String() method is automatically called when printing
	fmt.Printf("Turning right to %v\n", StrEast)
	fmt.Printf("Continuing to %v\n", StrSouth)
	fmt.Printf("Finally turning to %v\n", StrWest)

	// Using HttpStatus with string representation and behavior
	fmt.Println("\nHTTP Status examples:")

	// Simulating some responses
	responses := []StrHttpStatus{
		StrStatusOK,
		StrStatusCreated,
		StrStatusBadRequest,
		StrStatusNotFound,
		StrStatusServerError,
	}

	for _, status := range responses {
		// The String() method is called automatically in formatting
		fmt.Printf("Response status: %v\n", status)
		fmt.Printf("  Is success: %v\n", status.IsSuccess())
		fmt.Printf("  Is client error: %v\n", status.IsClientError())
		fmt.Printf("  Is server error: %v\n", status.IsServerError())
	}

	utils.PrintKey(`
KEY TAKEAWAYS:
- Adding String() methods implements the fmt.Stringer interface for automatic string conversion
- String representations make enums more readable in logs and debugging
- Enum types can have methods for additional behavior (e.g., IsSuccess(), IsError())
- Implementation options include switch statements or maps
- Default case should handle unexpected values
- String enums maintain type safety while adding human-readable representation
- Pattern works well with JSON encoding/decoding when using MarshalJSON/UnmarshalJSON
`)
}
