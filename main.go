package main

import (
	"restapi/configs"
	"restapi/controllers"
	"restapi/models"
)

func main() {
	configs.ReadConfig("config.json")
	models.ConnectDB(configs.DBconfig())
	controllers.Route(configs.Configs.HTTP)
}
