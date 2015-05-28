#Turkish Identification Number Validation

Validation of Turkish Identification Number in golang

#Usage:

Check main file in test folder or follow this code:

~~~~{.shell}
go get github.com/aligoren/gotcvalidate
~~~~

and import your project

~~~~{.go}
import "github.com/aligoren/gotcvalidate"
~~~~

#Sample Usage:

~~~~{.go}
package main

import (
    "fmt"
    "github.com/aligoren/gotcvalidate"
)

func main() {
    /* initialize TCValidate (TCValidate) */
    ti := gotcvalidate.TCValidate{}
    /* First number was given to Atatürk, which is 10000000146. */
    if ti.TiValidate("10000000146") {
        fmt.Println("True")
    }else{
        fmt.Println("False")
    }
}

~~~~
