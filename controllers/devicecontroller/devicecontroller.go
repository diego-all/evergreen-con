package devicecontroller

import (
	"evergreen-con/entities"
	"evergreen-con/models/deviceTypemodel"
	"evergreen-con/models/devicemodel"
	"strconv"
	"time"

	"net/http"
	"text/template"
)

func List(w http.ResponseWriter, r *http.Request) {
	devices := devicemodel.GetAll()
	data := map[string]any{
		"devices": devices,
	}

	temp, err := template.ParseFiles("views/device/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/device/create.html")
		if err != nil {
			panic(err)
		}

		deviceTypes := deviceTypemodel.GetAll()
		data := map[string]any{
			"deviceTypes": deviceTypes,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var device entities.Device

		deviceTypeId, err := strconv.Atoi(r.FormValue("device_types_id"))
		if err != nil {
			panic(err)
		}

		device.Name = r.FormValue("name")
		device.Location = r.FormValue("location")
		device.Parameters = r.FormValue("parameters")
		device.Type.Id = uint(deviceTypeId)
		device.Model = r.FormValue("model")

		device.CreatedAt = time.Now()
		device.UpdatedAt = time.Now()

		ok := devicemodel.Create(device)
		if !ok {
			temp, _ := template.ParseFiles("views/device/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/devices", http.StatusSeeOther)
	}
}

func Read(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	device := devicemodel.Detail(id)
	data := map[string]any{
		"device": device,
	}

	temp, err := template.ParseFiles("views/device/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/device/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		device := devicemodel.Detail(id)
		deviceTypes := deviceTypemodel.GetAll()
		data := map[string]any{
			"device":      device,
			"deviceTypes": deviceTypes,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var device entities.Device

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		deviceTypeId, err := strconv.Atoi(r.FormValue("device_types_id"))
		if err != nil {
			panic(err)
		}

		device.Name = r.FormValue("name")
		device.Location = r.FormValue("location")
		device.Parameters = r.FormValue("parameters")
		//device.Type = r.FormValue("type")
		// falta terminarlo como con Add
		device.Type.Id = uint(deviceTypeId)
		device.Model = r.FormValue("model")
		device.UpdatedAt = time.Now()

		if ok := devicemodel.Update(id, device); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/devices", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := devicemodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/devices", http.StatusSeeOther)
}
