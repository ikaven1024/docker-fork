package main

import (
	"fmt"
	"os"
	"strings"
)

type Options []string

func (o Options) Has(names ...string) bool {
	for _, opt := range o {
		for _, name := range names {
			// --name foo
			if opt == name {
				return true
			}
			// --name=foo
			if strings.HasPrefix(opt, name+"=") {
				return true
			}
		}
	}
	return false
}

func ParseFlags() (cid string, options Options) {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}
	return args[0], args[1:]
}

func usage() {
	fmt.Println("Usage: ")
}
