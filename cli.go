package serial

import (
    "bytes"
    "errors"
    "os/exec"
    "runtime"
)

func runCommand(command string) (string, error) {
    var cmd *exec.Cmd = nil

    switch runtime.GOOS {
    case "darwin", "linux":
        cmd = exec.Command("/bin/sh", "-c", command)
        break
    case "windows":
        cmd = exec.Command("cmd.exe", "/c", command)
        break
    default:
        break
    }

    if cmd == nil {
        return "", errors.New("unsupported OS")
    }

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
