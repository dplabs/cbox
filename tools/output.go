package tools

import (
	"fmt"
	"strings"

	"github.com/dpecos/cmdbox/models"
	"github.com/logrusorgru/aurora"
)

func PrintCommand(cmd models.Cmd, full bool, sourceOnly bool) {
	if !sourceOnly {
		if len(cmd.Tags) != 0 {
			tags := strings.Join(cmd.Tags, ", ")
			fmt.Printf("%d - (%s) %s", aurora.Red(aurora.Bold(cmd.ID)), aurora.Brown(tags), aurora.Blue(aurora.Bold(cmd.Title)))
		} else {
			fmt.Printf("%d - %s", aurora.Red(aurora.Bold(cmd.ID)), aurora.Blue(aurora.Bold(cmd.Title)))
		}
		t := cmd.CreatedAt
		fmt.Println(aurora.Sprintf(aurora.Cyan(" %d-%02d-%02d %02d:%02d:%02d"), t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
		if full {
			if cmd.Description != "" {
				fmt.Printf("\n%s\n", aurora.Green(cmd.Description))
			}
			if cmd.URL != "" {
				fmt.Printf("\n%s\n", aurora.Blue(cmd.URL))
			}
			fmt.Printf("\n%s\n\n", cmd.Cmd)
		}
	} else {
		fmt.Println(cmd.Cmd)
	}
}

func PrintTag(tag string) {
	fmt.Printf("%s\n", aurora.Brown(tag))
}