package devicemodel

import (
	"evergreen-con/config"
	"evergreen-con/entities"
)

func GetAll() []entities.Device {

	//rows, err := config.DB.Query("SELECT * FROM devices")
	rows, err := config.DB.Query(`
		SELECT
		    devices.id,
		    devices.name,
			locations.identifier as location_identifier,
		    devices.parameters,
		    device_types.name as device_type_name,
		    devices.model,
		    devices.created_at,
		    devices.updated_at FROM devices
		JOIN device_types ON devices.type_id = device_types.id
		JOIN locations ON devices.location_id = locations.id;
	`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var devices []entities.Device

	for rows.Next() {
		var device entities.Device
		if err := rows.Scan(
			&device.Id,
			&device.Name,
			&device.Location.Identifier,
			&device.Parameters,
			&device.Type.Name,
			&device.Model,
			&device.CreatedAt,
			&device.UpdatedAt,
		); err != nil {
			panic(err)
		}

		devices = append(devices, device)
	}

	return devices
}

func Create(device entities.Device) bool {
	result, err := config.DB.Exec(`
		INSERT INTO devices (name, location_id, parameters, type_id, model, created_at, updated_at) 
		VALUE (?, ?, ?, ?, ?, ?, ?)`,
		device.Name,
		device.Location.Id,
		device.Parameters,
		device.Type.Id,
		device.Model,
		device.CreatedAt,
		device.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Device {
	//row := config.DB.QueryRow(`SELECT id, name, location, parameters, type_id, model FROM devices WHERE id = ? `, id)

	// row := config.DB.QueryRow(`
	// 	SELECT
	// 		devices.id,
	// 		devices.name,
	// 		devices.location_id,
	// 		devices.parameters,
	// 		device_types.name as device_type_name,
	// 		devices.model,
	// 		devices.created_at,
	// 		devices.updated_at FROM devices
	// 	JOIN device_types ON devices.type_id = device_types.id
	// 	WHERE devices.id = ?
	// `, id)

	row := config.DB.QueryRow(`
		SELECT
			devices.id,
			devices.name,
			locations.identifier as location_identifier,
			devices.parameters,
			device_types.name as device_type_name,
			devices.model,
			devices.created_at,
			devices.updated_at FROM devices
		JOIN device_types ON devices.type_id = device_types.id
		JOIN locations ON devices.location_id = locations.id
		WHERE devices.id = ?
	 `, id)

	var device entities.Device

	if err := row.Scan(
		&device.Id, &device.Name, &device.Location.Identifier, &device.Parameters, &device.Type.Name, &device.Model, &device.CreatedAt, &device.UpdatedAt); err != nil {
		panic(err.Error())
	}

	return device
}

func Update(id int, device entities.Device) bool {
	// le hace falta editar los otros campos
	query, err := config.DB.Exec(`
		UPDATE devices SET 
			name = ?,
			location_id = ?,
			parameters = ?,
			type_id = ?,
			model = ?,
			updated_at = ? 
		WHERE id = ?`,
		device.Name,
		device.Location.Id,
		device.Parameters,
		device.Type.Id,
		device.Model,
		device.UpdatedAt,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM devices WHERE id = ?", id)
	return err
}
