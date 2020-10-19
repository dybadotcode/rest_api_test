package controllers

import (
	"net/http"
	"time"

	"fmt"
	"restapi/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//RssInput ...
type RssInput struct {
	AccessTime   string `json:"accessTime" binding:"required"`
	ResponseTime string `json:"responseTime" binding:"required"`
	UserAgent    string `json:"userAgent" binding:"required"`
	URL          string `json:"url" binding:"required"`
	Content      string `json:"content" binding:"required"`
}

//RssUpdate ...
type RssUpdate struct {
	AccessTime   string `json:"accessTime"`
	ResponseTime string `json:"responseTime"`
	UserAgent    string `json:"userAgent"`
	URL          string `json:"url"`
	Content      string `json:"content"`
}

//RssQuery ...
type RssQuery struct {
	AccessTimeFrom   string `json:"accessTimeFrom"`
	AccessTimeTo     string `json:"accessTimeTo"`
	ResponseTimeFrom string `json:"responseTimeFrom"`
	ResponseTimeTo   string `json:"responseTimeTo"`
	UserAgent        string `json:"userAgent"`
	URL              string `json:"url"`
	Content          string `json:"content"`
}

// CreateRss ...
//POST /rsses
func CreateRss(context *gin.Context) {
	var input RssInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessTime, _ := time.Parse(time.RFC3339, input.AccessTime)
	responseTime, _ := time.Parse(time.RFC3339, input.ResponseTime)
	rss := models.Rss{AccessTime: accessTime, ResponseTime: responseTime, UserAgent: input.UserAgent, URL: input.URL, Content: input.Content}
	models.DB.Create(&rss)
	context.JSON(http.StatusOK, gin.H{"rsses": rss})
}

// GetAllRsses ...
// GET /rsses
func GetAllRsses(context *gin.Context) {
	var rss []models.Rss
	models.DB.Find(&rss)
	context.JSON(http.StatusOK, gin.H{"rsses": rss})
}

// GetRss ...
// GET /rsses/:id
func GetRss(context *gin.Context) {
	// Проверяем имеется ли запись
	var rss models.Rss
	if err := models.DB.Where("id = ?", context.Param("id")).First(&rss).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"rsses": rss})
}

// SearchRss ...
// GET /rsses/search/
func SearchRss(context *gin.Context) {
	var rssQuery RssQuery
	var rsses []models.Rss
	if err := context.ShouldBindJSON(&rssQuery); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Where ... AND ...
	// Adds condition if rssQuery.Parametr != ""
	addQueryPart(&models.DB, "access_time >= ?", rssQuery.AccessTimeFrom)
	addQueryPart(&models.DB, "access_time <= ?", rssQuery.AccessTimeTo)
	addQueryPart(&models.DB, "response_time >= ?", rssQuery.ResponseTimeFrom)
	addQueryPart(&models.DB, "response_time <= ?", rssQuery.ResponseTimeTo)
	addQueryPart(&models.DB, "user_agent = ?", rssQuery.UserAgent)
	addQueryPart(&models.DB, "url = ?", rssQuery.URL)
	addQueryPart(&models.DB, "content = ?", rssQuery.Content)
	//Find
	if err := models.DB.Find(&rsses).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Записи не существуют"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"rsses": rsses})

}

// UpdateRss ...
// PATCH /rsses/:id
func UpdateRss(context *gin.Context) {
	// Проверяем имеется ли такая запись перед тем как её менять
	var rss models.Rss
	if err := models.DB.Where("id = ?", context.Param("id")).First(&rss).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input RssUpdate
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&rss).Update(input)

	context.JSON(http.StatusOK, gin.H{"rsses": rss})
}

// DeleteRss ...
// DELETE /rsses/:id
func DeleteRss(context *gin.Context) {
	// Проверяем имеется ли такая запись перед тем как её удалять
	var rss models.Rss
	if err := models.DB.Where("id = ?", context.Param("id")).First(&rss).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}
	models.DB.Delete(&rss)

	context.JSON(http.StatusOK, gin.H{"rss": true})
}

// addQueryPart ...
func addQueryPart(db **gorm.DB, queryPart string, data string) {
	if data != "" {
		fmt.Println(queryPart + " " + data)
		*db = (*db).Where(queryPart, data)
	}
}
