#Turkish Identification Number Validation

Validation of Turkish Identification Number in golang

[![GoDoc](https://godoc.org/github.com/aligoren/gotcvalidate?status.svg)](https://godoc.org/github.com/aligoren/gotcvalidate)

The identification number is a unique 11-digit number given by the MERNIS computer on the basis of the citizen's registration record that is kept by the registration office. The number does not reflect any personal information about the citizen. It is not possible to change the identification number once applied.

[www.nvi.gov.tr/English/Mernis_EN,Mernis_En.html?pageindex=1](www.nvi.gov.tr/English/Mernis_EN,Mernis_En.html?pageindex=1)

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

Run your test code

~~~~{.shell}
go run main.go

True
~~~~
