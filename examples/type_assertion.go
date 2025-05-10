package examples

import (
	"fmt"

	"go-interface-enum-explorer/utils"
)

// Define Animal interface for the TypeAssertion example
type AssertAnimal interface {
	Speak() string
}

// Define Dog for the TypeAssertion example
type AssertDog struct {
	Breed string
}

func (d AssertDog) Speak() string {
	return "Woof!"
}

// Define Cat for the TypeAssertion example
type AssertCat struct {
	Color string
}

func (c AssertCat) Speak() string {
	return "Meow!"
}

// Define Duck for the TypeAssertion example
type AssertDuck struct {
	Habitat string
}

func (d AssertDuck) Speak() string {
	return "Quack!"
}

// TypeAssertion demonstrates how to extract and use concrete types from interfaces
func TypeAssertion() {
	utils.PrintExplanation(`
TYPE ASSERTIONS AND TYPE SWITCHES IN GO
=====================================

When working with interfaces, you often need to access the concrete value or check the underlying type.
Go provides two mechanisms for this: type assertions and type switches.

Type Assertion: x.(T)
- Used to extract the concrete value of type T from an interface
- Panics if the interface doesn't hold a value of type T (unless using the "comma ok" form)

Type Switch: switch x.(type)
- Evaluates the type of an interface value
- Executes the case that matches the concrete type

Key points:
- Type assertions retrieve the concrete value from an interface
- The "comma ok" idiom safely checks if an assertion is valid
- Type switches handle multiple possible types in a concise way
- Both are essential for working with interfaces effectively
`)

	utils.PrintCode(`
// Simple hierarchy using interfaces
type Animal interface {
        Speak() string
}

type Dog struct {
        Breed string
}

func (d Dog) Speak() string {
        return "Woof!"
}

type Cat struct {
        Color string
}

func (c Cat) Speak() string {
        return "Meow!"
}

type Duck struct {
        Habitat string
}

func (d Duck) Speak() string {
        return "Quack!"
}

// Function that demonstrates type assertion with "comma ok" idiom
func AnimalDetail(a Animal) {
        // Type assertion with "comma ok" idiom
        if dog, ok := a.(Dog); ok {
                fmt.Printf("Dog of breed: %s\n", dog.Breed)
        } else if cat, ok := a.(Cat); ok {
                fmt.Printf("Cat of color: %s\n", cat.Color)
        } else {
                fmt.Println("Unknown animal type")
        }
}

// Function that demonstrates type switch
func DescribeAnimal(a Animal) {
        fmt.Printf("Animal says: %s\n", a.Speak())
        
        // Type switch
        switch v := a.(type) {
        case Dog:
                fmt.Printf("This is a %s dog\n", v.Breed)
        case Cat:
                fmt.Printf("This is a %s cat\n", v.Color)
        case Duck:
                fmt.Printf("This is a duck that lives in %s\n", v.Habitat)
        default:
                fmt.Println("This is an unknown animal type")
        }
}

func main() {
        // Create some animals
        animals := []Animal{
                Dog{Breed: "Labrador"},
                Cat{Color: "Black"},
                Duck{Habitat: "Pond"},
        }
        
        // Demonstrate type assertion
        fmt.Println("Using type assertion:")
        for _, animal := range animals {
                AnimalDetail(animal)
        }
        
        // Demonstrate type switch
        fmt.Println("\nUsing type switch:")
        for _, animal := range animals {
                DescribeAnimal(animal)
        }
        
        // Demonstrate unsafe type assertion (would panic if not for comma-ok)
        fmt.Println("\nSafe vs. unsafe type assertion:")
        var a Animal = Dog{Breed: "Poodle"}
        
        // Safe - using comma ok idiom
        if cat, ok := a.(Cat); ok {
                fmt.Printf("Cat color: %s\n", cat.Color)
        } else {
                fmt.Println("Not a cat")
        }
        
        // This would panic if we did a direct assertion without checking:
        // cat := a.(Cat) // would panic
}
`)

	// Actual implementation
	utils.PrintOutput("Running the code...")

	// Function for type assertion with comma-ok
	animalDetail := func(a AssertAnimal) {
		// Type assertion with "comma ok" idiom
		if dog, ok := a.(AssertDog); ok {
			fmt.Printf("Dog of breed: %s\n", dog.Breed)
		} else if cat, ok := a.(AssertCat); ok {
			fmt.Printf("Cat of color: %s\n", cat.Color)
		} else {
			fmt.Println("Unknown animal type")
		}
	}

	// Function for type switching
	describeAnimal := func(a AssertAnimal) {
		fmt.Printf("Animal says: %s\n", a.Speak())

		// Type switch
		switch v := a.(type) {
		case AssertDog:
			fmt.Printf("This is a %s dog\n", v.Breed)
		case AssertCat:
			fmt.Printf("This is a %s cat\n", v.Color)
		case AssertDuck:
			fmt.Printf("This is a duck that lives in %s\n", v.Habitat)
		default:
			fmt.Println("This is an unknown animal type")
		}
	}

	// Create some animals
	animals := []AssertAnimal{
		AssertDog{Breed: "Labrador"},
		AssertCat{Color: "Black"},
		AssertDuck{Habitat: "Pond"},
	}

	// Demonstrate type assertion
	fmt.Println("Using type assertion:")
	for _, animal := range animals {
		animalDetail(animal)
	}

	// Demonstrate type switch
	fmt.Println("\nUsing type switch:")
	for _, animal := range animals {
		describeAnimal(animal)
	}

	// Demonstrate safe type assertion
	fmt.Println("\nSafe vs. unsafe type assertion:")
	var a AssertAnimal = AssertDog{Breed: "Poodle"}

	// Safe - using comma ok idiom
	if cat, ok := a.(AssertCat); ok {
		fmt.Printf("Cat color: %s\n", cat.Color)
	} else {
		fmt.Println("Not a cat")
	}

	utils.PrintKey(`
KEY TAKEAWAYS:
- Type assertions (x.(T)) extract concrete types from interfaces
- The "comma ok" idiom (value, ok := x.(T)) provides safe type assertion
- Type switches are cleaner than multiple if-else assertions
- Type assertions are essential when working with empty interfaces
- Type switches are particularly useful when handling multiple possible types
- Without type assertions, interface values can only be used through their methods
`)
}
