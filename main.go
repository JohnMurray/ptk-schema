package main

import (
	"flag"
)

func main() {
	/*
	 * Parse incomign flags, dispatch to appropriate command or show help
	 * message.
	 */
	flag.Parse()
	var args []string = flag.Args()

	if len(args) == 0 {
		println("No command was given\n")
		helpCommand()
	} else {
		commands := getCommands()

		if cmd, ok := commands[args[0]]; ok {
            cmd()
        } else {
            helpCommand()
        }
	}
}

/**
 * Return a map of command-names (from user-input) to functions to
 * run.
 */
func getCommands() map[string]func() {
	return map[string]func(){
		"help": helpCommand,
	}
}


/**
 * Print help information to the user... Straight forward enough taht I'll
 * just stop typing about it now.
 *
 * TODO: define non-command help ouput in terms of flags specified for each
 *       command (not quite sure how to do this)
 */
func helpCommand() {
	println("Usage: schema [command] [options]")
	println("")
	println("Commands:")
	println("  help         Print this help message")
}
