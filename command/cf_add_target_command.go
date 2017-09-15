package command

import (
	"errors"
	"fmt"
)

type CFAddTargetCommand struct {
	StateManager StateManager
}

func (c CFAddTargetCommand) Run(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("Missing required argument")
	}

	targetName := args[0]

	err := c.StateManager.SaveEnv(targetName)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`Added target "%s"`, targetName), nil
}
