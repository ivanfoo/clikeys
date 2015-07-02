package main

import (
    "fmt"
    "github.com/ivanfoo/go-keys"
)

func main() {
    fmt.Println("Hiya")
    keys.GenRSAKey(2048)
}

