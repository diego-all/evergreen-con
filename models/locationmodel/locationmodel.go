package locationmodel

import (
	"evergreen-con/config"
	"evergreen-con/entities"
)

func Getall() []entities.Location {
	rows, err := config.DB.Query("SELECT * FROM locations")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var locations []entities.Location

	for rows.Next() {
		var location entities.Location
		if err := rows.Scan(
			&location.Id,
			&location.Identifier,
			&location.Name,
			&location.Project,
			&location.Plot,
			&location.Property,
			&location.CreatedAt,
			&location.UpdatedAt,
		); err != nil {
			panic(err)
		}

		locations = append(locations, location)
	}

	return locations
}
