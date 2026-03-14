// Checking Queue
package main

import (
    "fmt"
)

// printMessage prints the given message a specified number of times.
// It returns an error if the count is invalid.
func printMessage(message string, count int) error {
    if count <= 0 {
        return fmt.Errorf("count must be greater than zero")
    }

    for i := 1; i <= count; i++ {
        fmt.Printf("%s #%d\n", message, i)
    }
    return nil
}

func main() {
    // Example usage
    if err := printMessage("Hello World", 5); err != nil {
        fmt.Println("Error:", err)
    }
}
