package models

var Wrld World
var Plr Player

type ToDoList []string
type Routes []RoomInfo
type Gear [][]string
type Furniture [][]string
type Action func([]string) string

type RoomInfo struct {
	RoomName string
	Door     bool
}

type Room struct {
	RoomName    string
	Description string
	Routes      Routes
	Furniture   Furniture
}

type World map[string]Room

func routeExists(routes []RoomInfo, routeToFind string) (bool, bool) {
	var flag, door bool
	for _, route := range routes {
		if route.RoomName == routeToFind {
			flag = true
			door = route.Door
			break
		}
	}
	return flag, door
}

func findRouteIndexByName(routes []RoomInfo, routeToFind string) int {
	for routeIdx, route := range routes {
		if route.RoomName == routeToFind {
			return routeIdx
		}
	}
	return -1
}
