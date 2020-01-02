package serial

import (
    "regexp"
    "strconv"
    "strings"
)

func removeCRLF(str string) string {
    return strings.ReplaceAll(strings.ReplaceAll(str, "\r", ""), "\n", "")
}

func extractIntFromStr(s string) (int, error) {
    re := regexp.MustCompile(`[^0-9.]`)
    i, err := strconv.Atoi(re.ReplaceAllString(s, `$1`))

    if err != nil {
        return -1, err
    }

    return i, nil
}
