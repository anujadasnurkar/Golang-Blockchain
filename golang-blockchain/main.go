package main

import (
	"os"
    "example.com/cli"
)

func main() {
	defer os.Exit(0)
	cli := CommandLine{}
	cli.Run()
}