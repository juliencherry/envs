package command

import "strings"

type CFTargetsCommand struct {
	StateManager StateManager
}

func (c CFTargetsCommand) Run(args []string) (string, error) {
	envs, err := c.StateManager.GetEnvs()
	if err != nil {
		return "", err
	}

	if len(envs) == 0 {
		return "No targets available", nil
	}
	return strings.Join(envs, "\n"), nil
}
