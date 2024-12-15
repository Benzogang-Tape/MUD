package furniture

import (
	"fmt"

	"github.com/Benzogang-Tape/MUD/internal/models/items"
)

type Door struct {
	Paths  map[string]string
	Opened bool
}

func (door *Door) Trigger(trigger *items.Item) string {
	if trigger.Name != "ключи" {
		return fmt.Sprintf("%s %s %s %s", "нельзя применить", trigger.Name, "к предмету:", door.Name())
	}
	door.Opened = !door.Opened
	if door.Opened {
		return "дверь открыта"
	}
	return "дверь закрыта"
}

func (door *Door) Name() string {
	return "дверь"
}
