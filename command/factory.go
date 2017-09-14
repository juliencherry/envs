package command

import "errors"

type Command interface {
	Run([]string) (string, error)
}

func Build(cmd string) (Command, error) {
	switch cmd {
	case "cf-add-target":
		return CFAddTargetCommand{}, nil
	case "cf-target":
		return CFTargetCommand{}, nil
	case "cf-targets":
		return CFTargetsCommand{}, nil
	default:
		return nil, errors.New("cannot build command")
	}
}
