package cli

import (
	"log"

	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/tools"
	"github.com/spf13/cobra"
)

var unassignTag string
var tagCmd = &cobra.Command{
	Use:   "tag",
	Args:  cobra.MinimumNArgs(1),
	Short: "Manage tags assigned to a command",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		cmdID := tools.StringToInt(args[0])
		if unassignTag != "" {
			if len(args) != 1 {
				log.Fatalf("Tags can only be deleted one by one")
			}
			db.UnassignTag(cmdID, unassignTag)
		} else {
			for _, tag := range args[1:] {
				db.AssignTag(cmdID, tag)
			}
		}

		command := db.Find(cmdID)
		tools.PrintCommand(command, false, false)
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)

	tagCmd.Flags().StringVarP(&unassignTag, "unassign", "u", "", "Unassign specified tags from command")
}
