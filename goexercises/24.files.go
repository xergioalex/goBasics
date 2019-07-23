package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

func main() {
	// Get information about program execution
	_, filename, _, ok := runtime.Caller(0)
	// Validate execution
	if !ok {
		panic("No caller information")
	}
	/* Read a file */
	data, err := ioutil.ReadFile(path.Dir(filename) + "/test.txt")
	// Stop execution if error any error happen
	if err != nil {
		panic(err)
	}
	// Print file content
	fmt.Println(string(data))
}