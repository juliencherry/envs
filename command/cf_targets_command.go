package command

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type CFTargetsCommand struct{}

func (CFTargetsCommand) Run(args []string) (string, error) {
	envStateFilepath := os.Getenv("ENV_STATE_FILEPATH")
	if envStateFilepath == "" {
		envStateFilepath = filepath.Join(os.Getenv("HOME"), ".envs")
	}

	envState, err := getFileContents(envStateFilepath)
	if err != nil {
		return "", err
	}

	if envState == "" {
		return "No targets available", nil
	}
	return envState, nil

}

func getFileContents(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)

	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

	return strings.TrimSpace(string(contents)), nil
}
