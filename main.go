package main

import (
	"evergreen-con/config"
	"evergreen-con/controllers/applicationcontroller"
	"evergreen-con/controllers/devicecontroller"
	"evergreen-con/controllers/homecontroller"
	"evergreen-con/controllers/locationcontroller"
	"evergreen-con/production"
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

	//config.QueryApps(config.DB)
	//production.ReadJSONmock()

	http.HandleFunc("/locations/sync", func(w http.ResponseWriter, r *http.Request) {
		// Ejecuta la función de sincronización.
		production.ReadJSONmock()

		// Devuelve una respuesta al usuario.
		message := "Synchronization was successful."
		fmt.Fprintf(w, message)
	})

	// http.HandleFunc("/locations/reset", func(w http.ResponseWriter, r *http.Request) {
	// 	// Ejecuta la función de sincronización.
	// 	production.ResetLocation()

	// 	// Devuelve una respuesta al usuario.
	// 	message := "The locations was successfully cleared."
	// 	fmt.Fprintf(w, message)
	// })

	http.HandleFunc("/configurations/backup", func(w http.ResponseWriter, r *http.Request) {
		// Ejecuta la función de sincronización.
		production.BackupDatabase()

		// Devuelve una respuesta al usuario.
		message := "Backup Completed Successfully."
		fmt.Fprintf(w, message)
	})

	http.HandleFunc("/configurations/status", func(w http.ResponseWriter, r *http.Request) {
		// Ejecuta la función de sincronización.

		// Devuelve una respuesta al usuario.
		message := "Backup Completed Successfully."
		fmt.Fprintf(w, message)
	})

	// 1.Home
	http.HandleFunc("/", homecontroller.Home)
	http.HandleFunc("/controls", homecontroller.Controls)

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

	// 4. Locations
	http.HandleFunc("/locations", locationcontroller.List)

	// Run server
	log.Println("Server running on port: 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))

}
