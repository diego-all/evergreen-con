package entities

import "time"

type Device struct {
	Id   int //uint despues
	Name string
	//Location string
	Location   Location
	Parameters string
	//Type       string
	Type      DeviceType
	Model     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
