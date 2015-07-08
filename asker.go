package main

import (
    "fmt"
    "log"
)

func AskUserForInt(message string) int {
    var num int
    fmt.Print(message)
    _, err := fmt.Scanf("%d", &num)
    
    if err != nil {
        log.Fatalf("Invalid value") 
    }

    return num
}

func AskUserForString(message string) string {
    var str string
    fmt.Print(message)
    _, err := fmt.Scanln(&str)

    if err != nil {
        log.Fatalf("Invalid value") 
    }

    return str
}
