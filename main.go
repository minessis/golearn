package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.litecoinpool.org/api?api_key=b9d0e9b9edda375686b1218e266e644c")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Built-in function:
	// Give me an empty byte slice that has
	// 99,999 elements inside it
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	lw := logWriter{}

	// io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
