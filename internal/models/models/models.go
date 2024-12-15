package models

import (
	"strings"

	"github.com/Benzogang-Tape/MUD/internal/models/items"
	"github.com/Benzogang-Tape/MUD/internal/models/room"
)

var Wrld World

type World map[string]*room.Room
type ToDoList []string
type Gear map[string]*items.Item
type Action func([]string) string

func (list ToDoList) TODO() string {
	return "надо " + strings.Join(list, " и ")
}
