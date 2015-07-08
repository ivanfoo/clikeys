package main

import (
    "github.com/ivanfoo/urkel"
)

type Key struct {
    length int
    private string
    public string
    outFile string
}

func main() {
    key := new(Key) 
    key.length = AskUserForInt("Key length: ")
    key.outFile = AskUserForString("Output file: ")
    key.private, key.public = urkel.GenPemRSA(key.length)
    WriteStrToFile(key.private, key.outFile)
}


