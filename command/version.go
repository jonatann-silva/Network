package command

import (
	"context"
	"strings"

	cli "jonatann-silva/network/pkg/cli"
	version "jonatann-silva/network/version"
)

// VersionCommand :
type VersionCommand struct {
	UI cli.UI
}

// Name :
func (c *VersionCommand) Name() string {
	return "version"
}

// Synopsis :
func (c *VersionCommand) Synopsis() string {
	return "Print the DSN version"
}

// Run :
func (c *VersionCommand) Run(ctx context.Context, args []string) int {
	c.UI.Output(version.GetVersion().VersionNumber())
	return 0
}

// Help :
func (c *VersionCommand) Help() string {
	h := `
Usage: dsn version [options]

  Version prints out the DSN version.

General Options:
` + GlobalOptions() + `
`
	return strings.TrimSpace(h)
}
