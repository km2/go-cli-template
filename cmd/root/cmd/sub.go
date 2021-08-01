package cmd

import (
	"fmt"

	"github.com/km2/go-cli-template/cli"
	"github.com/spf13/cobra"
)

var runOptions cli.SubOptions

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Sub command",
	Long:  "Sub command.",
	RunE:  runSub,
}

func runSub(cmd *cobra.Command, args []string) error {
	c := cli.New()
	if err := c.RunSub(runOptions); err != nil {
		return fmt.Errorf("failed to run sub command: %w", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(subCmd)
}
