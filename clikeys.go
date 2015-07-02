package main

import (
    "fmt"
    "github.com/ivanfoo/urkel"
)

func main() {
    fmt.Println("Running...")
    urkel.GenKeyToFile(2048, "my_key")  
}

