package main

import (
        "fmt"
        "conditioners/internal/stream_utils"
        "conditioners/internal/process"
)

func main() {
        err := process.ProcessOffice()
        if err != nil {
                stream_utils.FlushInput()
                fmt.Println(err)
        }
}
