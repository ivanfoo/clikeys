package main

import (
    "io"
    "os"
    "log"
)

func WriteStrToFile(str string, fileName string) {
    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY,0644)
   
    if err != nil {
        log.Fatal(err)
    }

    _, err = io.WriteString(file, str)

    if err != nil {
        log.Fatal(err)
    }

    file.Close()
}
