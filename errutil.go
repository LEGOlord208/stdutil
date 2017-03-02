package stdutil;

import (
	"fmt"
	"os"
	"runtime/debug"
	"io"
)

var (
	// Set whether or not PrintErr() should print a stack trace.
	ShouldTrace bool;

	// Sets the output PrintErr() should use.
	// Defaults to os.Stderr
	ErrOutput io.Writer = os.Stderr;

	// A slice of functions (events) to call before PrintErr() is called.
	// The final error message that will be printed out it sent as argument.
	// If returned true, the function is cancelled completely.
	EventPrePrintError []func(string) bool;

	// A slice of functions (events) to call after PrintErr() is called.
	// The final error message that will be printed out it sent as argument.
	EventPostPrintError []func(string);
)

// Prints an error message.
// Any argument may be zero.
func PrintErr(msg string, err error){
	var text string;
	if(err == nil){
		text = msg;
	} else if(msg == ""){
		text = "An error occured: " + err.Error();
	} else {
		text = msg + ": " + err.Error();
	}

	defer func(){
		for _, event := range EventPostPrintError{
			event(text);
		}
	}();

	for _, event := range EventPrePrintError{
		if(event(text)){
			return;
		}
	}

	fmt.Fprintln(ErrOutput, text);

	if(ShouldTrace){
		debug.PrintStack();
	}
}
