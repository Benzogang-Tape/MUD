package models

import (
	"github.com/Benzogang-Tape/MUD/utils"
	"slices"
	"strings"
)

type Player struct {
	NecessaryGear    *Gear
	Location         string
	Gear             Gear
	ToDoList         ToDoList
	AvailableActions map[string]Action
}

func (player *Player) needToDo() string {
	return ", надо " + strings.Join(player.ToDoList, " и ")
}

func (player *Player) updateToDoList() {
	for containerIdx, container := range player.Gear {
		itemsToBeTaken, itemsTaken := len((*player.NecessaryGear)[containerIdx]), 0
		for _, item := range (*player.NecessaryGear)[containerIdx] {
			if slices.Contains(container, item) {
				itemsTaken++
			}
		}
		if itemsTaken == itemsToBeTaken {
			player.ToDoList = slices.Delete(player.ToDoList, containerIdx, containerIdx+1)
		}
	}
}

func (player *Player) showEnvironment() string {
	outputParts := make([]string, 0, 2)
	currentRoomFurniture := Wrld[player.Location].Furniture
	currentRoomFurniture = slices.DeleteFunc(currentRoomFurniture, func(furnitureItem []string) bool {
		return len(furnitureItem) <= 1
	})
	if len(currentRoomFurniture) == 0 {
		return "пустая комната"
	}
	for _, furnitureItem := range currentRoomFurniture {
		outputParts = append(outputParts, furnitureItem[0]+strings.Join(furnitureItem[1:], ", "))
	}
	return strings.Join(outputParts, ", ")
}

func (player *Player) canGoTo() string {
	routes := make([]string, 0, 3)
	currentRoutes := Wrld[player.Location].Routes
	for _, route := range currentRoutes {
		routes = append(routes, route.RoomName)
	}
	return "можно пройти - " + strings.Join(routes, ", ")
}

func (player *Player) LookAround([]string) string {
	var playerOutput string
	if player.Location != "комната" {
		playerOutput += utils.Prefixer("ты ")(Wrld[player.Location].RoomName)
	}
	environment := player.showEnvironment()

	if environment != "" && playerOutput != "" {
		playerOutput += ", "
	}
	playerOutput += environment
	if player.Location == "кухня" {
		playerOutput += player.needToDo()
	}
	answer := utils.FormatPrefixer(playerOutput)
	return answer(player.canGoTo())
}

func (player *Player) GoTo(args []string) string {
	goTo := args[0]
	if goTo == player.Location {
		return "ты уже здесь"
	}
	if routeExists, doorStatus := routeExists(Wrld[player.Location].Routes, goTo); !routeExists {
		return "нет пути в " + goTo
	} else if !doorStatus {
		return "дверь закрыта"
	}
	player.Location = goTo
	answer := utils.FormatPrefixer(Wrld[player.Location].Description)
	return answer(player.canGoTo())
}

func (player *Player) PutOn(args []string) string {
	container := args[0]
	currentRoomFurniture := Wrld[player.Location].Furniture
	for furnitureIdx, furnitureItem := range currentRoomFurniture {
		if slices.Contains(furnitureItem, container) {
			if !utils.ContainerExists(*player.NecessaryGear, container) {
				return "нельзя надеть " + container
			}
			if utils.ContainerExists(player.Gear, container) {
				return "у вас уже есть " + container
			}
			player.Gear = append(player.Gear, []string{container})
			containerIdx := slices.Index(Wrld[player.Location].Furniture[furnitureIdx], container)
			currentRoomFurniture[furnitureIdx] = slices.Delete(currentRoomFurniture[furnitureIdx], containerIdx, containerIdx+1)
			return "вы надели: " + container
		}
	}
	return "нет такого"
}

func (player *Player) GetItem(args []string) string {
	item := args[0]
	currentRoomFurniture := Wrld[player.Location].Furniture
	for furnitureIdx, furnitureContent := range currentRoomFurniture {
		if slices.Contains(furnitureContent, item) {
			if len(player.Gear) == 0 || len(player.Gear[0]) == 0 {
				return "некуда класть"
			}
			for containerIdx, containerContent := range *player.NecessaryGear {
				if slices.Contains(containerContent, item) && slices.Index(containerContent, item) != 0 {

					itemIdx := slices.Index(currentRoomFurniture[furnitureIdx], item)
					player.Gear[containerIdx] = append(player.Gear[containerIdx], item)
					currentRoomFurniture[furnitureIdx] = slices.Delete(currentRoomFurniture[furnitureIdx], itemIdx, itemIdx+1)
					player.updateToDoList()
					return "предмет добавлен в инвентарь: " + item
				}
			}
			return "вам не нужен предмет: " + item
		}
	}
	return "нет такого"
}

func (player *Player) UseItem(args []string) string {
	item, target, destination := args[0], args[1], "улица"
	if len(args) > 2 {
		destination = args[2]
	}
	if doorExists, _ := routeExists(Wrld[player.Location].Routes, destination); !doorExists {
		return "не к чему применить"
	}
	currentDoorIdx := findRouteIndexByName(Wrld[player.Location].Routes, destination)
	for _, containerContent := range player.Gear {
		if slices.Contains(containerContent, item) {
			if item != "ключи" {
				return "нельзя применить " + item + " к предмету: " + target
			}
			if target != "дверь" {
				return "не к чему применить"
			}
			Wrld[player.Location].Routes[currentDoorIdx].Door = !Wrld[player.Location].Routes[currentDoorIdx].Door
			if Wrld[player.Location].Routes[currentDoorIdx].Door {
				return "дверь открыта"
			}
			return "дверь закрыта"
		}
	}
	return "нет предмета в инвентаре - " + item
}
