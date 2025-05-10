package examples

import (
        "fmt"

        "go-interface-enum-explorer/utils"
)

// Define Shape interface for the BasicInterfaces example
type BasicShape interface {
        Area() float64
        Perimeter() float64
}

// Define Rectangle for the BasicInterfaces example
type BasicRectangle struct {
        Width  float64
        Height float64
}

func (r BasicRectangle) Area() float64 {
        return r.Width * r.Height
}

func (r BasicRectangle) Perimeter() float64 {
        return 2 * (r.Width + r.Height)
}

// Define Circle for the BasicInterfaces example
type BasicCircle struct {
        Radius float64
}

func (c BasicCircle) Area() float64 {
        return 3.14159 * c.Radius * c.Radius
}

func (c BasicCircle) Perimeter() float64 {
        return 2 * 3.14159 * c.Radius
}

// BasicInterfaces demonstrates the fundamental concepts of interfaces in Go
func BasicInterfaces() {
        utils.PrintExplanation(`
BASIC INTERFACES IN GO
=====================

An interface in Go is a collection of method signatures that a type can implement.
It defines behavior, not structure. Any type that implements all the methods
of an interface implicitly satisfies that interface.

Key points:
- Interfaces define behavior through method signatures
- Types implement interfaces implicitly (no "implements" keyword)
- A type can implement multiple interfaces
- Interfaces allow for polymorphism in Go
`)

        utils.PrintCode(`
// Shape is an interface that defines a common behavior for shapes
type Shape interface {
        Area() float64
        Perimeter() float64
}

// Rectangle is a concrete type that will implement the Shape interface
type Rectangle struct {
        Width  float64
        Height float64
}

// Area calculates the area of a Rectangle
// This method implements the Area method from the Shape interface
func (r Rectangle) Area() float64 {
        return r.Width * r.Height
}

// Perimeter calculates the perimeter of a Rectangle
// This method implements the Perimeter method from the Shape interface
func (r Rectangle) Perimeter() float64 {
        return 2 * (r.Width + r.Height)
}

// Circle is another concrete type that will implement the Shape interface
type Circle struct {
        Radius float64
}

// Area calculates the area of a Circle
func (c Circle) Area() float64 {
        return 3.14159 * c.Radius * c.Radius
}

// Perimeter calculates the perimeter of a Circle
func (c Circle) Perimeter() float64 {
        return 2 * 3.14159 * c.Radius
}

// PrintShapeInfo takes a Shape interface and prints information about it
func PrintShapeInfo(s Shape) {
        fmt.Printf("Area: %.2f\n", s.Area())
        fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
        // Create a Rectangle instance
        rect := Rectangle{Width: 5, Height: 4}
        
        // Create a Circle instance
        circle := Circle{Radius: 3}
        
        fmt.Println("Rectangle:")
        PrintShapeInfo(rect)
        
        fmt.Println("\nCircle:")
        PrintShapeInfo(circle)
}
`)

        // Actual implementation
        utils.PrintOutput("Running the code...")
        
        // Function that uses the interface
        printShapeInfo := func(s BasicShape) {
                fmt.Printf("Area: %.2f\n", s.Area())
                fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
        }
        
        // Create a Rectangle
        rect := BasicRectangle{Width: 5, Height: 4}
        
        // Create a Circle
        circle := BasicCircle{Radius: 3}
        
        fmt.Println("Rectangle:")
        printShapeInfo(rect)
        
        fmt.Println("\nCircle:")
        printShapeInfo(circle)

        utils.PrintKey(`
KEY TAKEAWAYS:
- Both Rectangle and Circle implement the Shape interface by providing Area() and Perimeter() methods
- No explicit declaration is needed to say a type implements an interface
- The PrintShapeInfo function can accept any type that satisfies the Shape interface
- This allows for polymorphic behavior - different types responding to the same method calls
`)
}
