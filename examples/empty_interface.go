package examples

import (
        "fmt"
        "reflect"

        "go-interface-enum-explorer/utils"
)

// EmptyInterface demonstrates the use and application of empty interfaces in Go
func EmptyInterface() {
        utils.PrintExplanation(`
EMPTY INTERFACE IN GO
===================

The empty interface (interface{}) has no methods, so all types satisfy it implicitly.
This means it can hold values of any type, making it useful for functions that need
to handle values of unknown types, like fmt.Println() or when you need generic containers.

In Go 1.18+, generics provide an alternative to empty interfaces for many use cases.

Key points:
- Empty interface can hold values of any type
- Written as interface{} or any in Go 1.18+
- Useful for functions that need to handle unknown types
- To use the value, you need type assertions or reflection
- Use with care - it bypasses Go's static type checking
`)

        utils.PrintCode(`
// PrintAny can take any type of value
func PrintAny(v interface{}) {
        fmt.Printf("Value: %v, Type: %T\n", v, v)
}

// Stack is a simple generic stack using an empty interface
type Stack struct {
        items []interface{}
}

// Push adds an item to the stack
func (s *Stack) Push(item interface{}) {
        s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() (interface{}, bool) {
        if len(s.items) == 0 {
                return nil, false
        }
        
        index := len(s.items) - 1
        item := s.items[index]
        s.items = s.items[:index]
        return item, true
}

func main() {
        // PrintAny can take any type
        PrintAny(42)
        PrintAny("hello")
        PrintAny(true)
        PrintAny(3.14)
        PrintAny([]int{1, 2, 3})
        
        // Using a stack with different types
        stack := Stack{}
        
        // Push different types to the stack
        stack.Push(42)
        stack.Push("Go")
        stack.Push(true)
        
        // Pop items and check their types
        for i := 0; i < 3; i++ {
                if item, ok := stack.Pop(); ok {
                        fmt.Printf("Popped %v of type %T\n", item, item)
                }
        }
}
`)

        // Actual implementation
        // Function that accepts any type
        printAny := func(v interface{}) {
                fmt.Printf("Value: %v, Type: %T\n", v, v)
        }

        // Stack implementation using empty interface
        type Stack struct {
                items []interface{}
        }

        // Push method
        push := func(s *Stack, item interface{}) {
                s.items = append(s.items, item)
        }

        // Pop method
        pop := func(s *Stack) (interface{}, bool) {
                if len(s.items) == 0 {
                        return nil, false
                }
                
                index := len(s.items) - 1
                item := s.items[index]
                s.items = s.items[:index]
                return item, true
        }

        // Function that uses reflection with interface{}
        describeValue := func(v interface{}) string {
                val := reflect.ValueOf(v)
                
                switch val.Kind() {
                case reflect.String:
                        return fmt.Sprintf("String with %d characters", val.Len())
                case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
                        return fmt.Sprintf("Integer with value %d", val.Int())
                case reflect.Float32, reflect.Float64:
                        return fmt.Sprintf("Float with value %f", val.Float())
                case reflect.Bool:
                        return fmt.Sprintf("Boolean set to %t", val.Bool())
                case reflect.Slice, reflect.Array:
                        return fmt.Sprintf("Sequence with %d elements", val.Len())
                default:
                        return "Unknown type"
                }
        }

        // Now actually use the code
        utils.PrintOutput("Running the code...")
        
        // Demonstrate PrintAny with different types
        fmt.Println("Using PrintAny with different types:")
        printAny(42)
        printAny("hello")
        printAny(true)
        printAny(3.14)
        printAny([]int{1, 2, 3})
        
        // Using a stack with different types
        fmt.Println("\nUsing a generic Stack:")
        stack := Stack{}
        
        // Push different types to the stack
        push(&stack, 42)
        push(&stack, "Go")
        push(&stack, true)
        
        // Pop items and check their types
        for i := 0; i < 3; i++ {
                if item, ok := pop(&stack); ok {
                        fmt.Printf("Popped %v of type %T\n", item, item)
                }
        }
        
        // Demonstrate reflection with interface{}
        fmt.Println("\nUsing reflection with interface{}:")
        fmt.Println(describeValue("Hello, Go!"))
        fmt.Println(describeValue(42))
        fmt.Println(describeValue(3.14159))
        fmt.Println(describeValue(true))
        fmt.Println(describeValue([]string{"a", "b", "c"}))

        utils.PrintKey(`
KEY TAKEAWAYS:
- The empty interface (interface{}) can hold values of any type
- It's useful for generic functions like PrintAny or data structures like Stack
- To work with values stored in an empty interface, you need type assertions or reflection
- In Go 1.18+, generics offer a more type-safe alternative for many use cases
- Empty interfaces sacrifice compile-time type checking for flexibility
- Common uses: fmt package functions, containers/collections, plugins, and configuration
`)
}
