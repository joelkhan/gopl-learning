// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	fmt.Println("name:\n ", os.Args[0])
	fmt.Println("args:")
	for idx, arg := range os.Args[1:] {
		fmt.Println(" ", idx, arg)
	}
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}

//!-
