package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/dpecos/cbox/models"
	homedir "github.com/mitchellh/go-homedir"
)

const DEFAULT_SPACE_NAME = "default"

func resolveSpaceFile(spaceName string) string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	return home + "/.cbox/spaces/" + spaceName + ".json"
}

func SpaceStore(name string, title string) models.Space {
	space := models.Space{
		Name:  name,
		Title: title,
	}

	raw, err := json.MarshalIndent(space, "", "  ")
	if err != nil {
		log.Fatalf("Error generating json for space %s: %s", name, err)
	}

	err = ioutil.WriteFile(resolveSpaceFile(name), raw, os.ModeExclusive)
	if err != nil {
		log.Fatalf("Error writing json file for space %s: %s", name, err)
	}

	return space
}

func SpaceDelete(space models.Space) {
	err := os.Remove(resolveSpaceFile(space.Name))
	if err != nil {
		log.Fatalf("Error deleting space %s: %s", space.Name, err)
	}
}

func SpacesList() []string {
	return []string{DEFAULT_SPACE_NAME}
}

func SpaceLoad(name string) models.Space {
	raw, err := ioutil.ReadFile(resolveSpaceFile(name))
	if err != nil {
		log.Fatalf("Error loading json file for space %s: %s", name, err)
	}

	var space models.Space
	err = json.Unmarshal(raw, &space)
	if err != nil {
		log.Fatalf("Error parsing json file for space %s: %s", name, err)
	}

	return space
}
