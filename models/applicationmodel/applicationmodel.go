package applicationmodel

import (
	"evergreen-con/config"
	"evergreen-con/entities"
)

func Getall() []entities.Application {
	//rows, err := config.DB.Query("SELECT * FROM applications")
	rows, err := config.DB.Query(`
		SELECT
		    applications.id,
		    applications.identifier,
		    applications.name,
		    applications.port,
			applications.status,
			applications.type,
			applications.language,
		    devices.name as device_name,
			applications.created_at,
			applications.updated_at FROM applications
		JOIN devices ON applications.device_id = devices.id;
	`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var applications []entities.Application

	for rows.Next() {
		var application entities.Application
		if err := rows.Scan(
			&application.Id,
			&application.Identifier,
			&application.Name,
			&application.Port,
			&application.Status,
			&application.Type,
			&application.Language,
			&application.Device.Name,
			&application.CreatedAt,
			&application.UpdatedAt,
		); err != nil {
			panic(err)
		}

		applications = append(applications, application)
	}

	return applications
}

func Create(applications entities.Application) bool {
	result, err := config.DB.Exec(`
		INSERT INTO applications(
			identifier, name, port, status, type, language, device_id, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		applications.Identifier,
		applications.Name,
		applications.Port,
		applications.Status,
		applications.Type,
		applications.Language,
		applications.Device.Id,
		applications.CreatedAt,
		applications.UpdatedAt,
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

func Detail(id int) entities.Application {

	//row := config.DB.QueryRow(`SELECT id, identifier, name, port, status, type, language, device_id, created_at, updated_at FROM applications WHERE id = ? `, id)
	row := config.DB.QueryRow(`
		SELECT
		    applications.id,
		    applications.identifier,
		    applications.name,
		    applications.port,
			applications.status,
			applications.type,
			applications.language,
		    devices.name as device_name,
			applications.created_at,
			applications.updated_at FROM applications
		JOIN devices ON applications.device_id = devices.id
		WHERE applications.id = ?
	`, id)

	var application entities.Application

	if err := row.Scan(
		&application.Id,
		&application.Identifier,
		&application.Name,
		&application.Port,
		&application.Status,
		&application.Type,
		&application.Language,
		&application.Device.Name,
		&application.CreatedAt,
		&application.UpdatedAt); err != nil {
		panic(err.Error())
	}

	return application
}

func Update(id int, application entities.Application) bool {
	query, err := config.DB.Exec(`
		UPDATE applications SET 
			identifier = ?,
			name = ?, 
			port = ?,
			status = ?,
			type = ?,
			language = ?,
			device_id = ?,
			updated_at = ?
		WHERE id = ?`,
		application.Identifier,
		application.Name,
		application.Port,
		application.Status,
		application.Type,
		application.Language,
		application.Device.Id,
		application.UpdatedAt,
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
	_, err := config.DB.Exec("DELETE FROM applications WHERE id = ?", id)
	return err
}
