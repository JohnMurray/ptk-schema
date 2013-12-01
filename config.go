package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

type Config struct {
	CommentToken string
}

/**
 * Return the config for the schema tool. This config could either be global
 * or user-local/provided depending on what is found. The tool will look in
 * the following places, in order. Configs found 'later' (in the order
 * specified) will overwrite previous configurations where applicable. This
 * means that each subsequent configuration can become more specific than the
 * previous.
 *
 * Order:
 *   + Static Default - (defined in this file) (incomplete configuration)
 *   + /etc/schema.conf
 *   + /usr/local/etc/schema.conf
 *   + $HOME/.schema.conf
 *   + ./.schema.conf - (current working directory)
 *
 * This function guarantees that a complete configuration object is always
 * returned. Failure to do so will result in a runtime-panic with a detailed
 * message of what went wrong.
 *
 * Returns a Config struct
 */
func GetConfig() *Config {
	var config = new(Config)

    var configLocations = getConfigPaths()

	for i := range configLocations {
		config = conflateConfigs(config,
			readAndParseJson(configLocations[i]))
	}

	return config
}

/**
 * Return all the paths that we will try to load configs from. The paths
 * returned will be in priority order (first = lowest, last = highest)
 */
func getConfigPaths() []string {
	var userConf string
	usr, err := user.Current()
	if err == nil {
		userConf = usr.HomeDir + string(os.PathSeparator) + ".schema.conf"
	}
	return []string{
		"/etc/schema.conf",
		"/usr/local/etc/schema.conf",
		userConf,
		"./.schema.conf",
	}
}

/**
 * Givne two configs, take care of merging them where appropriate. Any fields
 * that are specified within the override that are nil will be disregarded in
 * terms of override. This means that nil cannot be treated as a useful value
 * by the application.
 */
func conflateConfigs(original *Config, overrides *Config) *Config {
	newConf := new(Config)

	if overrides.CommentToken != "" {
		newConf.CommentToken = overrides.CommentToken
	} else {
		newConf.CommentToken = original.CommentToken
	}

	return newConf
}

/**
 * Return a config object for a JSON config file given the file name. If no
 * file happens to be found, then just return an empty Config object.
 */
func readAndParseJson(filename string) *Config {
	config := new(Config)

	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("%+v\n", e)
		return config
	}

	err := json.Unmarshal(file, config)

	if err != nil {
	}

	return config
}
