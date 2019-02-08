package cli

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var spacesCmd = &cobra.Command{
	Use:     "spaces",
	Aliases: []string{"space"},
	Args:    cobra.ExactArgs(0),
	Short:   "Show available spaces in your cbox",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.SpacesList() },
}

var spacesCreateCmd = &cobra.Command{
	Use:   "create",
	Args:  cobra.ExactArgs(0),
	Short: "Creates a new space to your cbox",
	Long:  tools.Logo,
	Run:   func(cmd *cobra.Command, args []string) { ctrl.SpacesCreate() },
}

var spacesEditCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e", "ed"},
	Args:    cobra.ExactArgs(1),
	Short:   "Edit an space from your cbox",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.SpacesEdit(args[0]) },
}

var spacesDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Args:  cobra.ExactArgs(1),
	Short: "Destroys a space from your cbox",
	Long:  tools.Logo,
	Run:   func(cmd *cobra.Command, args []string) { ctrl.SpacesDestroy(args[0]) },
}

func init() {
	rootCmd.AddCommand(spacesCmd)
	spacesCmd.AddCommand(spacesCreateCmd)
	spacesCmd.AddCommand(spacesEditCmd)
	spacesCmd.AddCommand(spacesDestroyCmd)
}