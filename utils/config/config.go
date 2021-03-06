package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

type AppConfig struct {
	CommentToken string
	AlterExt     string
}

/**
 * Look at the environment variables and determine if we are running in
 * DEBUG mode or not.
 *
 * Returns a boolean value indicating debug'iness
 */
func IsDebug() bool {
	debug := os.Getenv("SCHEMA_DEBUG")
	if debug != "" {
		return true
	}
	return false
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
 * Returns a Config object
 */
func GetAppConfig() *AppConfig {
	var config = new(AppConfig)

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
func conflateConfigs(original *AppConfig, overrides *AppConfig) *AppConfig {
	newConf := new(AppConfig)

	if overrides.CommentToken != "" {
		newConf.CommentToken = overrides.CommentToken
	} else {
		newConf.CommentToken = original.CommentToken
	}

	if overrides.AlterExt != "" {
		newConf.AlterExt = overrides.AlterExt
	} else {
		newConf.AlterExt = original.AlterExt
	}

	return newConf
}

/**
 * Return a config object for a JSON config file given the file name. If no
 * file happens to be found, then just return an empty Config object.
 */
func readAndParseJson(filename string) *AppConfig {
	config := new(AppConfig)

	file, e := ioutil.ReadFile(filename)
	if e != nil {
		if IsDebug() {
			fmt.Printf("Could not read config-file: %s\n", e)
		}
		return config
	}

	err := json.Unmarshal(file, config)

	if err != nil {
		fmt.Printf("Could not read config-file: %s\n", file)
	}

	return config
}
