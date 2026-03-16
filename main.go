package main

import (
    "fmt"
    "log"
    "strings"
)

func generateMessages(message string, count int) ([]string, error) {
    if strings.TrimSpace(message) == "" {
        return nil, fmt.Errorf("message cannot be empty")
    }
    if count <= 0 {
        return nil, fmt.Errorf("count must be greater than zero")
    }

    msgs := make([]string, count)
    for i := 0; i < count; i++ {
        msgs[i] = fmt.Sprintf("%s #%d", message, i+1)
    }
    return msgs, nil
}

func main() {
    msgs, err := generateMessages("Hello World", 5)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    for _, msg := range msgs {
        fmt.Println(msg)
    }
}
