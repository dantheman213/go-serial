package serial

import (
    "errors"
    "runtime"
)

const (
    BaudRate115200 = 115200
)

type PortInfo struct {
    Status string
    PortName string
    PortNumber int
    MaxBaudRate int
}

func ListPorts() (*[]PortInfo, error) {
    if runtime.GOOS == "linux" {
        // TODO
    } else if runtime.GOOS == "darwin" {
        // TODO
    } else if runtime.GOOS == "windows" {
        return windowsListSerialPorts()
    }

    return nil, errors.New("unknown")
}

func OpenPort() {
    // TODO
}

func ClosePort() {
    // TODO
}

func Read() {
    // TODO
}

func Write() {
    // TODO
}
