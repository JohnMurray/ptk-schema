package main

import (
	"flag"
	"fmt"
)


/*
 * Represents a command available to the program
 */
type Command struct {
	action      func()
	description string
}


/**
 * Parse incomign flags, dispatch to appropriate command or show help
 * message.
 */
func main() {
	flag.Parse()
	var args []string = flag.Args()

	if len(args) == 0 {
		println("No command was given\n")
		helpCommand()
	} else {
		commands := getCommands()

		if cmd, ok := commands[args[0]]; ok {
			cmd.action()
		} else {
			helpCommand()
		}
	}
}


/**
 * Return a map of command-names (from user-input) to functions to
 * run.
 */
func getCommands() map[string]Command {
	return map[string]Command{
		"help": Command{
			helpCommand,
			"Prints this help message",
		},
	}
}

/**
 * COMMAND 'help'
 *
 * Print help information to the user... Straight forward enough taht I'll
 * just stop typing about it now.
 *
 * TODO: define non-command help ouput in terms of flags specified for each
 *       command (not quite sure how to do this)
 */
func helpCommand() {
	commands := getCommands()

	println("Usage: schema [command] [options]")
	println("")
	println("Commands:")
	for k, v := range commands {
		fmt.Printf("  %s         %s\n", k, v.description)
	}
}
