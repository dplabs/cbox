package cli

import (
	"github.com/dpecos/cbox/internal/pkg"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Args:    cobra.MaximumNArgs(2),
	Short:   "Search for commands in a given space",
	Long:    pkg.Logo,
	Run:     ctrl.SearchCommands,
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().BoolVarP(&viewSnippet, "view", "v", false, "Show all details about commands")
}