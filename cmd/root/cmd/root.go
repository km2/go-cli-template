package cmd

import (
	"github.com/spf13/cobra"
)

var version bool

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Root command",
	Long:  "Root command.",
	RunE:  runRoot,
}

func runRoot(cmd *cobra.Command, args []string) error {
	if version {
		return runVersion(cmd, args)
	}

	return nil
}

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "V", false, "show version")
}

func Execute() {
	rootCmd.Execute() //nolint:errcheck
}
