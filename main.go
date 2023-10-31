package main

import (
	"gocrud/config"
	"gocrud/routes"
)

func main() {
	// eksekusi function ConnectDB dalam package config yang kita buat sendiri
	config.ConnectDB("del", "2409", "crud-golang")

	// eksekusi function Init dalam package routes yang kita buat sendiri
	routes.Init()

	// eksekusi function CreateServer dalam package config yang kita buat sendiri
	s := config.CreateServer("127.0.0.1", "8000")

	// run server
	err := s.ListenAndServe()

	// handle server execution error
	if err != nil {
		panic(err)
	}
}
