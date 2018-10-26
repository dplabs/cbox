package models

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type CBox struct {
	path   string
	Spaces []Space
}

func (cbox *CBox) spaceFindPositionByLabel(spaceLabel string) (int, error) {
	if spaceLabel == "" {
		return -1, fmt.Errorf("could not search by empty label")
	}
	for i, space := range cbox.Spaces {
		if space.Label == spaceLabel {
			return i, nil
		}
	}
	return -1, fmt.Errorf("space with label '%s' not found", spaceLabel)
}

func (cbox *CBox) spaceFindPositionByID(spaceID uuid.UUID) (int, error) {
	if spaceID == uuid.Nil {
		return -1, fmt.Errorf("could not search by empty ID")
	}
	for i, space := range cbox.Spaces {
		if space.ID == spaceID {
			return i, nil
		}
	}
	return -1, fmt.Errorf("space with ID '%s' not found", spaceID)
}

func (cbox *CBox) SpaceLabels() []string {
	labels := make(map[string]bool)
	for _, space := range cbox.Spaces {
		if _, ok := labels[space.Label]; !ok {
			labels[space.Label] = true
		}
	}
	result := []string{}
	for k, _ := range labels {
		result = append(result, k)
	}
	return result
}

func (cbox *CBox) SpaceFind(spaceLocator string) (*Space, error) {
	pos, err := cbox.spaceFindPositionByLabel(spaceLocator)
	if err != nil {
		id, e := uuid.FromString(spaceLocator)
		if e != nil {
			return nil, fmt.Errorf("find space (uuid): %v", err)
		}
		pos, err = cbox.spaceFindPositionByID(id)
		if err != nil {
			return nil, fmt.Errorf("find space: %v", err)
		}
	}
	return &cbox.Spaces[pos], nil
}

func (cbox *CBox) SpaceCreate(space *Space) error {
	s, err := cbox.SpaceFind(space.Label)
	if err == nil && s != nil {
		return fmt.Errorf("space create: space with label '%s' already exists", space.Label)
	}
	if space.Entries == nil {
		space.Entries = []Command{}
	}
	if space.CreatedAt == NilUnixTime {
		now := UnixTimeNow()
		space.CreatedAt = now
		space.UpdatedAt = now
	}
	cbox.Spaces = append(cbox.Spaces, *space)
	return nil
}

func (cbox *CBox) SpaceDestroy(space *Space) error {
	pos, err := cbox.spaceFindPositionByID(space.ID)
	if err != nil {
		return fmt.Errorf("space destroy: could not found space with ID '%s'", space.ID)
	}

	cbox.Spaces = append(cbox.Spaces[:pos], cbox.Spaces[pos+1:]...)

	return nil
}

func (cbox *CBox) SpaceEdit(space *Space, previousLabel string) error {

	if space.Label != previousLabel && len(cbox.SpaceLabels()) != len(cbox.Spaces) {
		return fmt.Errorf("space edit: duplicate label '%s", space.Label)
	}

	space.UpdatedAt = UnixTimeNow()

	return nil
}