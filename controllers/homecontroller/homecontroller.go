package homecontroller

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func Controls(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/controls.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
