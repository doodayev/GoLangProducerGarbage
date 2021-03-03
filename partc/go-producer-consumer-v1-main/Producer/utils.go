package main

import (
	"os"
)

// Fetch the environment variable if it exists or else return the defaultValue set in the code.
func GetValueFromEnvVariable(variableName, defaultValue string) string {
	environmentValue := os.Getenv(variableName)
	if environmentValue == "" {
		return defaultValue
	}
	return environmentValue
}
