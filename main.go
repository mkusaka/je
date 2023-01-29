package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type str struct {
	V string `json:"v,omitempty"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: je <arg>")
		os.Exit(1)
	}

	in := os.Args[1]
	fmt.Println("in", in)

	hoge := str{
		V: in,
	}
	out, err := json.Marshal(hoge)
	if err != nil {
		fmt.Println(errors.WithStack(err))
		os.Exit(1)
	}
	outString := strings.TrimSuffix(strings.TrimPrefix(string(out), `{"v":"`), `"}`)
	fmt.Println(outString)
}
