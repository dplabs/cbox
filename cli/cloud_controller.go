package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/dpecos/cbox/tools/console"

	"github.com/spf13/cobra"
)

func (ctrl *CLIController) CloudLogin(cmd *cobra.Command, args []string) {
	fmt.Println(tools.Logo)
	fmt.Printf("Open this URL in a browser and follow the authentication process: \n\n%s\n\n", fmt.Sprintf("%s/auth/", core.CloudURL()))

	jwt := console.ReadString("JWT Token")
	fmt.Println()

	_, _, name, err := core.CloudLogin(jwt)
	if err != nil {
		console.PrintError("Error trying to parse JWT token. Try to login again")
		log.Fatalf("cloud: login: %v", err)
	}

	console.PrintSuccess("Hi " + name + "!")
}

func (ctrl *CLIController) CloudLogout(cmd *cobra.Command, args []string) {
	fmt.Println(tools.Logo)
	core.CloudLogout()
	console.PrintSuccess("Successfully logged out from cbox cloud. See you back soon!")
}

func (ctrl *CLIController) CloudSpacePublish(cmd *cobra.Command, args []string) {
	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("cloud: publish space: %v", err)
	}

	cbox := core.LoadCbox("")

	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("cloud: publish space: %v", err)
	}

	tools.PrintSpace("Space", space)

	if console.Confirm("Publish?") {
		fmt.Printf("Publishing space '%s'...\n", space.Label)

		cloud, err := core.CloudClient()
		if err != nil {
			log.Fatalf("cloud: publish space: %v", err)
		}
		err = cloud.SpacePublish(space)
		if err != nil {
			log.Fatalf("cloud: publish space: %v", err)
		}

		console.PrintSuccess("Space published successfully!")
	} else {
		console.PrintError("Publish cancelled")
	}
}

func (ctrl *CLIController) CloudSpaceClone(cmd *cobra.Command, args []string) {
	selector, err := models.ParseSelectorForCloudCommand(args[0])
	if err != nil {
		log.Fatalf("cloud: clone space: invalid cloud selector: %v", err)
	}

	cloud, err := core.CloudClient()
	if err != nil {
		log.Fatalf("cloud: clone space: cloud client: %v", err)
	}

	space, err := cloud.SpaceRetrieve(selector, nil)
	if err != nil {
		log.Fatalf("cloud: clone space: %v", err)
	}

	tools.PrintSpace("Space to clone", space)

	if console.Confirm("Clone?") {
		cbox := core.LoadCbox("")

		err := cbox.SpaceCreate(space)
		for err != nil {
			console.PrintError("Space already found in your cbox. Try a different one")
			space.Label = strings.ToLower(console.ReadString("Label"))
			err = cbox.SpaceCreate(space)
		}

		core.PersistCbox(cbox)

		console.PrintSuccess("Space cloned successfully!")
	} else {
		console.PrintError("Clone cancelled")
	}
}

func (ctrl *CLIController) CloudSpacePull(cmd *cobra.Command, args []string) {
	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("cloud: pull space: invalid cloud selector: %v", err)
	}

	cloud, err := core.CloudClient()
	if err != nil {
		log.Fatalf("cloud: pull space: cloud client: %v", err)
	}

	cbox := core.LoadCbox("")

	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("cloud: pull space: %v", err)
	}

	spaceCloud, err := cloud.SpaceRetrieve(nil, &space.ID)
	if err != nil {
		log.Fatalf("cloud: pull space: %v", err)
	}

	// Note: Label is not overwritten because user can renamed his local copy of the space
	space.Entries = spaceCloud.Entries
	space.UpdatedAt = spaceCloud.UpdatedAt
	space.Description = spaceCloud.Description

	core.PersistCbox(cbox)

	tools.PrintSpace("Pulled space", space)

	console.PrintSuccess("Space pulled successfully!")
}

func (ctrl *CLIController) CloudCommandList(cmd *cobra.Command, args []string) {
	selector, err := models.ParseSelectorForCloudCommand(args[0])
	if err != nil {
		log.Fatalf("cloud: list commands: invalid cloud selector: %v", err)
	}

	cloud, err := core.CloudClient()
	if err != nil {
		log.Fatalf("cloud: list commands: cloud client: %v", err)
	}

	commands, err := cloud.CommandList(selector)
	if err != nil {
		log.Fatalf("cloud: list commands: %v", err)
	}

	tools.PrintCommandList(commands, viewSnippet, false)
}

func (ctrl *CLIController) CloudCommandCopy(cmd *cobra.Command, args []string) {
	cmdSelector, err := models.ParseSelectorForCloudCommand(args[0])
	if err != nil {
		log.Fatalf("cloud: copy command: invalid cloud selector: %v", err)
	}

	spaceSelector, err := models.ParseSelectorMandatorySpace(args[1])
	if err != nil {
		log.Fatalf("cloud: copy command: invalid space selector: %v", err)
	}

	cbox := core.LoadCbox("")

	space, err := cbox.SpaceFind(spaceSelector.Space)
	if err != nil {
		log.Fatalf("cloud: copy command: local space: %v", err)
	}

	cloud, err := core.CloudClient()
	if err != nil {
		log.Fatalf("cloud: copy command: cloud client: %v", err)
	}

	commands, err := cloud.CommandList(cmdSelector)
	if err != nil {
		log.Fatalf("cloud: copy command: retrieving matches: %v", err)
	}

	if len(commands) == 0 {
		console.PrintError(fmt.Sprintf("Command '%s' not found", cmdSelector))
	}

	tools.PrintCommandList(commands, false, false)
	fmt.Println()

	if console.Confirm(fmt.Sprintf("Copy these commands into %s?", spaceSelector)) {

		failures := false
		for _, command := range commands {
			err = space.CommandAdd(&command)
			if err != nil {
				failures = true
				log.Printf("cloud: copy command: %v", err)
			}
		}

		core.PersistCbox(cbox)

		if failures {
			console.PrintError("\nSome commands could not be stored")
		} else {
			console.PrintSuccess("Commands copied successfully!")
		}
	} else {
		console.PrintError("Copy cancelled")
	}
}
