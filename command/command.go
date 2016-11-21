package command

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/command/boot"
	"github.com/xh3b4sd/anna/command/version"
)

// New creates a new annad command.
func New() *Command {
	command := &Command{}

	command.SetBootCommand(boot.New())
	command.SetVersionCommand(version.New())

	return command
}

// Command represents the annad command.
type Command struct {
	// Dependencies.

	bootCommand    *boot.Command
	versionCommand *version.Command
}

// Execute represents the cobra run method.
func (c *Command) Execute(cmd *cobra.Command, args []string) {
	cmd.HelpFunc()(cmd, nil)
}

// New creates a new cobra command for the annad command.
func (c *Command) New() *cobra.Command {
	newCommand := &cobra.Command{
		Use:   "annad",
		Short: "Manage the daemon of the anna project. For more information see https://github.com/the-anna-project/annad.",
		Long:  "Manage the daemon of the anna project. For more information see https://github.com/the-anna-project/annad.",
		Run:   c.Execute,
	}

	newCommand.AddCommand(c.bootCommand.New())
	newCommand.AddCommand(c.versionCommand.New())

	return newCommand
}

// BootCommand returns the boot subcommand of the annad command.
func (c *Command) BootCommand() *boot.Command {
	return c.bootCommand
}

// SetBootCommand sets the boot subcommand for the annad command.
func (c *Command) SetBootCommand(command *boot.Command) {
	c.bootCommand = command
}

// SetVersionCommand sets the version subcommand for the annad command.
func (c *Command) SetVersionCommand(command *version.Command) {
	c.versionCommand = command
}

// VersionCommand returns the version subcommand of the annad command.
func (c *Command) VersionCommand() *version.Command {
	return c.versionCommand
}
