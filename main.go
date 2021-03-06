package main

// @title Natrya Forum API
// @version 1.0
// @description Natrya Forum menggunakan bahasa pemrograman golang

// @host localhost:8888/api
// @BasePath /v1
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

import (
	"natrya.co.id/ryanthe/forum/api"
	_ "natrya.co.id/ryanthe/forum/docs" // docs is generated by Swag CLI, you have to import it.
)

func main() {

	api.Run()

}
