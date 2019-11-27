package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// The silly way...
	//fmt.Printf("%+v", resp)

	// The complicated way... (but we understand now the internals)
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// The "go" way of doing this...
	// io.Copy(os.Stdout, resp.Body)

	// Our own implementation of Writer interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("The body contains:", string(bs))
	fmt.Println("Number of bytes returned:", len(bs))
	return len(bs), nil
}
