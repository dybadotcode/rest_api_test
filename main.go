package main

import (
	"go/dev/golang/src/restapi/controllers"
	"go/dev/golang/src/restapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// Подключение к базе данных 666
	models.ConnectDB()

	// Маршруты
	route.GET("/articles", controllers.GetAllArticles)
	route.POST("/articles", controllers.CreateArticle)
	route.GET("/articles/:id", controllers.GetArticle)
	route.PATCH("/articles/:id", controllers.UpdateArticle)
	route.DELETE("/articles/:id", controllers.DeleteArticle)

	// Запуск сервера
	route.Run()
}
