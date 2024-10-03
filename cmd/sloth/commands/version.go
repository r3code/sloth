package commands

import (
	"context"
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/slok/sloth/internal/info"
)

type versionCommand struct{}

// NewVersionCommand returns the version command.
func NewVersionCommand(app *kingpin.Application) Command {
	c := &versionCommand{}
	app.Command("version", "Shows version.")

	return c
}

func (versionCommand) Name() string { return "version" }
func (versionCommand) Run(_ context.Context, config RootConfig) error {
	fmt.Fprint(config.Stdout, info.Version)
	return nil
}
