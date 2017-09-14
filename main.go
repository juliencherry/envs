package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var envStateFilepath string

func main() {
	if envStateFilepath == "" {
		envStateFilepath = filepath.Join(os.Getenv("HOME"), ".envs")
	}

	if len(os.Args) < 2 {
		fmt.Println("Invalid command")
		os.Exit(1)
	}
	command := os.Args[1]

	var commandErr error
	switch command {
	case "cf-add-target":

		if len(os.Args) < 3 {
			fmt.Println("Must specify target to add")
			os.Exit(1)
		}

		if err := cfAddTarget(os.Args[2]); err != nil {
			commandErr = fmt.Errorf("Failed to add the target: %s", err.Error())
		}
	case "cf-targets":
		if err := cfTargets(); err != nil {
			commandErr = fmt.Errorf("Failed to list targets: %s", err.Error())
		}
	default:
		commandErr = errors.New("Invalid command")
	}

	if commandErr != nil {
		fmt.Println(commandErr.Error())
		os.Exit(1)
	}
}

func cfAddTarget(targetName string) error {

	contents, err := getFileContents(envStateFilepath)
	if err != nil {
		return err
	}

	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		if line == targetName {
			return fmt.Errorf("target “%s” already exists", targetName)
		}
	}

	file, err := os.OpenFile(envStateFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(targetName + "\n")
	if err != nil {
		return err
	}

	return nil
}

func cfTargets() error {
	envState, err := getFileContents(envStateFilepath)
	if err != nil {
		return err
	}

	if envState == "" {
		fmt.Println("No targets available")
	} else {
		fmt.Print(envState)
	}

	return nil
}

func getFileContents(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)

	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

	return string(contents), nil
}
