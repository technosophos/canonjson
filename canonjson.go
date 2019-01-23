package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/docker/go/canonical/json"
)

func main() {
	if len(os.Args) < 2 {
		fail("the path to a JSON file is required")
	}

	file := os.Args[1]
	if fi, err := os.Stat(file); err != nil {
		fail("no such file: %s", file)
	} else if fi.IsDir() {
		fail("%s is a directory. JSON file required", file)
	}

	data, err := ioutil.ReadFile(file)
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
	fmt.Println(string(out))

}

func fail(msg string, v ...interface{}) {
	fmt.Printf(msg+"\n", v...)
	os.Exit(1)
}
