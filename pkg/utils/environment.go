package utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func Environment() error {
	// Check if the .env file exists
	if !IsFileExist(".env") {
		return errors.New(".env file does not exist")
	}

	// Read the .env file
	content, err := os.ReadFile(".env")
	if err != nil {
		return errors.New("cannot read .env file")
	}

	lines := bufio.NewScanner(strings.NewReader(string(content)))
	for lines.Scan() {
		line := lines.Text()

		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		// Interpreting each line
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key, value := parts[0], parts[1]

			// Set the environment variable
			err := os.Setenv(key, value)
			if err != nil {
				return errors.New("cannot set environment variable: " + key + " = " + value)
			}
		} else {
			LoggerWarn.Println(Warn + "Invalid line in .env file: " + line)
			continue
		}
	}

	if err := lines.Err(); err != nil {
		return errors.New("cannot read .env file")
	}

	return nil
}

func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func GetEnv(key string) string {
	err := Environment()
	if err != nil {
		LoggerError.Println(Error + err.Error() + Reset)
	}
	return os.Getenv(key)
}
