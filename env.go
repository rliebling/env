package env

import (
	"errors"
	"os"
	"strings"
)

func GetFirst(envVars ...string) (string, error) {
	for _, envVar := range envVars {
		if envValue := os.Getenv(envVar); envValue != "" {
			return envValue, nil
		}
	}
	return "", errors.New("Environment variable(s) not found: " + strings.Join(envVars, ","))
}

func GetHomedir() (string, error) {
	homedir, err := GetFirst("HOME")
	if homedir != "" {
		return homedir, err
	}

	homeDrive, err := GetFirst("HOMEDRIVE")
	if err != nil {
		return "", err
	}

	homePath, err := GetFirst("HOMEPATH")
	if err != nil {
		return "", err
	}
	return homeDrive + homePath, nil
}
