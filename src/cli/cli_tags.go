package cli

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:     "tags",
	Aliases: []string{"tag"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "List the tags available in your cbox",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.TagsList(optionalSelector(args, 0)) },
}

var tagsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "Delete specified tag",
	Long:  tools.Logo,
	Run:   func(cmd *cobra.Command, args []string) { ctrl.TagsDelete(args[0]) },
}

func init() {
	rootCmd.AddCommand(tagsCmd)
	tagsCmd.AddCommand(tagsDeleteCmd)
}
