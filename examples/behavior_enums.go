package examples

import (
	"fmt"

	"go-interface-enum-explorer/utils"
)

// Define PaymentMethod for the BehaviorEnums example
type BehPaymentMethod int

const (
	BehCreditCard BehPaymentMethod = iota
	BehPayPal
	BehBankTransfer
	BehCryptocurrency
)

func (p BehPaymentMethod) String() string {
	names := []string{
		"Credit Card",
		"PayPal",
		"Bank Transfer",
		"Cryptocurrency",
	}

	if int(p) < len(names) {
		return names[int(p)]
	}
	return fmt.Sprintf("Unknown Payment Method(%d)", p)
}

func (p BehPaymentMethod) ProcessingFee() float64 {
	switch p {
	case BehCreditCard:
		return 0.029 // 2.9%
	case BehPayPal:
		return 0.039 // 3.9%
	case BehBankTransfer:
		return 0.015 // 1.5%
	case BehCryptocurrency:
		return 0.01 // 1.0%
	default:
		return 0.05 // Default 5% for unknown methods
	}
}

func (p BehPaymentMethod) ProcessingTime() int {
	switch p {
	case BehCreditCard, BehPayPal:
		return 0 // Instant
	case BehBankTransfer:
		return 72 // 3 days
	case BehCryptocurrency:
		return 1 // 1 hour for confirmations
	default:
		return 24 // Default 24 hours
	}
}

func (p BehPaymentMethod) IsInstant() bool {
	return p.ProcessingTime() == 0
}

// Define Season for the BehaviorEnums example
type BehSeason int

const (
	BehSpring BehSeason = iota
	BehSummer
	BehAutumn
	BehWinter
)

func (s BehSeason) String() string {
	names := []string{"Spring", "Summer", "Autumn", "Winter"}
	if int(s) < len(names) {
		return names[int(s)]
	}
	return fmt.Sprintf("Unknown Season(%d)", s)
}

func (s BehSeason) MonthsInNorthernHemisphere() []string {
	switch s {
	case BehSpring:
		return []string{"March", "April", "May"}
	case BehSummer:
		return []string{"June", "July", "August"}
	case BehAutumn:
		return []string{"September", "October", "November"}
	case BehWinter:
		return []string{"December", "January", "February"}
	default:
		return []string{}
	}
}

func (s BehSeason) AverageTemperature(region string) string {
	switch region {
	case "Northern Europe":
		switch s {
		case BehSpring:
			return "5°C to 15°C"
		case BehSummer:
			return "15°C to 25°C"
		case BehAutumn:
			return "5°C to 15°C"
		case BehWinter:
			return "-5°C to 5°C"
		}
	case "Mediterranean":
		switch s {
		case BehSpring:
			return "15°C to 20°C"
		case BehSummer:
			return "25°C to 35°C"
		case BehAutumn:
			return "15°C to 25°C"
		case BehWinter:
			return "5°C to 15°C"
		}
	}
	return "Temperature data not available"
}

