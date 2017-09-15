package command

import (
	"errors"
)

func Build(cmd string, stateManager StateManager) (Command, error) {
	switch cmd {
	case "cf-add-target":
		return CFAddTargetCommand{stateManager}, nil
	case "cf-target":
		return CFTargetCommand{stateManager}, nil
	case "cf-targets":
		return CFTargetsCommand{stateManager}, nil
	default:
		return nil, errors.New("cannot build command")
	}
}
