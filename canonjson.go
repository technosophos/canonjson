package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/docker/go/canonical/json"
)

func main() {
	const argErr = "The first argument is the input JSON file, and the second is an optional output file."
	if len(os.Args) < 2 {
		fail(argErr)
	}

	inputFile := os.Args[1]
	if fi, err := os.Stat(inputFile); err != nil {
		fail("no such file: %s", inputFile)
	} else if fi.IsDir() {
		fail("%s is a directory. JSON file required", inputFile)
	}

	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fail(err.Error())
	}

	obj := map[string]interface{}{}
	if err := json.Unmarshal(data, &obj); err != nil {
		fail(err.Error())
	}

	out, err := json.MarshalCanonical(obj)
	if err != nil {
		fail(err.Error())
	}

	if len(os.Args) == 2 {
		fmt.Println(string(out))
	} else if len(os.Args) == 3 {
		outputFile := os.Args[2]
		f, err := os.Create(outputFile)
		if err != nil {
			fail("cannot create output file %v: %v", outputFile, err)
		}
		defer f.Close()

		_, err = f.Write(out)
		if err != nil {
			fail(err.Error())
		}
	} else {
		fail(argErr)
	}
}

func fail(msg string, v ...interface{}) {
	fmt.Printf(msg+"\n", v...)
	os.Exit(1)
}
