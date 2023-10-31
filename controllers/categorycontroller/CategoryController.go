package categorycontroller

import (
	"gocrud/entities"
	"gocrud/models/categorymodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()

	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/categories/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func StoreHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/categories/create.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Store(category); !ok {
			temp, _ := template.ParseFiles("views/categories/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		categories := categorymodel.Get(r.URL.Query().Get("id"))

		data := map[string]any{
			"categories": categories,
		}

		temp, err := template.ParseFiles("views/categories/edit.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var category entities.Category

		id, _ := strconv.Atoi(r.FormValue("id"))

		category.Id = uint(id)
		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		categorymodel.Update(category)

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var category entities.Category

		id, _ := strconv.Atoi(r.FormValue("id"))

		category.Id = uint(id)

		categorymodel.Destroy(category.Id)

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}
