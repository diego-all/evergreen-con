package entities

import "time"

type Device struct {
	Id         int
	Name       string
	Location   string
	Parameters string
	//Type       string
	Type      DeviceType
	Model     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
