package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	t := os.Args[1]

	f, err := os.Open(t)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// _, err = io.Copy(os.Stdout, f)
	// if err != nil {
	// 	fmt.Printf("An error occurred copying the contents of file %s to stdout: %q", t, err)
	// 	os.Exit(1)
	// }
	io.Copy(os.Stdout, f)

}
