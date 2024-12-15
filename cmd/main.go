package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Benzogang-Tape/MUD/internal/models/furniture"
	"github.com/Benzogang-Tape/MUD/internal/models/items"
	"github.com/Benzogang-Tape/MUD/internal/models/models"
	"github.com/Benzogang-Tape/MUD/internal/models/player"
	"github.com/Benzogang-Tape/MUD/internal/models/room"
	"github.com/Benzogang-Tape/MUD/pkg/stringfmt"
)

var plr player.Player

func main() {
	initGame()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command, err := scanner.Text(), scanner.Err()
		if err != nil {
			fmt.Println(err)
		}
		if command == "exit" {
			fmt.Println("Выход из игры...")
			break
		}
		fmt.Println(handleCommand(command))
	}
}

func initGame() {
	defaultPostfix := stringfmt.Postfixer("ничего интересного")
	playerPrefix := stringfmt.Prefixer("ты ")
	streetDoor := &furniture.Door{
		Paths: map[string]string{
			"улица":   "домой",
			"коридор": "улица",
			"домой":   "коридор",
		},
		Opened: false,
	}
	kitchenDoor := &furniture.Door{
		Paths: map[string]string{
			"кухня":   "коридор",
			"коридор": "кухня",
		},
		Opened: true,
	}
	roomDoor := &furniture.Door{
		Paths: map[string]string{
			"комната": "коридор",
			"коридор": "комната",
		},
		Opened: true,
	}
	tea := &items.Item{
		Name:        "чай",
		IsNecessary: false,
		IsContainer: false,
		Target:      nil,
	}
	keys := &items.Item{
		Name:        "ключи",
		IsNecessary: true,
		IsContainer: false,
		Target:      streetDoor,
	}
	notes := &items.Item{
		Name:        "конспекты",
		IsNecessary: true,
		IsContainer: false,
		Target:      nil,
	}
	backpack := &items.Item{
		Name:        "рюкзак",
		IsNecessary: true,
		IsContainer: true,
		Target:      nil,
	}
	chair := &furniture.Item{
		Name: "на стуле",
		Contents: map[string]*items.Item{
			"рюкзак": backpack,
		},
		IterationOrder: []string{"рюкзак"},
	}
	table := &furniture.Item{
		Name: "на столе",
		Contents: map[string]*items.Item{
			"ключи":     keys,
			"конспекты": notes,
		},
		IterationOrder: []string{"ключи", "конспекты"},
	}
	kitchenTable := &furniture.Item{
		Name: "на столе",
		Contents: map[string]*items.Item{
			"чай": tea,
		},
		IterationOrder: []string{"чай"},
	}
	kitchen := &room.Room{
		Name:        "кухня",
		View:        playerPrefix("находишься на кухне"),
		Description: defaultPostfix("кухня, "),
		Furniture: furniture.Furniture{
			kitchenTable,
		},
		Routes: room.Routes{
			kitchenDoor,
		},
		ShowMissions: true,
	}
	livingRoom := &room.Room{
		Name:        "комната",
		View:        "",
		Description: playerPrefix("в своей комнате"),
		Furniture: furniture.Furniture{
			table,
			chair,
		},
		Routes: room.Routes{
			roomDoor,
		},
		ShowMissions: false,
	}
	hallway := &room.Room{
		Name:        "коридор",
		View:        playerPrefix("в коридоре"),
		Description: defaultPostfix(""),
		Furniture:   furniture.Furniture{},
		Routes: room.Routes{
			kitchenDoor,
			roomDoor,
			streetDoor,
		},
		ShowMissions: false,
	}
	outside := &room.Room{
		Name:        "улица",
		View:        playerPrefix("на улице"),
		Description: "на улице весна",
		Furniture:   furniture.Furniture{},
		Routes: room.Routes{
			streetDoor,
		},
		ShowMissions: false,
	}
	models.Wrld = models.World{
		"кухня":   kitchen,
		"комната": livingRoom,
		"коридор": hallway,
		"улица":   outside,
		"домой":   hallway,
	}

	plr = player.Player{
		Location:  kitchen,
		Container: nil,
		Gear:      models.Gear{},
		NecessaryGear: models.Gear{
			"ключи":     keys,
			"конспекты": notes,
		},
		TODOList: models.ToDoList{
			"собрать рюкзак",
			"идти в универ",
		},
		Stocked: false,
		AvailableActions: map[string]models.Action{
			"осмотреться": plr.LookAround,
			"идти":        plr.GoTo,
			"надеть":      plr.PutOn,
			"взять":       plr.GetItem,
			"применить":   plr.UseItem,
		},
	}
}

func handleCommand(command string) string {
	commands := strings.Split(command, " ")
	action, actionExists := plr.AvailableActions[commands[0]]
	if !actionExists {
		return "неизвестная команда"
	}
	return action(commands[1:])
}
