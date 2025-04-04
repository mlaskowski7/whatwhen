package cli

import "github.com/spf13/cobra"

var cli = &cobra.Command{
	Use:   "whatwhen",
	Short: "A CLI tool to manage your whats and your whens",
}

func Execute() error {
	return cli.Execute()
}

func init() {
	cli.AddCommand(viewCmd)

	AddCreateCmdFlags()
	cli.AddCommand(createCmd)
}
