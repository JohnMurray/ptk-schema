package main

type Config struct {
	commentToken string
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
	return new(Config)
}
