package serial

import (
    "bytes"
    "errors"
    "os/exec"
)

func runCommand(command string, args []string) (string, error) {
    cmd := exec.Command(command, args...)

    var stdout bytes.Buffer
    var stderr bytes.Buffer

    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()

    if err != nil {
        return "", err
    }

    stdoutStr := stdout.String()
    stdoutErr := stderr.String()

    if stdoutErr != "" {
        return "", errors.New(stdoutErr)
    }

    return stdoutStr, nil
}
