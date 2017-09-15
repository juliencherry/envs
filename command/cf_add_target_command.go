package command

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
)

type CFAddTargetCommand struct {
	StateManager StateManager
}

func (c CFAddTargetCommand) Run(args []string) (string, error) {
	flags := flag.NewFlagSet("cf-add-target", flag.ContinueOnError)
	flags.SetOutput(ioutil.Discard)

	envName := flags.String("name", "", "")
	api := flags.String("api", "", "")
	username := flags.String("username", "", "")
	password := flags.String("password", "", "")
	skipSSLValidation := flags.Bool("skip-ssl-validation", false, "")

	flags.StringVar(envName, "n", "", "")
	flags.StringVar(api, "a", "", "")
	flags.StringVar(username, "u", "", "")
	flags.StringVar(password, "p", "", "")
	flags.BoolVar(skipSSLValidation, "s", false, "")

	err := flags.Parse(args)
	if err != nil {
		return "", fmt.Errorf("cannot parse flags: %s", err)
	}

	if *envName == "" {
		return "", errors.New("`name` flag required")
	}

	if *api == "" {
		return "", errors.New("`api` flag required")
	}

	if *username == "" {
		return "", errors.New("`username` flag required")
	}

	if *password == "" {
		return "", errors.New("`password` flag required")
	}

	err = c.StateManager.SaveEnv(*envName, *api, *username, *password, *skipSSLValidation)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`Added target "%s"`, *envName), nil
}
