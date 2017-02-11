package stdutil;

import (
	"bufio"
	"os"
	"io"
	"strings"
)

var scanner *bufio.Scanner;

func INit() bool{
	if(IsINit()){
		return false;
	}
	scanner = bufio.NewScanner(os.Stdin);
	return true;
}
func IsINit() bool{
	return scanner != nil;
}

func Scanln() (string, error){
	INit();

	if(scanner.Scan()){
		return scanner.Text(), nil;
	} else {
		err := scanner.Err();
		if(err == nil){
			PrintErr("End of file from STDIN", nil);
			return "", io.EOF;
		} else {
			PrintErr("Could not read from STDIN", err);
			return "", err;
		}
	}
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
func MustScanLower() string{
	return strings.ToLower(MustScanln());
}
