package stdutil;

import (
	"fmt"
	"os"
	"runtime/debug"
)

var ShouldTrace bool;

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
