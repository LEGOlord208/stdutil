# STDutil
Minimal wrapper around some STD actions in Go... To make them quicker than ever!

## Example:
```
package main;

import (
	"fmt"
	"github.com/legolord208/stdutil"
	"strings"
	"regexp"
)

func main(){
	stdutil.ShouldTrace = true;

	fmt.Print("Name: ");
	name := stdutil.MustScanLower();

	if ok, _ := regexp.MatchString("[0-9]+", name); ok{
		stdutil.PrintErr("Name can't contain numbers!!!", nil);
		return;
	}

	fmt.Println("Hello, " + strings.Title(name));
}
```

Also check out it's [GoDocs](https://godoc.org/github.com/LEGOlord208/stdutil)!

## Install
Should be as easy as
```
go get github.com/legolord208/stdutil
```

# Have fun!
