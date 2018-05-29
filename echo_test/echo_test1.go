package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {

	// test 1
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf(s + "\r\n")

	// test 2
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf(s + "\r\n")

	// test 3
	fmt.Printf(strings.Join(os.Args[1:], " ") + "\r\n")

	// length
	fmt.Printf("len(os.Args) = %d\r\n", len(os.Args))
	fmt.Printf("len(os.Args) = %d\r\n", len(os.Args[1:]))
	// fmt.Printf("len(os.Args) = %d\r\n", len(os.Args[2:]))
	fmt.Printf("exec = " + os.Args[0] + "\r\n")
}

//!-
