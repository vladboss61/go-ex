package main

import (
    "fmt";
    "time"
)

func time_week() {
    defer fmt.Println("execute lately")

    fmt.Printf("Day of week: %v\n", time.Now().Weekday())
}