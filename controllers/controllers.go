package controllers

import (
	"net/http"

	"restapi/models"

	"github.com/gin-gonic/gin"
)

//ContentIDPair ...
type ContentIDPair struct {
	ID1 string `json:"id1"`
	ID2 string `json:"id2"`
}

// Route ...
func Route(http string) {
	route := gin.Default()
	route.GET("/rsses", GetAllRsses)
	route.GET("/rsses/:id", GetRss)
	route.GET("/content/:id", GetContent)
	route.GET("/contentCompare", ContentCompare)
	route.GET("/search", SearchRss)
	route.Run(http)
}

// GetAllRsses ... запрос
func GetAllRsses(context *gin.Context) {
	rsses := models.GetALLRSSesWhithoutContent()
	context.JSON(http.StatusOK, gin.H{"rsses": *rsses})
}

// GetRss ...
// GET /rsses/:id
func GetRss(context *gin.Context) {
	rss, err := models.GetContentByID(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"rss": rss})
}

// GetContent ...
// GET /contents/:id
func GetContent(context *gin.Context) {
	content, err := models.GetContentByID(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"content": content})
}

// ContentCompare ...
// GET /contents/compare
func ContentCompare(context *gin.Context) {
	var contentIDPair ContentIDPair
	if err := context.ShouldBindJSON(&contentIDPair); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c1, err := models.GetContentByID(contentIDPair.ID1)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Контент по id1 не существует"})
		return
	}
	c2, err := models.GetContentByID(contentIDPair.ID2)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Контент по id2 не существует"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"content": models.RssFeedCompareString(c1.Content, c2.Content)})
}

// SearchRss ...
// GET /rsses/search/
func SearchRss(context *gin.Context) {
	var rssQuery *models.RssQuery
	if err := context.ShouldBindJSON(rssQuery); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rsses, err := models.Search(rssQuery)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Записи не существуют"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"rsses": rsses})

}
