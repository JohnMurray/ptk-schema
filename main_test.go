package main

import (
	"os"
	"testing"
)

func testSetup(t *testing.T) {
	// validate in project directory running tests
	if _, err := os.Stat("main_test.go"); os.IsNotExist(err) {
		t.Error("main_test.go was not found. Make sure " +
			"to run tests within the root project directory")
	}

	// change current working directory to the test-dir
	if err := os.Chdir("./etc/test-env"); err != nil {
		t.Fatal("Could not find test-directory './etc/test-env'")
	}
}

func testTeardown(t *testing.T) {
	// go back up to top-level directory
	if err := os.Chdir("./../.."); err != nil {
		t.Fatal("Could not go _back_ to root direcotry from test " +
			"directory './etc/test-env'")
	}
}

func TestGetCurrentSchemDir(t *testing.T) {
	testSetup(t)
	defer testTeardown(t)
}

func TestFileLoad(t *testing.T) {
	testSetup(t)
	defer testTeardown(t)
}
