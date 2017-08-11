package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

var cmdMerge = &Command{
	UsageLine: "merge [-m msg] base commit...",
	Short:     "merge the named commits into the base branch",
	Long: `
Merge performs the merging commits into the base branch.

The -m flag sets the commit message to be used for the merge commit.
`,
}

var mergeMsg string

func init() {
	cmdMerge.Run = runMerge
	cmdMerge.Flag.StringVar(&mergeMsg, "m", "", "")
}

func runMerge(cmd *Command, args []string) {
	if len(args) != 2 {
		return
	}

	var msg *string
	if mergeMsg != "" {
		msg = &mergeMsg
	}
	res, err := client.PerformMerge(args[0], args[1], msg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var buf bytes.Buffer
	e := json.NewEncoder(&buf)
	if err := e.Encode(res); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Fprintln(os.Stdout, buf.String())
}
