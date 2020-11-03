package envreader

import (
	"os"
	"strconv"
)

// ReadEnvAsInt reads env variable and converts it to string, if not found or failed returns default value
func ReadEnvAsInt(envName string, defaultValue int) int {
	envValue, exists := os.LookupEnv(envName)
	if !exists {
		return defaultValue
	}

	integerEnvValue, err := strconv.Atoi(envValue)
	if err != nil {
		return defaultValue
	}

	return integerEnvValue
}

// ReadEnvAsString reads env variable as string, if not found or failed returns default value
func ReadEnvAsString(envName string, defaultValue string) string {
	envValue, exists := os.LookupEnv(envName)

	if exists {
		return envValue
	}
	return defaultValue
}
