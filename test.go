package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonConfig struct {
	CommentToken string
}

func main() {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File Error: %v\n", e)
		os.Exit(1)
	}

	// json_string := string(file)
	// fmt.Println(json_string)

	config := new(JsonConfig)
	err := json.Unmarshal(file, &config)

	if err != nil {
		fmt.Printf("Parse Error: %v\n", err)
	}

	fmt.Printf("%+v\n", config)
    println(config.CommentToken)
}
