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
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
