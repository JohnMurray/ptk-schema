package main

import (
	"fmt"
	"os"
)

/*
 * Represents a command available to the program
 */
type Command struct {
	action      func()
	description string
}

/**
 * Look at the first argument coming in and dispatch to the appropriate
 * command-handler. Flag/option parsing will be done by each command
 * separately.
 */
func main() {
	c := GetConfig()
	fmt.Printf("%+v\n", c.CommentToken)
	var args []string = os.Args[1:] // first arg is command-name

	if len(args) == 0 {
		helpCommand()
	} else {
		// reset command to avoid uneccessary warnings
		if args[0] == "-h" || args[0] == "-help" || args[0] == "--help" {
			args[0] = "help"
		}

		commands := getCommands()

		// dispatch
		if cmd, ok := commands[args[0]]; ok {
			cmd.action()
		} else {
			fmt.Printf("Command not found '%s'\n", args[0])
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

//
// COMMANDS
//
// Each method below specifies a command that can be used with the tool. Any
// utility functions that the commands use can be found below the commands in
// the 'UTILITIES' section.
//

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

//
// UTILITIES
//
// Represents a shared set of utilities that can be used to interact with
// the alter-chain and underlying file-system.
//
