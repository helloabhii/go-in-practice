package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] // taking arguments/
	result, _ := concat(args...)
	fmt.Printf("Concatenated String: '%s'\n", result)
}

func concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("no strings supplied")
	}
	return strings.Join(parts, " "), nil
}
