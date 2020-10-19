package main

import (
	"fmt"
	"restapi/controllers"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Printf("Start\n")
	// Подключение к базе данных
	models.ConnectDB()
	// Маршруты
	route := gin.Default()
	route.POST("/rsses", controllers.CreateRss)
	route.GET("/rsses", controllers.GetAllRsses)
	route.GET("/rsses/:id", controllers.GetRss)
	route.DELETE("/rsses/:id", controllers.DeleteRss)
	route.PATCH("/rsses/:id", controllers.UpdateRss)
	route.GET("/search", controllers.SearchRss)
	route.Run(":8083")
	fmt.Printf("End\n")
}
