package main

import (
    "fmt"
    "strconv"
    "strings"
    "unicode"
)

func Unpack(input string) (string, error) {
    var result strings.Builder
    escape := false

    for i, r := range input {
        if unicode.IsDigit(r) && !escape {
            if i == 0 {
                return "", fmt.Errorf("invalid string")
            }
            count, _ := strconv.Atoi(string(r))
            result.WriteString(strings.Repeat(string(input[i-1]), count-1))
        } else if r == '\\' && !escape {
            escape = true
        } else {
            escape = false
            result.WriteRune(r)
        }
    }
    return result.String(), nil
}

func main() {
    examples := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`}

    for _, example := range examples {
        result, err := Unpack(example)
        if err != nil {
            fmt.Printf("Error unpacking %s: %v\n", example, err)
        } else {
            fmt.Printf("Unpacked %s to %s\n", example, result)
        }
    }
}
