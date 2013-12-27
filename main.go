package main

import (
	"./utils/chain"
	"./utils/config"
	"flag"
	"fmt"
	"os"
)

/*
 * Represents data that is required by the command
 */
type CommandContext struct {
	config *config.AppConfig
}

/*
 * Represents a command available to the program
 */
type Command struct {
	action      func(CommandContext)
	description string
	context     CommandContext
}

const META_FILENAME = ".schema.meta"

/**
 * Look at the first argument coming in and dispatch to the appropriate
 * command-handler. Flag/option parsing will be done by each command
 * separately.
 */
func main() {
	var args []string = os.Args[1:] // first arg is command-name

	if len(args) == 0 {
		helpCommand(CommandContext{})
	} else {
		// reset command to avoid uneccessary warnings
		if args[0] == "-h" || args[0] == "--help" {
			args[0] = "help"
		} else if args[0] == "-v" || args[0] == "--version" {
			args[0] = "version"
		}

		commands := getCommands()

		// dispatch
		if cmd, ok := commands[args[0]]; ok {
			cmd.action(cmd.context)
		} else {
			fmt.Printf("Command not found '%s'\n", args[0])
			helpCommand(commands["help"].context)
		}
	}
}

/**
 * Return a map of command-names (from user-input) to functions to
 * run.
 */
func getCommands() map[string]Command {
	context := CommandContext{
		config: config.GetAppConfig(),
	}
	return map[string]Command{
		"help": Command{
			helpCommand,
			"Prints this help message",
			context,
		},
		"new": Command{
			newCommand,
			"Creates a new alter file (an up and possibly a down alter)",
			context,
		},
		"version": Command{
			versionCommand,
			"Lists the current version",
			context,
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
 * USAGE: schema help|-h|--help|
 *
 * TODO: define non-command help ouput in terms of flags specified for each
 *       command (not quite sure how to do this)
 */
func helpCommand(cmdcontext CommandContext) {
	commands := getCommands()

	println("Usage: schema [command] [options]")
	println("")
	println("Commands:")

	for k, v := range commands {
		fmt.Printf("  %-8s         %s\n", k, v.description)
	}
}

/**
 * COMMAND 'version'
 *
 * Prints out the current version.
 *
 * USAGE: schema version|-v|--version
 */
func versionCommand(cmdcontext CommandContext) {
	fmt.Printf("Version: %s\n", Version)
}

/**
 * COMMAND 'new'
 *
 * Creates a new 'up' and 'down' file-set which represents a single alter in
 * the chain. The file will be created at the end (defined as the furthest
 * point on the chain) of the chain.
 *
 * USAGE: schema new [options] filename
 *
 * Options:
 *  -n  --no-down     Do not create a down alter (could fail 'check' command)
 *
 * TODO: implement
 */
func newCommand(cmdcontext CommandContext) {
	context := &chain.ChainContext{
		AlterExt:     cmdcontext.config.AlterExt,
		MetaFileName: META_FILENAME,
	}

	if !chain.CwdIsSchemaDir(context) {
		fmt.Print("The current directory contains not alters. Try running 'init' first.\n")
		os.Exit(1)
	}

	// parse stuff
	flag.Bool("name", true, "usage")
	flag.Parse()

	// tailRef := getTailRef()
	// println(tailRef)
}
