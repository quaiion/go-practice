package main

import (
        "fmt"
        "github.com/quaiion/go-practice/conditioners/internal/streamUtils"
        "github.com/quaiion/go-practice/conditioners/internal/process"
)

func main() {
        err := process.ProcessOffice()
        if err != nil {
                streamUtils.FlushInput()
                fmt.Println(err)
        }
}
