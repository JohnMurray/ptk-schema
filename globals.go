package main

import (
    "errors"
)

//
// GLOBAL
//
// Contains all of the variables that are global to the application. Maybe
// not the best thing, but definitely the easiest for an application of this
// size. Perhaps I'll come back at some point and git rid of this file
// (really need to read up on good Go development practices as well).
//

/*
 * Determines if debug information should be printed out during program
 * execution. Should be set from environment variable.
 */
var Debug bool = false

/*
 * The global config object (set on main) that is a 'composed' config based
 * on overrides.
 */
var Config *AppConfig

/*
 * Global set of errors that the applciation can return.
 */
var (
	ErrNotSchemaDir = errors.New(
		"user: current directory is not a valid schema directory")
)
