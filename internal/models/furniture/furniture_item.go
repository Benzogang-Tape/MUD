package furniture

import "github.com/Benzogang-Tape/MUD/internal/models/items"

type Item struct {
	Name           string
	Contents       map[string]*items.Item
	IterationOrder []string
}
