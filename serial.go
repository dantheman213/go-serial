package serial

import (
    "errors"
    "runtime"
)

const (
    BaudRateUnknown = -1
    BaudRate115200 = 115200
)

type PortInfo struct {
    Status string
    PortName string
    MaxBaudRate int
}

func ListPorts() (*[]PortInfo, error) {
    if runtime.GOOS == "linux" {
        return linuxListSerialPorts()
    } else if runtime.GOOS == "darwin" {
        return darwinListSerialPorts()
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
