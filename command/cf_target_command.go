package command

type CFTargetCommand struct{}

func (CFTargetCommand) Run(args []string) (string, error) {
	return "No environment targeted", nil
}
