package command

type Command interface {
	Run([]string) (string, error)
}

//go:generate counterfeiter . CFCLIWrapper
type CFCLIWrapper interface {
	Login(api string, username string, password string, skipSSLValidation bool) error
}

//go:generate counterfeiter . StateManager
type StateManager interface {
	GetEnvDetails(env string) (api string, username string, password string, skipSSLValidation bool, err error)
	GetEnvs() ([]string, error)
	SaveEnv(envName string, api string, username string, password string, skipSSLValidation bool) error
}
