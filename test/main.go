package main

import (
	"fmt"
	"github.com/aligoren/gotcvalidate"
)

func main() {
	/* initialize tiva (TCValidate) */
	ti := gotcvalidate.TCValidate{}
	/* First number was given to Atatürk, which is 10000000146. */
	if ti.TiValidate("10000000146") {
		fmt.Println("True")
	}else{
		fmt.Println("False")
	}
}
