package player

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Benzogang-Tape/MUD/internal/models/furniture"
	"github.com/Benzogang-Tape/MUD/internal/models/items"
	"github.com/Benzogang-Tape/MUD/internal/models/models"
	"github.com/Benzogang-Tape/MUD/internal/models/room"
	"github.com/Benzogang-Tape/MUD/pkg/stringfmt"
)

type Player struct {
	Location         *room.Room
	Container        *items.Item
	Gear             models.Gear
	NecessaryGear    models.Gear
	TODOList         models.ToDoList
	Stocked          bool
	AvailableActions map[string]models.Action
}

const wrongCommandUsage = "неправильное использование команды"

func (player *Player) currentRoom() *room.Room {
	return player.Location
}

func (player *Player) currentFurniture() furniture.Furniture {
	return player.currentRoom().Furniture
}

func (player *Player) needTODO() string {
	return ", " + player.TODOList.TODO()
}

func (player *Player) updateTODOList() {
	if player.Stocked {
		return
	}
	itemsToBeTaken, itemsTaken := len(player.NecessaryGear), 0
	for itemName, mustHaveItem := range player.NecessaryGear {
		if item, hasItem := player.Gear[itemName]; hasItem && item == mustHaveItem {
			itemsTaken++
		}
	}
	if itemsTaken == itemsToBeTaken {
		player.Stocked = true
		player.TODOList = player.TODOList[1:]
	}
}

func (player *Player) canGoTo() string {
	return "можно пройти - " + player.currentRoom().AvailableRoutes()
}

func (player *Player) LookAround([]string) string {
	var b strings.Builder
	currentRoom := player.currentRoom()
	b.WriteString(currentRoom.View)
	environment := currentRoom.Environment()

	if b.String() != "" {
		b.WriteString(", ")
	}
	b.WriteString(environment)
	if player.currentRoom().ShowMissions {
		b.WriteString(player.needTODO())
	}
	answer := stringfmt.FormatPrefixer(b.String())
	return answer(player.canGoTo())
}

func (player *Player) GoTo(args []string) string {
	if args == nil || len(args) != 1 {
		return wrongCommandUsage
	}
	goTo := args[0]
	if goTo == player.currentRoom().Name {
		return "ты уже здесь"
	}
	routeExists, isOpen := player.currentRoom().RouteExists(goTo)
	if !routeExists {
		return "нет пути в " + goTo
	}
	if !isOpen {
		return "дверь закрыта"
	}
	player.Location = models.Wrld[goTo]
	answer := stringfmt.FormatPrefixer(player.currentRoom().Description)
	return answer(player.canGoTo())
}

func (player *Player) PutOn(args []string) string {
	if args == nil || len(args) != 1 {
		return wrongCommandUsage
	}
	container := args[0]
	for _, furnitureItem := range player.currentFurniture() {
		if item, itemExists := furnitureItem.Contents[container]; itemExists {
			if !item.IsContainer {
				return "нельзя надеть " + container
			}
			if player.Container != nil {
				return "у вас уже есть " + container
			}
			player.Container = item
			delete(furnitureItem.Contents, container)
			containerIdx := slices.Index(furnitureItem.IterationOrder, container)
			furnitureItem.IterationOrder = slices.Delete(furnitureItem.IterationOrder, containerIdx, containerIdx+1)
			return "вы надели: " + container
		}
	}
	return "нет такого"
}

func (player *Player) GetItem(args []string) string {
	if args == nil || len(args) != 1 {
		return wrongCommandUsage
	}
	itemToGet := args[0]
	for _, furnitureItem := range player.currentFurniture() {
		if item, itemExists := furnitureItem.Contents[itemToGet]; itemExists {
			if player.Container == nil {
				return "некуда класть"
			}
			if !item.IsNecessary {
				return "вам не нужен предмет: " + itemToGet
			}
			player.Gear[itemToGet] = item
			delete(furnitureItem.Contents, itemToGet)
			itemIdx := slices.Index(furnitureItem.IterationOrder, itemToGet)
			furnitureItem.IterationOrder = slices.Delete(furnitureItem.IterationOrder, itemIdx, itemIdx+1)
			player.updateTODOList()
			return "предмет добавлен в инвентарь: " + itemToGet
		}
	}
	return "нет такого"
}

func (player *Player) UseItem(args []string) string {
	if args == nil || len(args) != 2 {
		return wrongCommandUsage
	}
	itemToUse, target := args[0], args[1]
	item, itemExists := player.Gear[itemToUse]
	if !itemExists {
		return "нет предмета в инвентаре - " + itemToUse
	}
	switch door := item.Target.(type) {
	case *furniture.Door:
		if _, nearTheDoor := door.Paths[player.currentRoom().Name]; !nearTheDoor {
			return "подойди ближе"
		}
		return item.Use(target)
	default:
		return fmt.Sprintf("%s %s %s %s", "нельзя применить", itemToUse, "к предмету:", target)
	}
}
