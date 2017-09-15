package command

import (
	"fmt"
)

type CFTargetCommand struct {
	CFCLIWrapper CFCLIWrapper
	StateManager StateManager
}

func (c CFTargetCommand) Run(args []string) (string, error) {

	if len(args) < 1 {
		return "No environment targeted", nil
	}
	envName := args[0]

	api, username, password, skipSSLValidation, err := c.StateManager.GetEnvDetails(envName)
	if err != nil {
		return "", fmt.Errorf("cannot target environment: %s", err)
	}

	if err := c.CFCLIWrapper.Login(api, username, password, skipSSLValidation); err != nil {
		return "", fmt.Errorf("cannot target environment: %s", err)
	}
	return fmt.Sprintf(`Targeted environment "%s"`, envName), nil
}
