package main

import (
    "fmt"
)

// printMessage prints the given message n times.
func printMessage(message string, n int) {
    for i := 1; i <= n; i++ {
        fmt.Printf("%s #%d\n", message, i)
    }
}

func main() {
    // Clear intent: print "Hello World" five times
    printMessage("Hello World", 5)
}
