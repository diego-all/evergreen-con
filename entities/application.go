package entities

import "time"

type Application struct {
	Id         int
	Identifier string
	Name       string
	Port       string
	Status     string
	Type       string
	Language   string
	Device     Device
	//Device    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Application2 struct {
	Id         int
	Identifier string
	Name       string
	Port       string
	Status     string
	Type       string
	Language   string
	DeviceId   int
	CreatedAt  string
	UpdatedAt  string
}
