package env

import (
	"fmt"
	"os"
)

type Key string

func OrDefault(env Key, defaultValue string) string {
	value, ok := os.LookupEnv(string(env))
	if !ok {
		return defaultValue
	}

	return value
}

func Must(env Key) string {
	value, ok := os.LookupEnv(string(env))
	if !ok {
		panic(fmt.Errorf("env var %s is not set", env))
	}

	return value
}
