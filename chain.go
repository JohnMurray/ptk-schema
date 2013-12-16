/*
 * File: chain.go
 *
 * Purpose: Contains all of the utilities related to obtaining, validating,
 *			and working with alter-chains.
 */

package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type Direction int

const (
	up   Direction = iota
	down Direction = iota
)

var (
	ErrNotSchemaDir = errors.New(
		"user: current directory is not a valid schema directory")
)

const metaFileName = ".schema.meta"

type Chain struct {
	ref      string
	backref  *Chain
	hasDown  bool
	downFile string
	upFile   string
}

type Meta struct {
	ref       string
	backref   string
	direction Direction
	filename  string
}

// func GetMeta() []Meta* {

// }

// func GetChain() []Chain* {
// 	return GetChainWithMeta(GetMeta())
// }

// func GetChainWithMeta([]Meta*) []Chain* {
// 	return []Chain*{}
// }

/**
 * For the current working directory, return a list of files that meet
 * the criteria to be an alter (fit the naming convention) if the CWD
 * is a schema directory. If the current directory is not a schema dir,
 * then return and error
 */
func fileList(conf *AppConfig) ([]string, error) {
	if !cwdIsSchemaDir() {
		return nil, ErrNotSchemaDir
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}

	alterFiles := make([]string, len(files), cap(files))
	i := 0
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), conf.AlterExt) {
			alterFiles[i] = file.Name()
			i += 1
		}
	}

	alterFiles = alterFiles[:i]

	return alterFiles, nil
}

/**
 * Check whether the current working directory is a schema directory. This
 * can be checked by looking for the .schema.meta file within the CWD.
 */
func cwdIsSchemaDir() bool {
	if currDir, err := os.Getwd(); err == nil {
		metaFile := currDir + string(os.PathSeparator) + metaFileName

		_, err = os.Stat(metaFile)
		return err == nil
	}
	return false
}

/**
 * Given a (valid) path to a file, read through the file and return a Meta
 * object. If no file is returned, then a nil Meta object may be returned.
 */
func parseMeta(filepath string) /* *Meta*/ {

}
