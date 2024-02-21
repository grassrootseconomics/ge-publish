package command

type (
	CommandOpts struct {
	}

	Command struct {
	}
)

func NewCommandContainer(o CommandOpts) *Command {
	return &Command{}
}
