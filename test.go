package main

import (
    "fmt"
    "os"
	"os/user"
)

func main() {
	usr, err := user.Current()

	if err != nil {
        fmt.Printf("%+v", err)
        os.Exit(1)
	}

    println(usr.HomeDir)
}
