package locationcontroller

import (
	"evergreen-con/models/locationmodel"
	"evergreen-con/production"
	"html/template"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	locations := locationmodel.Getall()
	data := map[string]any{
		"locations": locations,
	}

	temp, err := template.ParseFiles("views/location/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func SynLocations(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/controls.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)

	production.ReadJSONmock()
}

// func ResetLocation(w http.ResponseWriter, r *http.Request) {
// 	temp, err := template.ParseFiles("views/home/controls.html")
// 	if err != nil {
// 		panic(err)
// 	}

// 	temp.Execute(w, nil)

// 	production.ResetLocation()
// }
