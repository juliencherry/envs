package command

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CFAddTargetCommand struct{}

func (CFAddTargetCommand) Run(args []string) (string, error) {
	envStateFilepath := os.Getenv("ENV_STATE_FILEPATH")
	if envStateFilepath == "" {
		envStateFilepath = filepath.Join(os.Getenv("HOME"), ".envs")
	}

	if len(args) < 1 {
		return "", errors.New("Missing required argument")
	}

	targetName := args[0]
	contents, err := getFileContents(envStateFilepath)
	if err != nil {
		return "", err
	}

	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		if line == targetName {
			return "", fmt.Errorf(`target "%s" already exists`, targetName)
		}
	}

	file, err := os.OpenFile(envStateFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(targetName + "\n")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`Added target "%s"`, targetName), nil
}
