package main

import (
    "fmt"
    "github.com/dantheman213/serial"
)

func main() {
    ports, err := serial.ListPorts()
    if err != nil {
        panic(err)
    }

    fmt.Print(ports)
}
