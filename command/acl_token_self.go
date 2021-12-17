package command

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	structs "jonatann-silva/network/dsn/structs"
	cli "jonatann-silva/network/pkg/cli"

	table "github.com/rodaine/table"
	"github.com/spf13/pflag"
)

// ACLTokenSelfCommand :
type ACLTokenSelfCommand struct {
	UI cli.UI
	Command

	// Parsed flags
	json bool
}

func (c *ACLTokenSelfCommand) FlagSet() *pflag.FlagSet {

	flags := c.Command.FlagSet(c.Name())
	flags.Usage = func() { c.UI.Output("\n" + c.Help() + "\n") }

	// General options
	flags.BoolVar(&c.json, "json", false, "")

	return flags
}

// Name :
func (c *ACLTokenSelfCommand) Name() string {
	return "acl token self"
}

// Synopsis :
func (c *ACLTokenSelfCommand) Synopsis() string {
	return "Lookup self ACL token"
}

// Run :
func (c *ACLTokenSelfCommand) Run(ctx context.Context, args []string) int {

	flags := c.FlagSet()

	if err := flags.Parse(args); err != nil {
		return 1
	}

	args = flags.Args()
	if len(args) > 0 {
		c.UI.Error("This command takes no arguments")
		c.UI.Error(`For additional help, try 'dsn acl token self --help'`)
		return 1
	}

	// Get the HTTP client
	api, err := c.Command.APIClient()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error setting up API client: %s", err))
		return 1
	}

	token, err := api.ACLTokens().Self()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error retrieving self ACL token: %s", err))
		return 1
	}

	c.UI.Output(c.formatToken(token))

	return 0
}

// Help :
func (c *ACLTokenSelfCommand) Help() string {
	h := `
Usage: dsn acl token self <name> [options]

  Display information on the currently set ACL policy.

  Use the --json flag to see a detailed list of the rules associated with the token.

General Options:
` + GlobalOptions() + `

ACL Token Info Options:

  --json
    Enable JSON output.

`
	return strings.TrimSpace(h)
}

func (c *ACLTokenSelfCommand) formatToken(token *structs.ACLToken) string {

	var b bytes.Buffer

	if c.json {
		enc := json.NewEncoder(&b)
		enc.SetIndent("", "    ")
		formatted := map[string]interface{}{
			"id":        token.ID,
			"name":      token.Name,
			"type":      token.Type,
			"secret":    token.Secret,
			"policies":  token.Policies,
			"createdAt": token.CreatedAt,
			"updatedAt": token.UpdatedAt,
		}
		if err := enc.Encode(formatted); err != nil {
			c.UI.Error(fmt.Sprintf("Error formatting JSON output: %s", err))
		}
	} else {
		tbl := table.New("TOKEN ID", "NAME", "TYPE", "SECRET", "POLICIES").WithWriter(&b)
		tbl.AddRow(token.ID, token.Name, token.Type, token.Secret, len(token.Policies))
		tbl.Print()
	}

	return b.String()
}
