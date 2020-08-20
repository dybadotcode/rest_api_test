package controllers

import (
	"go/dev/golang/src/restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateArticleInput struct {
	Writer  string `json:"writer" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateArticleInput struct {
	Writer  string `json:"writer"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// GET /articles
// Получаем список всех
func GetAllArticles(context *gin.Context) {
	var article []models.Article
	models.DB.Find(&article)

	context.JSON(http.StatusOK, gin.H{"article": article})
}

// GET /articles/:id
// Получение одной новости по ID
func GetArticle(context *gin.Context) {
	// Проверяем имеется ли запись
	var article models.Article
	if err := models.DB.Where("id = ?", context.Param("id")).First(&article).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"articles": article})
}

// POST /articles
// Создание новости
func CreateArticle(context *gin.Context) {
	var input CreateArticleInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{Writer: input.Writer, Title: input.Title, Content: input.Content}
	models.DB.Create(&article)

	context.JSON(http.StatusOK, gin.H{"articles": article})
}

// PATCH /articles/:id
// Изменения информации
func UpdateArticle(context *gin.Context) {
	// Проверяем имеется ли такая запись перед тем как её менять
	var article models.Article
	if err := models.DB.Where("id = ?", context.Param("id")).First(&article).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateArticleInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&article).Update(input)

	context.JSON(http.StatusOK, gin.H{"articles": article})
}

// DELETE /articles/:id
// Удаление
func DeleteArticle(context *gin.Context) {
	// Проверяем имеется ли такая запись перед тем как её удалять
	var article models.Article
	if err := models.DB.Where("id = ?", context.Param("id")).First(&article).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	models.DB.Delete(&article)

	context.JSON(http.StatusOK, gin.H{"articles": true})
}
