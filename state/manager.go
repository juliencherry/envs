package state

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Manager struct {
	Path string
}

func (m Manager) GetEnvs() ([]string, error) {
	s, err := m.getState()
	if err != nil {
		return nil, err
	}

	var envs []string
	for environment := range s.Environments {
		envs = append(envs, environment)
	}

	return envs, nil
}

func (m Manager) SaveEnv(name string, api string, username string, password string, skipSSLValidation bool) error {
	s, err := m.getState()
	if err != nil {
		return err
	}

	if _, ok := s.Environments[name]; ok {
		return errors.New("target already exists")
	}

	s.Environments[name] = environment{
		API:               api,
		Username:          username,
		Password:          password,
		SkipSSLValidation: skipSSLValidation,
	}

	contents, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(m.Path, contents, 0600)
}

func (m Manager) GetEnvDetails(envName string) (string, string, string, bool, error) {
	s, err := m.getState()
	if err != nil {
		return "", "", "", false, err
	}

	env, ok := s.Environments[envName]
	if !ok {
		return "", "", "", false, errors.New("environment not found")
	}

	return env.API, env.Username, env.Password, env.SkipSSLValidation, nil
}

func (m Manager) getState() (state, error) {
	emptyState := state{
		Environments: make(map[string]environment),
	}

	contents, err := ioutil.ReadFile(m.Path)

	if err != nil {
		if os.IsNotExist(err) {
			return emptyState, nil
		}
		return state{}, err
	}

	if len(contents) == 0 {
		return emptyState, nil
	}

	var s state
	err = json.Unmarshal(contents, &s)
	if err != nil {
		return state{}, err
	}

	return s, nil
}
