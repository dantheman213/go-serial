package serial

import (
    "strconv"
    "strings"
)

// https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-serialport
const (
    Win32SerialPortStatusOK = "OK"
    Win32SerialPortStatusError = "Error"
    Win32SerialPortStatusStarting = "Starting"
    Win32SerialPortStatusStopping = "Stopping"
    Win32SerialPortStatusNoContact = "No Contact"
    Win32SerialPortStatusLostComm = "Lost Comm"
)

func windowsListSerialPorts() (*[]PortInfo, error)  {
    list := make([]PortInfo, 0)
    result, err := runCommand("wmic path Win32_SerialPort")
    if err != nil {
        if err.Error() == "No Instance(s) Available." {
            // no serial devices have been detected
            return &list, nil
        } else {
            return nil, err
        }
    }

    lines := strings.Split(result, "\n")
    for k := 1; k < len(lines); k++ {
        line := strings.TrimSpace(removeCRLF(lines[k]))
        if line != "" {
            fields := strings.Fields(line)

            b, err := strconv.Atoi(fields[13])
            if err != nil {
                return nil, err
            }

            list = append(list, PortInfo{
                Status: fields[33],
                PortName:    fields[12],
                MaxBaudRate: b,
            })
        }
    }

    return &list, nil
}