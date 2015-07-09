package main

import (
    "fmt"
    "github.com/ivanfoo/urkel"
)

const KeysStorage string = "keys/"

type KeysPair struct {
    length  int
    private string
    public  string
    priFile string
    pubFile string
}

func main() {
    keysPair := new(KeysPair)

    keysPair.length = AskUserForInt("Key length: ")
    keysPair.priFile = AskUserForString("Output file for private key: ")
    keysPair.pubFile = AskUserForString("Output file for public key: ")

    fmt.Print("Generating keys...")
    keysPair.private, keysPair.public = urkel.GenPemRSA(keysPair.length)
    fmt.Println("Done!")

    WriteStrToFile(keysPair.private, KeysStorage+keysPair.priFile)
    WriteStrToFile(keysPair.public, KeysStorage+keysPair.pubFile)
}
