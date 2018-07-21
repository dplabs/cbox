package tests

import (
	"math/rand"
	"testing"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	uuid "github.com/satori/go.uuid"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func findSpaceFile(space *models.Space) bool {
	spaces := core.SpaceListFiles()

	found := false
	for _, s := range spaces {
		if s.ID == space.ID {
			found = true
			break
		}
	}
	return found
}

func assertSpaceFileExists(t *testing.T, space *models.Space) {
	found := findSpaceFile(space)
	if !found {
		t.Fatal("space file could not be found (and should)")
	}
}

func assertSpaceFileNotExists(t *testing.T, space *models.Space) {
	found := findSpaceFile(space)
	if found {
		t.Fatal("new space found (and shouldn't)")
	}
}

func createSpace(t *testing.T) *models.Space {
	if cbox == nil {
		t.Fatal("cbox not initialized")
	}

	id, _ := uuid.NewV4()
	space := models.Space{
		Label:       randString(8),
		Description: randString(15),
	}
	space.ID = id

	err := cbox.SpaceAdd(&space)

	if err != nil {
		t.Error(err)
	}

	s, _ := cbox.SpaceFind(space.ID.String())
	return s
}