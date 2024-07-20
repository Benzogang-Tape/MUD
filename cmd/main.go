package main

import (
	"bufio"
	"fmt"
	"github.com/Benzogang-Tape/MUD/models"
	"github.com/Benzogang-Tape/MUD/utils"
	"os"
	"strings"
)

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
			defer fmt.Println("Данные не сохранены")
			break
		}
		fmt.Println(handleCommand(command))
	}
}

func initGame() {
	defaultPostfix := utils.Postfixer("ничего интересного")
	models.Wrld = models.World{
		"кухня": models.Room{
			RoomName:    "находишься на кухне",
			Description: defaultPostfix("кухня, "),
			Furniture: models.Furniture{
				{"на столе: ", "чай"},
			},
			Routes: models.Routes{
				{RoomName: "коридор", Door: true},
			},
		},
		"комната": models.Room{
			RoomName:    "в своей комнате",
			Description: utils.Prefixer("ты ")("в своей комнате"),
			Furniture: models.Furniture{
				{"на столе: ", "ключи", "конспекты"},
				{"на стуле: ", "рюкзак"},
			},
			Routes: models.Routes{
				{RoomName: "коридор", Door: true},
			},
		},
		"коридор": models.Room{
			RoomName:    "в коридоре",
			Description: defaultPostfix(""),
			Furniture:   models.Furniture{},
			Routes: models.Routes{
				{RoomName: "кухня", Door: true},
				{RoomName: "комната", Door: true},
				{RoomName: "улица", Door: false},
			},
		},
		"улица": models.Room{
			RoomName:    "на улице весна",
			Description: utils.Postfixer("")("на улице весна"),
			Furniture:   models.Furniture{},
			Routes: models.Routes{
				{RoomName: "домой", Door: true},
			},
		},
	}

	necessaryGear := models.Gear{
		{"рюкзак", "ключи", "конспекты"},
	}

	models.Plr = models.Player{
		NecessaryGear: &necessaryGear,
		Location:      "кухня",
		Gear:          models.Gear{},
		ToDoList: models.ToDoList{
			"собрать рюкзак",
			"идти в универ",
		},
		AvailableActions: map[string]models.Action{
			"осмотреться": models.Plr.LookAround,
			"идти":        models.Plr.GoTo,
			"надеть":      models.Plr.PutOn,
			"взять":       models.Plr.GetItem,
			"применить":   models.Plr.UseItem,
		},
	}
}

func handleCommand(command string) string {
	commands := strings.Split(command, " ")
	action, actionExists := models.Plr.AvailableActions[commands[0]]
	if actionExists {
		return action(commands[1:])
	}
	return "неизвестная команда"
}
