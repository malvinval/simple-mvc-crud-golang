package homecontroller

import (
	"net/http"
	"text/template"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("views/home/index.html")

	if err != nil {
		panic(err)
	}

	template.Execute(w, nil)
}
