package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type str struct {
	V string `json:"v,omitempty"`
}

var version string

func main() {
	showVersion := flag.Bool("version", false, "prints the version number")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	var r io.Reader

	var in string
	if len(os.Args) >= 2 {
		in = os.Args[1]
	}

	switch in {
	case "":
		r = os.Stdin
	case "-":
		r = os.Stdin
	default:
		r = bytes.NewBufferString(in)
	}

	b, err := io.ReadAll(r)

	if err != nil {
		fmt.Println(errors.WithStack(err))
		os.Exit(1)
	}

	inputString := string(b)

	inputStruct := str{
		V: inputString,
	}
	out, err := json.Marshal(inputStruct)
	if err != nil {
		fmt.Println(errors.WithStack(err))
		os.Exit(1)
	}
	outString := strings.TrimSuffix(strings.TrimPrefix(string(out), `{"v":"`), `"}`)
	fmt.Println(outString)
}
