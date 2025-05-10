package examples

import (
        "fmt"

        "go-interface-enum-explorer/utils"
)

// IotaEnums demonstrates using iota for creating sequential enum values
func IotaEnums() {
        utils.PrintExplanation(`
IOTA ENUMS IN GO
==============

Go provides the iota identifier for creating sequences of related constants.
When used in a const block, iota starts at 0 and increments by 1 for each constant.

Key points:
- iota starts at 0 and increments for each constant
- Can be used with expressions for more complex sequences
- Makes it easier to maintain sequential constants
- Often used with custom types for better type safety
- Constants remain at their assigned values even if reordered
`)

        utils.PrintCode(`
// Weekday represents days of the week
type Weekday int

// Define the days of the week using iota
const (
        Sunday Weekday = iota    // 0
        Monday                   // 1
        Tuesday                  // 2
        Wednesday                // 3
        Thursday                 // 4
        Friday                   // 5
        Saturday                 // 6
)

// LogLevel for controlling application logging
type LogLevel int

// Define log levels using iota with expressions
const (
        LogDebug LogLevel = iota * 10    // 0
        LogInfo                          // 10
        LogWarning                       // 20
        LogError                         // 30
        LogFatal                         // 40
)

// BitFlag demonstrates using iota for bit flags/masks
type BitFlag uint

// Define bit flags using shifts with iota
const (
        ReadPermission BitFlag = 1 << iota  // 1 (001)
        WritePermission                      // 2 (010)
        ExecutePermission                    // 4 (100)
)

// IsWeekend checks if a day is a weekend
func IsWeekend(day Weekday) bool {
        return day == Sunday || day == Saturday
}

// ShouldLog determines if a message at the given level should be logged
func ShouldLog(messageLevel, configuredLevel LogLevel) bool {
        return messageLevel >= configuredLevel
}

// CheckPermissions verifies if all required permissions are granted
func CheckPermissions(granted, required BitFlag) bool {
        return granted&required == required
}

func main() {
        // Using weekday enum
        today := Monday
        fmt.Printf("Today is %d. Is it a weekend? %v\n", today, IsWeekend(today))
        weekend := Saturday
        fmt.Printf("Saturday is %d. Is it a weekend? %v\n", weekend, IsWeekend(weekend))
        
        // Using log level enum
        systemLevel := LogInfo
        fmt.Println("\nLog level configured to:", systemLevel)
        fmt.Printf("Debug message (level %d) will be logged: %v\n", 
                LogDebug, ShouldLog(LogDebug, systemLevel))
        fmt.Printf("Info message (level %d) will be logged: %v\n", 
                LogInfo, ShouldLog(LogInfo, systemLevel))
        fmt.Printf("Error message (level %d) will be logged: %v\n", 
                LogError, ShouldLog(LogError, systemLevel))
        
        // Using bit flag enum
        userPermissions := ReadPermission | WritePermission  // 3 (011)
        fmt.Println("\nUser permissions:", userPermissions)
        
        fmt.Printf("Has read permission: %v\n", 
                CheckPermissions(userPermissions, ReadPermission))
        fmt.Printf("Has write permission: %v\n", 
                CheckPermissions(userPermissions, WritePermission))
        fmt.Printf("Has execute permission: %v\n", 
                CheckPermissions(userPermissions, ExecutePermission))
        fmt.Printf("Has read+write permissions: %v\n", 
                CheckPermissions(userPermissions, ReadPermission|WritePermission))
}
`)

        // Actual implementation
        // Weekday type and constants
        type Weekday int

        const (
                Sunday Weekday = iota
                Monday
                Tuesday
                Wednesday
                Thursday
                Friday
                Saturday
        )

        // LogLevel type and constants
        type LogLevel int

        const (
                LogDebug LogLevel = iota * 10
                LogInfo
                LogWarning
                LogError
                LogFatal
        )

        // BitFlag type and constants
        type BitFlag uint

        const (
                ReadPermission BitFlag = 1 << iota
                WritePermission
                ExecutePermission
        )

        // IsWeekend function
        isWeekend := func(day Weekday) bool {
                return day == Sunday || day == Saturday
        }

        // ShouldLog function
        shouldLog := func(messageLevel, configuredLevel LogLevel) bool {
                return messageLevel >= configuredLevel
        }

        // CheckPermissions function
        checkPermissions := func(granted, required BitFlag) bool {
                return granted&required == required
        }

        // Now actually use the code
        utils.PrintOutput("Running the code...")
        
        // Using weekday enum
        today := Monday
        fmt.Printf("Today is %d. Is it a weekend? %v\n", today, isWeekend(today))
        weekend := Saturday
        fmt.Printf("Saturday is %d. Is it a weekend? %v\n", weekend, isWeekend(weekend))
        
        // Using log level enum
        systemLevel := LogInfo
        fmt.Println("\nLog level configured to:", systemLevel)
        fmt.Printf("Debug message (level %d) will be logged: %v\n", 
                LogDebug, shouldLog(LogDebug, systemLevel))
        fmt.Printf("Info message (level %d) will be logged: %v\n", 
                LogInfo, shouldLog(LogInfo, systemLevel))
        fmt.Printf("Error message (level %d) will be logged: %v\n", 
                LogError, shouldLog(LogError, systemLevel))
        
        // Using bit flag enum
        userPermissions := ReadPermission | WritePermission
        fmt.Println("\nUser permissions:", userPermissions)
        
        fmt.Printf("Has read permission: %v\n", 
                checkPermissions(userPermissions, ReadPermission))
        fmt.Printf("Has write permission: %v\n", 
                checkPermissions(userPermissions, WritePermission))
        fmt.Printf("Has execute permission: %v\n", 
                checkPermissions(userPermissions, ExecutePermission))
        fmt.Printf("Has read+write permissions: %v\n", 
                checkPermissions(userPermissions, ReadPermission|WritePermission))

        utils.PrintKey(`
KEY TAKEAWAYS:
- iota provides an automatic way to create sequential constants
- Type safety is improved by using custom types for enums
- iota can be combined with expressions for more complex sequences
- Common patterns include:
  - Simple incrementing values (0, 1, 2, ...)
  - Multiplied values for levels or priorities (0, 10, 20, ...)
  - Bit shifts for flags and masks (1, 2, 4, 8, ...)
- iota resets to 0 in each const block
- Values can be skipped or customized as needed
`)
}
