package entities

import "time"

type Location struct {
	Id         uint //uint despues
	Identifier string
	Name       string
	Project    string //Production Project/Agronomic Project
	Plot       string //Plot/Block/Parcela
	Property   string //Property/Predio
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
