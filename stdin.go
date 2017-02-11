package stdutil;

import (
	"bufio"
	"os"
	"io"
	"strings"
)

var scanner *bufio.Scanner;

func Scanln() (string, error){
	if(scanner != nil){
		scanner = bufio.NewScanner(os.Stdin);
	}

	if(scanner.Scan()){
		return scanner.Text(), nil;
	} else {
		err := scanner.Err();
		if(err == nil){
			return "", io.EOF;
		} else {
			PrintErr("Could not read", err);
			return "", err;
		}
	}
}
func ScanTrim() (string, error){
	str, err := Scanln();
	return strings.TrimSpace(str), err;
}

func MustScanln() string{
	in, err := Scanln();
	if(err != nil){
		os.Exit(1);
		return "";
	} else {
		return in;
	}
}
func MustScanTrim() string{
	return strings.TrimSpace(MustScanln());
}
func MustScanLower() string{
	return strings.ToLower(MustScanTrim());
}
