package stdutil;

import (
	"fmt"
	"os"
	"runtime/debug"
)

// Set whether or not PrintErr should print a stack trace.
var ShouldTrace bool;

// Prints an error message.
// Both arguments may be zero.
func PrintErr(text string, err error){
	if(err == nil){
		fmt.Fprintln(os.Stderr, text);
	} else if(text == ""){
		fmt.Fprintln(os.Stderr, "An error occured:", err);
	} else {
		fmt.Fprintln(os.Stderr, text + ":", err);
	}

	if(ShouldTrace){
		debug.PrintStack();
	}
}
