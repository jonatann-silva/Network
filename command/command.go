package command

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"

	api "jonatann-silva/network/api"
	cli "jonatann-silva/network/pkg/cli"
)

// Command is the base command
type Command struct {
	address string
	token   string
}

// FlagSet declares flags that are common to all commands,
// returning a pflag.FlagSet struct that will hold their values after
// pflag.Parse() is called by the command.
func (c *Command) FlagSet(name string) *pflag.FlagSet {

	flags := pflag.NewFlagSet(name, pflag.ContinueOnError)

	flags.Usage = func() {}

	flags.StringVar(&c.address, "address", "", "")
	flags.StringVar(&c.token, "token", "", "")

	// TODO: direct output to UI
	flags.SetOutput(nil)

	return flags
}

// APIClient returns a new api.Client struct,
// which can be used to interact with the DSN HTTP API.
func (c *Command) APIClient() (*api.Client, error) {
	return api.NewClient(&api.Config{
		Address: c.address,
		Token:   c.token,
	})
}

// GlobalOptions returns the global usage options string.
func GlobalOptions() string {
	text := `
  --address=<addr>
    The address of the DSN server.
    Overrides the DRAGO_ADDR environment variable if set.
    Default = http://127.0.0.1:8080

  --token=<token>
    The token used to authenticate with the DSN server.
    Overrides the DRAGO_TOKEN environment variable if set.
    Default = ""
`
	return text
}

// DefaultErrorMessage returns the default error message for this command
func DefaultErrorMessage(cmd cli.NamedCommand) string {
	return fmt.Sprintf("For additional help try 'dsn %s --help'", cmd.Name())
}

// manyStrings
type manyStrings []string

func (s *manyStrings) Set(val string) error {
	*s = append(*s, val)
	return nil
}

func (s *manyStrings) String() string {
	return strings.Join(*s, ",")
}
