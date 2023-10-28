package applicationcontroller

import (
	"evergreen-con/entities"
	"evergreen-con/models/applicationmodel"
	"evergreen-con/models/devicemodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func List(w http.ResponseWriter, r *http.Request) {
	applications := applicationmodel.Getall()
	data := map[string]any{
		"applications": applications,
	}

	temp, err := template.ParseFiles("views/application/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/application/create.html")
		if err != nil {
			panic(err)
		}
		devices := devicemodel.GetAll()
		data := map[string]any{
			"devices": devices,
		}

		temp.Execute(w, data)

	}

	if r.Method == "POST" {
		var application entities.Application

		deviceId, err := strconv.Atoi(r.FormValue("device_id"))
		if err != nil {
			panic(err)
		}

		//device.Type.Id = uint(deviceTypeId)
		application.Identifier = r.FormValue("identifier")
		application.Name = r.FormValue("name")
		application.Port = r.FormValue("port")
		application.Status = r.FormValue("status")
		application.Type = r.FormValue("type")
		application.Language = r.FormValue("language")
		application.Device.Id = deviceId
		application.CreatedAt = time.Now()
		application.UpdatedAt = time.Now()

		if ok := applicationmodel.Create(application); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/applications", http.StatusSeeOther)
	}
}

func Read(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	application := applicationmodel.Detail(id)
	data := map[string]any{
		"application": application,
	}

	temp, err := template.ParseFiles("views/application/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/application/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		application := applicationmodel.Detail(id)
		devices := devicemodel.GetAll()
		data := map[string]any{
			"application": application,
			"devices":     devices,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var application entities.Application

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		deviceId, err := strconv.Atoi(r.FormValue("device_id"))
		if err != nil {
			panic(err)
		}

		application.Identifier = r.FormValue("identifier")
		application.Name = r.FormValue("name")
		application.Port = r.FormValue("port")
		application.Status = r.FormValue("status")
		application.Type = r.FormValue("type")
		application.Language = r.FormValue("language")
		application.Device.Id = deviceId
		application.UpdatedAt = time.Now()

		if ok := applicationmodel.Update(id, application); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/applications", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := applicationmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/applications", http.StatusSeeOther)
}
