package config

import (
	"database/sql"
	"evergreen-con/entities"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const maxOpenDbConn = 5               // numero maximo de conexiones permitidas abiertas
const maxIdleDbConn = 5               // numero maximo de conexiones inactivas (ociosas) Abiertas y disponibles para reutilizacion
const maxDbLifeTime = 5 * time.Minute // tiempo antes de que se considere inactiva una conexion

func ConnectDatabase(dsn string) (*sql.DB, error) {

	// export DSN="root:password@tcp(127.0.0.1:33090)/evergreen_con?parseTime=true"
	//db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:33090)/evergreen_con?parseTime=true")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	// Very importart to set the global variable DB
	DB = db

	log.Println("Database connected")

	return DB, nil
}

func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		fmt.Println("Error!", err)
		return err
	}

	log.Println("*** Pinged database successfully! ***")

	testQuery(db)

	return nil
}

func testQuery(db *sql.DB) {

	var result string

	err := db.QueryRow("SELECT ?", "*** Conexio Est ... Habemus conexión ***").Scan(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("From testQuery", result)
}

func QueryApps(db *sql.DB) {

	//rows, err := db.Query(`select * from applications2`)
	rows, err := db.Query(`
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
	//var applications2 []entities.Application2

	for rows.Next() {
		var application entities.Application
		//var application2 entities.Application2
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
			// &application2.Id,
			// &application2.Identifier,
			// &application2.Name,
			// &application2.Port,
			// &application2.Status,
			// &application2.Type,
			// &application2.Language,
			// &application2.DeviceId,
			// &application2.CreatedAt,
			// &application2.UpdatedAt,
		); err != nil {
			panic(err)
		}

		applications = append(applications, application)
		//applications2 = append(applications2, application2)
	}

	//fmt.Println("From queryApps", applications2)
	fmt.Println("From queryApps", applications)

	// 	rows, err := db.Query(`
	// 	SELECT
	// 		devices.id,
	// 		devices.name,
	// 		devices.location,
	// 		devices.parameters,
	// 		device_types.name as device_type_name,
	// 		devices.model,
	// 		devices.created_at,
	// 		devices.updated_at FROM devices
	// 	JOIN device_types ON devices.type_id = device_types.id;
	// `)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	defer rows.Close()

	// 	var devices []entities.Device

	// 	for rows.Next() {
	// 		var device entities.Device
	// 		if err := rows.Scan(
	// 			&device.Id,
	// 			&device.Name,
	// 			&device.Location,
	// 			&device.Parameters,
	// 			&device.Type.Name,
	// 			&device.Model,
	// 			&device.CreatedAt,
	// 			&device.UpdatedAt,
	// 		); err != nil {
	// 			panic(err)
	// 		}

	// 		devices = append(devices, device)
	// 	}

	// 	fmt.Println("From queryApps", devices)

}

type Application struct {
	Id         int
	Identifier string
	Name       string
	Port       string
	Status     string
	Type       string
	Language   string
	DeviceName string
	CreatedAt  string
	UpdatedAt  string
}

func QueryApplications(db *sql.DB) ([]Application, error) {
	query := `
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
		    applications.updated_at
		FROM applications
		JOIN devices ON applications.device_id = devices.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applications []Application

	for rows.Next() {
		var application Application
		if err := rows.Scan(
			&application.Id,
			&application.Identifier,
			&application.Name,
			&application.Port,
			&application.Status,
			&application.Type,
			&application.Language,
			&application.DeviceName,
			&application.CreatedAt,
			&application.UpdatedAt,
		); err != nil {
			return nil, err
		}

		applications = append(applications, application)
	}

	fmt.Println(applications)
	return applications, nil
}
