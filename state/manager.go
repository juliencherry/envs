package state

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Manager struct {
	Path string
}

func (m Manager) GetEnvs() ([]string, error) {
	contents, err := ioutil.ReadFile(m.Path)

	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return nil, nil
	}

	envs := strings.TrimSpace(string(contents))
	if len(envs) == 0 {
		return nil, nil
	}

	return strings.Split(envs, "\n"), nil
}

func (m Manager) SaveEnv(env string) error {
	contents, err := ioutil.ReadFile(m.Path)

	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, line := range strings.Split(string(contents), "\n") {
		if env == line {
			return fmt.Errorf("target already exists")
		}
	}

	file, err := os.OpenFile(m.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(env + "\n")
	if err != nil {
		return err
	}

	return nil
}