// BehaviorEnums demonstrates adding behavior directly to enum values
func BehaviorEnums() {
	utils.PrintExplanation(`
BEHAVIOR ENUMS IN GO
==================

Go's type system allows adding behavior to enum values through methods.
This creates rich enums that not only represent a value but also encapsulate
related functionality. This pattern can lead to more maintainable and expressive code.

Key points:
- Add methods to enum types to associate behavior with values
- Allows for polymorphic behavior without interfaces
- Makes code more self-contained and maintainable
- Can combine with the String() method pattern
- Particularly useful for domain-specific logic
`)

	utils.PrintCode(`
// PaymentMethod represents different ways to pay for an order
type PaymentMethod int

// Define the payment methods
const (
        CreditCard PaymentMethod = iota
        PayPal
        BankTransfer
        Cryptocurrency
)

// String returns the string representation of a PaymentMethod
func (p PaymentMethod) String() string {
        names := []string{
                "Credit Card",
                "PayPal",
                "Bank Transfer",
                "Cryptocurrency",
        }
        
        if int(p) < len(names) {
                return names[int(p)]
        }
        return fmt.Sprintf("Unknown Payment Method(%d)", p)
}

// ProcessingFee returns the fee percentage for this payment method
func (p PaymentMethod) ProcessingFee() float64 {
        switch p {
        case CreditCard:
                return 0.029 // 2.9%
        case PayPal:
                return 0.039 // 3.9%
        case BankTransfer:
                return 0.015 // 1.5%
        case Cryptocurrency:
                return 0.01 // 1.0%
        default:
                return 0.05 // Default 5% for unknown methods
        }
}

// ProcessingTime returns the typical processing time in hours
func (p PaymentMethod) ProcessingTime() int {
        switch p {
        case CreditCard, PayPal:
                return 0 // Instant
        case BankTransfer:
                return 72 // 3 days
        case Cryptocurrency:
                return 1 // 1 hour for confirmations
        default:
                return 24 // Default 24 hours
        }
}

// IsInstant returns whether the payment is processed instantly
func (p PaymentMethod) IsInstant() bool {
        return p.ProcessingTime() == 0
}

// Season represents seasons of the year
type Season int

const (
        Spring Season = iota
        Summer
        Autumn
        Winter
)

// String returns the name of the season
func (s Season) String() string {
        names := []string{"Spring", "Summer", "Autumn", "Winter"}
        if int(s) < len(names) {
                return names[int(s)]
        }
        return fmt.Sprintf("Unknown Season(%d)", s)
}

// MonthsInNorthernHemisphere returns the months this season occurs in the Northern Hemisphere
func (s Season) MonthsInNorthernHemisphere() []string {
        switch s {
        case Spring:
                return []string{"March", "April", "May"}
        case Summer:
                return []string{"June", "July", "August"}
        case Autumn:
                return []string{"September", "October", "November"}
        case Winter:
                return []string{"December", "January", "February"}
        default:
                return []string{}
        }
}

// AverageTemperature returns a general temperature range (°C) for the season
func (s Season) AverageTemperature(region string) string {
        switch region {
        case "Northern Europe":
                switch s {
                case Spring:
                        return "5°C to 15°C"
                case Summer:
                        return "15°C to 25°C"
                case Autumn:
                        return "5°C to 15°C"
                case Winter:
                        return "-5°C to 5°C"
                }
        case "Mediterranean":
                switch s {
                case Spring:
                        return "15°C to 20°C"
                case Summer:
                        return "25°C to 35°C"
                case Autumn:
                        return "15°C to 25°C"
                case Winter:
                        return "5°C to 15°C"
                }
        }
        return "Temperature data not available"
}

func main() {
        // Using PaymentMethod enums with behavior
        fmt.Println("Payment Method Examples:")
        
        paymentMethods := []PaymentMethod{
                CreditCard,
                PayPal,
                BankTransfer,
                Cryptocurrency,
        }
        
        // Calculating payment for a $100 order with different methods
        orderAmount := 100.0
        
        fmt.Println("Payment options for a $100 order:")
        for _, method := range paymentMethods {
                fee := method.ProcessingFee() * orderAmount
                total := orderAmount + fee
                
                fmt.Printf("- %s:\n", method)
                fmt.Printf("  Processing fee: $%.2f (%.1f%%)\n", fee, method.ProcessingFee()*100)
                fmt.Printf("  Total amount: $%.2f\n", total)
                fmt.Printf("  Processing time: %d hours (Instant: %v)\n", 
                        method.ProcessingTime(), method.IsInstant())
        }
        
        // Using Season enums with behavior
        fmt.Println("\nSeason Examples:")
        
        for season := Spring; season <= Winter; season++ {
                fmt.Printf("Season: %s\n", season)
                fmt.Printf("  Months (Northern Hemisphere): %v\n", 
                        season.MonthsInNorthernHemisphere())
                fmt.Printf("  Temperature in Northern Europe: %s\n", 
                        season.AverageTemperature("Northern Europe"))
                fmt.Printf("  Temperature in Mediterranean: %s\n", 
                        season.AverageTemperature("Mediterranean"))
        }
}
`)

	// Actual implementation
	utils.PrintOutput("Running the code...")

	// Using PaymentMethod enums with behavior
	fmt.Println("Payment Method Examples:")

	paymentMethods := []BehPaymentMethod{
		BehCreditCard,
		BehPayPal,
		BehBankTransfer,
		BehCryptocurrency,
	}

	// Calculating payment for a $100 order with different methods
	orderAmount := 100.0

	fmt.Println("Payment options for a $100 order:")
	for _, method := range paymentMethods {
		fee := method.ProcessingFee() * orderAmount
		total := orderAmount + fee

		fmt.Printf("- %s:\n", method)
		fmt.Printf("  Processing fee: $%.2f (%.1f%%)\n", fee, method.ProcessingFee()*100)
		fmt.Printf("  Total amount: $%.2f\n", total)
		fmt.Printf("  Processing time: %d hours (Instant: %v)\n",
			method.ProcessingTime(), method.IsInstant())
	}

	// Using Season enums with behavior
	fmt.Println("\nSeason Examples:")

	for season := BehSpring; season <= BehWinter; season++ {
		fmt.Printf("Season: %s\n", season)
		fmt.Printf("  Months (Northern Hemisphere): %v\n",
			season.MonthsInNorthernHemisphere())
		fmt.Printf("  Temperature in Northern Europe: %s\n",
			season.AverageTemperature("Northern Europe"))
		fmt.Printf("  Temperature in Mediterranean: %s\n",
			season.AverageTemperature("Mediterranean"))
	}

	utils.PrintKey(`
KEY TAKEAWAYS:
- Adding methods to enum types creates "rich enums" with behavior
- Behavior related to an enum value is encapsulated with the value itself
- This approach is more object-oriented while still being idiomatic Go
- Benefits include:
  - Code organization - related functionality stays with the type
  - Self-contained logic - less scattered switch statements
  - More expressive domain model
  - Easier maintenance
- Can be combined with String() methods and other enum patterns
- Great for domain-specific behavior that varies by enum value
`)
}
