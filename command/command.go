package command

type Command interface {
	Run([]string) (string, error)
}

//go:generate counterfeiter . StateManager
type StateManager interface {
	GetEnvs() ([]string, error)
	SaveEnv(env string) error
}
