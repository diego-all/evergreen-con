package main

import (
	"evergreen-con/config"
	"evergreen-con/controllers/applicationcontroller"
	"evergreen-con/controllers/devicecontroller"
	"evergreen-con/controllers/homecontroller"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	//dsn := os.Getenv("DSN")
	var dsn string

	// instrucción condicional inline (if/else inline)
	if dsn = os.Getenv("DSN"); dsn == "" {
		fmt.Println("La variable de entorno DSN no está definida.")
	} else {
		fmt.Printf("El valor DSN es: %s\n", dsn)
	}

	// 	Database connection
	config.ConnectDatabase(dsn)

	// 1.Home
	http.HandleFunc("/", homecontroller.Home)

	// 2. Devices
	http.HandleFunc("/devices/add", devicecontroller.Create)
	http.HandleFunc("/devices/detail", devicecontroller.Read)
	http.HandleFunc("/devices/edit", devicecontroller.Update)
	http.HandleFunc("/devices", devicecontroller.List)
	http.HandleFunc("/devices/delete", devicecontroller.Delete)

	// 3. Applications
	http.HandleFunc("/applications/add", applicationcontroller.Create)
	http.HandleFunc("/applications/detail", applicationcontroller.Read)
	http.HandleFunc("/applications/edit", applicationcontroller.Update)
	http.HandleFunc("/applications", applicationcontroller.List)
	http.HandleFunc("/applications/delete", applicationcontroller.Delete)

	// Run server
	log.Println("Server running on port: 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))

}
