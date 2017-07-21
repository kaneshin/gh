package main

import (
	"flag"
)

// A Command is an implementation of a gh command.
type Command struct {
	Run   func(cmd *Command, args []string)
	Usage string
}

var commands = []*Command{
	cmdMerge,
}

func init() {
	flag.Parse()
}

func main() {
}
