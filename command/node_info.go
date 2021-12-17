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

// NodeInfoCommand :
type NodeInfoCommand struct {
	UI cli.UI
	Command

	// Parsed flags
	self bool
	json bool
}

func (c *NodeInfoCommand) FlagSet() *pflag.FlagSet {

	flags := c.Command.FlagSet(c.Name())
	flags.Usage = func() { c.UI.Output("\n" + c.Help() + "\n") }

	// General options
	flags.BoolVar(&c.self, "self", false, "")
	flags.BoolVar(&c.json, "json", false, "")

	return flags
}

// Name :
func (c *NodeInfoCommand) Name() string {
	return "node info"
}

// Synopsis :
func (c *NodeInfoCommand) Synopsis() string {
	return "Display info about a node"
}

// Run :
func (c *NodeInfoCommand) Run(ctx context.Context, args []string) int {

	flags := c.FlagSet()

	if err := flags.Parse(args); err != nil {
		return 1
	}

	args = flags.Args()
	if len(args) != 1 && !c.self {
		c.UI.Error("This command takes either one argument or a --self flag")
		c.UI.Error(`For additional help, try 'dsn node info --help'`)
		return 1
	}

	// Get the HTTP client
	api, err := c.Command.APIClient()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error setting up API client: %s", err))
		return 1
	}

	// Print status of a single node
	var nodeID string
	if !c.self {
		nodeID = args[0]
	} else {
		if nodeID, err = localAgentNodeID(api); err != nil {
			c.UI.Error(fmt.Sprintf("Error determining local node ID: %s", err))
			return 1
		}
	}

	node, err := api.Nodes().Get(nodeID)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error retrieving node status: %s", err))
		return 1
	}

	c.UI.Output(c.formatNode(node))

	return 0
}

// Help :
func (c *NodeInfoCommand) Help() string {
	h := `
Usage: dsn node info <node_id> [options]

  Display detailed information about an existing node.

  If ACLs are enabled, this option requires a token with the 'node:read' capability.

General Options:
` + GlobalOptions() + `

Node List Options:

  --self
    Query the status of the local node.

  --json
    Enable JSON output.

`
	return strings.TrimSpace(h)
}

func (c *NodeInfoCommand) formatNode(node *structs.Node) string {

	var b bytes.Buffer

	if c.json {
		enc := json.NewEncoder(&b)
		enc.SetIndent("", "    ")

		fnode := map[string]string{
			"id":               node.ID,
			"name":             node.Name,
			"advertiseAddress": node.AdvertiseAddress,
			"status":           node.Status,
		}

		if err := enc.Encode(fnode); err != nil {
			c.UI.Error(fmt.Sprintf("Error formatting JSON output: %s", err))
		}

	} else {
		tbl := table.New("NODE ID", "NAME", "ADVERTISE ADDRESS", "STATUS").WithWriter(&b)
		tbl.AddRow(node.ID, node.Name, node.AdvertiseAddress, node.Status)
		tbl.Print()
	}

	return b.String()
}
