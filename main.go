package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/juliencherry/envs/command"
	"github.com/juliencherry/envs/state"
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

	cmd, err := command.Build(os.Args[1], state.Manager{Path: envStateFilepath})
	if err != nil {
		fmt.Println("Invalid command:", err)
		os.Exit(1)
	}

	msg, err := cmd.Run(os.Args[2:])
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	fmt.Println(msg)
}
