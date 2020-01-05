package serial

import "strings"

func darwinListSerialPorts() (*[]PortInfo, error)  {
    list := make([]PortInfo, 0)
    result, err := runCommand("ls -l /dev/tty.*")
    if err != nil {
        return nil, err
    }

    lines := strings.Split(result, "\n")
    if len(lines) > 1 {
        if strings.Index(lines[0], "No such file or directory") > -1 {
            // no serial devices have been detected
            return &list, nil
        } else {
            for k := 0; k < len(lines); k++ {
                line := strings.TrimSpace(removeCRLF(lines[k]))
                if line != "" {
                    fields := strings.Fields(line)

                    list = append(list, PortInfo{
                        Status: "OK", // TODO
                        PortName:    fields[9],
                        MaxBaudRate: BaudRateUnknown, // TODO
                    })
                }
            }
        }
    }

    return &list, nil
}