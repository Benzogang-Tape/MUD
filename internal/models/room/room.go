package room

import (
	"strings"

	"github.com/Benzogang-Tape/MUD/internal/models/furniture"
)

type Routes []*furniture.Door

type Room struct {
	Name         string
	View         string
	Description  string
	Routes       Routes
	Furniture    furniture.Furniture
	ShowMissions bool
}

func (room Room) RouteExists(routeToFind string) (routeExists bool, doorState bool) {
	for _, door := range room.Routes {
		if _, ok := door.Paths[routeToFind]; ok {
			routeExists = true
			doorState = door.Opened
			break
		}
	}
	return routeExists, doorState
}

func (room Room) AvailableRoutes() (availableRoutes string) {
	routes := make([]string, 0, len(room.Routes))
	for _, route := range room.Routes {
		if way, ok := route.Paths[room.Name]; ok {
			routes = append(routes, way)
		}
	}
	return strings.Join(routes, ", ")
}

func (room Room) Environment() string {
	furn := room.Furniture
	if furn.IsEmpty() {
		return "пустая комната"
	}
	return furn.ShowContents()
}
