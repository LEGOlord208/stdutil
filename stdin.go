package stdutil

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var scanner *bufio.Scanner

// Scan a line.
// Return any errors.
func Scanln() (string, error) {
	if scanner == nil {
		scanner = bufio.NewScanner(os.Stdin)
	}

	if scanner.Scan() {
		return scanner.Text(), nil
	} else {
		err := scanner.Err()
		if err == nil {
			return "", io.EOF
		} else {
			PrintErr("Could not read", err)
			return "", err
		}
	}
}

// Scan a line, but trim the result.
// Return any errors.
func ScanTrim() (string, error) {
	str, err := Scanln()
	return strings.TrimSpace(str), err
}

// Scan a line.
// If failed, exit the program.
func MustScanln() string {
	in, err := Scanln()
	if err != nil {
		os.Exit(1)
		return ""
	} else {
		return in
	}
}

// Scan a line, but trim the result.
// If failed, exit the program.
func MustScanTrim() string {
	return strings.TrimSpace(MustScanln())
}

// Scan a line, but trim and make the result lowercase.
// If failed, exit the program.
func MustScanLower() string {
	return strings.ToLower(MustScanTrim())
}
