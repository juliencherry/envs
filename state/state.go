package state

type state struct {
	Environments map[string]environment
}

type environment struct {
	API               string
	Username          string
	Password          string
	SkipSSLValidation bool
}
