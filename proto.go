package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/resp"
)

const (
	CommandSET string = "SET"
	CommandGET string = "GET"
)

type Command interface {
}

type SetCommand struct {
	key, val string
}

func parseCommand(msg string) (Command, error) {
	rd := resp.NewReader(bytes.NewBufferString(msg))

	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if v.Type() == resp.Array {
			for i, v := range v.Array() {
				fmt.Printf("  #%d %s, value '%s'\n", i, v.Type(), v)
			}
		}
	}
	return nil, nil
}
