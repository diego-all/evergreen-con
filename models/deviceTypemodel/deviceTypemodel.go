package deviceTypemodel

import (
	"evergreen-con/config"
	"evergreen-con/entities"
)

func GetAll() []entities.DeviceType {

	//rows, err := config.DB.Query("SELECT * FROM devices")
	rows, err := config.DB.Query(`SELECT * FROM device_types`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var deviceTypes []entities.DeviceType
	//var devices []entities.Device

	for rows.Next() {
		//var device entities.Device
		var deviceType entities.DeviceType
		if err := rows.Scan(
			&deviceType.Id,
			&deviceType.Name,
			&deviceType.CreatedAt,
			&deviceType.UpdatedAt,
		); err != nil {
			panic(err)
		}

		deviceTypes = append(deviceTypes, deviceType)
	}

	return deviceTypes
}
