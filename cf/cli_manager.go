package cf

import (
	"fmt"
	"os/exec"
)

type CLIManager struct{}

func (CLIManager) Login(api string, username string, password string, skipSSLValidation bool) error {
	args := []string{"login", "-a", api, "-u", username, "-p", password}
	if skipSSLValidation {
		args = append(args, "--skip-ssl-validation")
	}

	out, err := exec.Command("cf", args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to log in: %s", string(out))
	}

	return nil
}
