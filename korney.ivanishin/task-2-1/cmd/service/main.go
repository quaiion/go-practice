package main

import (
        "fmt"
        "github.com/conditioners/internal/stream_utils"
        "github.com/conditioners/internal/process"
)

func main() {
        err := process.ProcessOffice()
        if err != nil {
                stream_utils.FlushInput()
                fmt.Println(err)
        }
}
