# STDutil
Minimal wrapper around some STD actions in Go... To make them quicker than ever!

## Example:
```Go
package main

import (
    "fmt"
    "github.com/jD91mZM2/stdutil"
    "regexp"
    "strings"
)

func main() {
    stdutil.ShouldTrace = true

    fmt.Print("Name: ")
    name := stdutil.MustScanLower()

    if ok, _ := regexp.MatchString("[0-9]+", name); ok {
        stdutil.PrintErr("Name can't contain numbers!!!", nil)
        return
    }

    fmt.Println("Hello, " + strings.Title(name))
}
```

Also check out its [GoDocs](https://godoc.org/github.com/jD91mZM2/stdutil)!

## Install
Should be as easy as
```
go get github.com/jD91mZM2/stdutil
```

# Have fun!
