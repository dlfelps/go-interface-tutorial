package examples

import (
	"fmt"

	"go-interface-enum-explorer/utils"
)

// BasicEnums demonstrates the most basic way to implement enum-like constants in Go
func BasicEnums() {
	utils.PrintExplanation(`
BASIC ENUMS IN GO
===============

Go doesn't have a built-in enum type like many other languages, but it provides
ways to create enum-like constructs using constants and custom types.

The simplest approach is to use a group of constants with an integer type.
This provides basic enumeration functionality but lacks type safety.

Key points:
- Use const blocks to group related constants
- Constants can be untyped or have explicit types
- Basic enums are simple but lack type safety
- Comments help document valid values
`)

	utils.PrintCode(`
// UserRole represents permission levels in the system
const (
        // Guest can only view public content
        Guest = 0
        // User can access their own data and create content
        User = 1
        // Moderator can edit and delete content from others
        Moderator = 2
        // Admin has full access to all features
        Admin = 3
)

// Order status values
const (
        StatusPending = 1
        StatusPaid = 2
        StatusShipped = 3
        StatusDelivered = 4
        StatusCancelled = 5
)

// Function that uses basic enum values
func CheckAccess(role int, feature string) bool {
        switch feature {
        case "view_content":
                return role >= Guest
        case "create_content":
                return role >= User
        case "moderate_content":
                return role >= Moderator
        case "system_settings":
                return role >= Admin
        default:
                return false
        }
}

// Function that works with order status
func DescribeStatus(status int) string {
        switch status {
        case StatusPending:
                return "Order is pending payment"
        case StatusPaid:
                return "Payment received, preparing shipment"
        case StatusShipped:
                return "Order has been shipped"
        case StatusDelivered:
                return "Order has been delivered"
        case StatusCancelled:
                return "Order was cancelled"
        default:
                return "Unknown status"
        }
}

func main() {
        // Using the role enum
        userRole := User
        fmt.Printf("User with role %d trying to access features:\n", userRole)
        
        features := []string{"view_content", "create_content", "moderate_content", "system_settings"}
        for _, feature := range features {
                if CheckAccess(userRole, feature) {
                        fmt.Printf("- Can access %s\n", feature)
                } else {
                        fmt.Printf("- Cannot access %s\n", feature)
                }
        }
        
        // Using the order status enum
        fmt.Println("\nOrder status descriptions:")
        statuses := []int{StatusPending, StatusPaid, StatusShipped, StatusDelivered, StatusCancelled}
        for _, status := range statuses {
                desc := DescribeStatus(status)
                fmt.Printf("Status %d: %s\n", status, desc)
        }
        
        // Showing issues with basic enums
        fmt.Println("\nPotential issues with basic enums:")
        
        // This is valid but doesn't make sense semantically
        invalidRole := 999
        fmt.Printf("Invalid role %d can access create_content: %v\n", 
                invalidRole, CheckAccess(invalidRole, "create_content"))
        
        // Mixing different enum types is possible (but shouldn't be)
        mixedValue := StatusPaid // This is 2, same as Moderator
        fmt.Printf("Mixing enum types - StatusPaid as role: %v\n", 
                CheckAccess(mixedValue, "moderate_content"))
}
`)

	// Actual implementation
	// User role constants
	const (
		Guest     = 0
		User      = 1
		Moderator = 2
		Admin     = 3
	)

	// Order status constants
	const (
		StatusPending   = 1
		StatusPaid      = 2
		StatusShipped   = 3
		StatusDelivered = 4
		StatusCancelled = 5
	)

	// Function that uses role enum values
	checkAccess := func(role int, feature string) bool {
		switch feature {
		case "view_content":
			return role >= Guest
		case "create_content":
			return role >= User
		case "moderate_content":
			return role >= Moderator
		case "system_settings":
			return role >= Admin
		default:
			return false
		}
	}

	// Function that works with order status
	describeStatus := func(status int) string {
		switch status {
		case StatusPending:
			return "Order is pending payment"
		case StatusPaid:
			return "Payment received, preparing shipment"
		case StatusShipped:
			return "Order has been shipped"
		case StatusDelivered:
			return "Order has been delivered"
		case StatusCancelled:
			return "Order was cancelled"
		default:
			return "Unknown status"
		}
	}

	// Now actually use the code
	utils.PrintOutput("Running the code...")

	// Using the role enum
	userRole := User
	fmt.Printf("User with role %d trying to access features:\n", userRole)

	features := []string{"view_content", "create_content", "moderate_content", "system_settings"}
	for _, feature := range features {
		if checkAccess(userRole, feature) {
			fmt.Printf("- Can access %s\n", feature)
		} else {
			fmt.Printf("- Cannot access %s\n", feature)
		}
	}

	// Using the order status enum
	fmt.Println("\nOrder status descriptions:")
	statuses := []int{StatusPending, StatusPaid, StatusShipped, StatusDelivered, StatusCancelled}
	for _, status := range statuses {
		desc := describeStatus(status)
		fmt.Printf("Status %d: %s\n", status, desc)
	}

	// Showing issues with basic enums
	fmt.Println("\nPotential issues with basic enums:")

	// This is valid but doesn't make sense semantically
	invalidRole := 999
	fmt.Printf("Invalid role %d can access create_content: %v\n",
		invalidRole, checkAccess(invalidRole, "create_content"))

	// Mixing different enum types is possible (but shouldn't be)
	mixedValue := StatusPaid // This is 2, same as Moderator
	fmt.Printf("Mixing enum types - StatusPaid as role: %v\n",
		checkAccess(mixedValue, "moderate_content"))

	utils.PrintKey(`
KEY TAKEAWAYS:
- Go doesn't have built-in enums, but we can use constants to create enum-like behavior
- Basic enum implementation using untyped constants is simple but lacks type safety
- Constants provide compile-time checks but don't prevent invalid values at runtime
- Potential issues with basic enums:
  - Can pass any integer value, even invalid ones
  - Can mix different enum types if they share the same underlying type
  - No automatic string conversion for debugging/logging
- More robust enum implementations address these limitations
`)
}
