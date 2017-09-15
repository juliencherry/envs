package command

type CFTargetCommand struct {
	StateManager StateManager
}

func (CFTargetCommand) Run(args []string) (string, error) {
	return "No environment targeted", nil
}
