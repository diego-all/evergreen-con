package config

import (
	"database/sql"
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
