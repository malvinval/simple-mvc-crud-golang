package routes

import (
	"gocrud/controllers/categorycontroller"
	"gocrud/controllers/homecontroller"
	"net/http"
)

func Init() {
	http.HandleFunc("/", homecontroller.WelcomeHandler)

	http.HandleFunc("/categories", categorycontroller.ReadHandler)
	http.HandleFunc("/categories/create", categorycontroller.StoreHandler)
	http.HandleFunc("/categories/edit", categorycontroller.UpdateHandler)
	http.HandleFunc("/categories/delete", categorycontroller.DeleteHandler)
}
