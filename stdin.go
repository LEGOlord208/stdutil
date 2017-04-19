package stdutil

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var scanner *bufio.Scanner

// Scanln scans a line from STDIN.
// It returns any errors, as well
// as printing it, if it's not EOF.
func Scanln() (string, error) {
	if scanner == nil {
		scanner = bufio.NewScanner(os.Stdin)
	}

	if scanner.Scan() {
		return scanner.Text(), nil
	}

	err := scanner.Err()
	if err == nil {
		return "", io.EOF
	}

	PrintErr("Could not read", err)
	return "", err
}

// ScanTrim scans a line from STDIN, but trims the result.
// Returns any errors.
func ScanTrim() (string, error) {
	str, err := Scanln()
	return strings.TrimSpace(str), err
}

// MustScanln scans a line from STDIN.
// If failed, exit the program.
func MustScanln() string {
	in, err := Scanln()
	if err != nil {
		os.Exit(1)
		return ""
	}
	return in
}

// MustScanTrim scans a line from STDIN, but trims the result.
// If failed, exit the program.
func MustScanTrim() string {
	return strings.TrimSpace(MustScanln())
}

// MustScanLower scans a line from STDIN, but trims and makes the result lowercase.
// If failed, exit the program.
func MustScanLower() string {
	return strings.ToLower(MustScanTrim())
}
