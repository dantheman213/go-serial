package serial

import (
    "strings"
)

func removeCRLF(str string) string {
    return strings.ReplaceAll(strings.ReplaceAll(str, "\r", ""), "\n", "")
}
