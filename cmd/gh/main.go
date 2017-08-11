package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kaneshin/gh"
)

// A Command is an implementation of a github command.
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string)

	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'gh help' output.
	Short string

	// Long is the long message shown in the 'gh help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet
}

// Name returns the command's name: the first word in the usage line.
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// Usage prints to standard error a usage message.
func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
	os.Exit(2)
}

// Runnable reports whether the command can be run.
func (c *Command) Runnable() bool {
	return c.Run != nil
}

type Commands []*Command

func (c Commands) get(name string) *Command {
	for _, cmd := range commands {
		if cmd.Name() == name && cmd.Runnable() {
			return cmd
		}
	}
	return nil
}

var commands = Commands([]*Command{
	cmdMerge,
})

var (
	client *gh.Client
	config *Config
)

func main() {
	var configFlag string
	var envFlag string

	flag.StringVar(&configFlag, "config", "", "")
	flag.StringVar(&envFlag, "env", "", "")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	// TODO
	if configFlag == "" {
	}
	if envFlag == "" {
	}
	config = MustNewConfig(configFlag).WithEnv(envFlag)
	client = gh.NewClient(config.AccessToken(), config.Owner(), config.Repo())

	cmd := commands.get(args[0])
	if cmd != nil {
		cmd.Flag.Usage = cmd.Usage
		cmd.Flag.Parse(args[1:])
		cmd.Run(cmd, cmd.Flag.Args())
		return
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n\n\tgh command [arguments]\n")
	os.Exit(2)
}

func help(args []string) {
	if len(args) == 0 {
		usage()
	}

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: gh help command\n\nToo many arguments given.\n")
		os.Exit(2)
	}

	cmd := commands.get(args[0])
	if cmd != nil {
		cmd.Usage()
	}
}
