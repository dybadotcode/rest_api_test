package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"restapi/controllers"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

//Configs ...
type Configs struct {
	dialect  string
	host     string
	port     string
	user     string
	dbname   string
	password string
	sslmode  string
	http     string
}

func main() {

	fmt.Printf("Start\n")
	configs, dbConfig := readConfig("config.json")
	// Подключение к базе данных

	models.ConnectDB(configs.dialect, dbConfig)
	// Маршруты

	route := gin.Default()
	route.POST("/rsses", controllers.CreateRss)
	route.GET("/rsses", controllers.GetAllRsses)
	route.GET("/rsses/:id", controllers.GetRss)
	route.DELETE("/rsses/:id", controllers.DeleteRss)
	route.PATCH("/rsses/:id", controllers.UpdateRss)
	route.GET("/search", controllers.SearchRss)
	route.Run(configs.http)
	fmt.Printf("End\n")
}

func readConfig(file string) (Configs, string) {
	var configs Configs
	rawdata, _ := ioutil.ReadFile("config.json")
	var config map[string]interface{}
	err := json.Unmarshal(rawdata, &config)
	if err != nil {
		panic("Cannot unmarshal the json ")
	}
	configs.dialect = fmt.Sprintf("%v", config["dialect"])
	configs.host = fmt.Sprintf("%v", config["host"])
	configs.port = fmt.Sprintf("%v", config["port"])
	configs.user = fmt.Sprintf("%v", config["user"])
	configs.dbname = fmt.Sprintf("%v", config["dbname"])
	configs.password = fmt.Sprintf("%v", config["password"])
	configs.sslmode = fmt.Sprintf("%v", config["sslmode"])
	configs.http = fmt.Sprintf("%v", config["http"])
	dbConfig := "host=" + configs.host +
		" port=" + configs.port +
		" user=" + configs.user +
		" dbname=" + configs.dbname +
		" password=" + configs.password +
		" sslmode=" + configs.sslmode
	return configs, dbConfig
}
