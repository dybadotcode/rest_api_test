package models

import (
	// ORM framework

	"strings"

	"github.com/jinzhu/gorm"

	//HTTP framework
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB ..
var DB *gorm.DB

//Result ...
type Result struct {
	ID           int64
	URL          string
	UserAgent    string
	AccessTime   string
	ResponseTime string
	Content      string
}

//ResultWhithoutContent ...
type ResultWhithoutContent struct {
	ID           int64
	URL          string
	UserAgent    string
	AccessTime   string
	ResponseTime string
	ContentID    int64
}

//Content ...
type Content struct {
	ID      int64
	Content string
}

//RssQuery ...
type RssQuery struct {
	AccessTimeFrom   string `json:"accessTimeFrom"`
	AccessTimeTo     string `json:"accessTimeTo"`
	ResponseTimeFrom string `json:"responseTimeFrom"`
	ResponseTimeTo   string `json:"responseTimeTo"`
	UserAgent        string `json:"userAgent"`
	URL              string `json:"url"`
}

//ConnectDB ...
func ConnectDB(rootConfigs string) {
	configs := strings.Split(rootConfigs, "/")
	dialect := configs[0]
	params := configs[1]
	db, err := gorm.Open(dialect, params)
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	DB = db
	//DB.LogMode(true)
}

//AccessesJoinByID ...
func AccessesJoinByID() *gorm.DB {
	db := DB.Table("accesses").Select("accesses.id, urls.url, bots.name as user_agent, accesses.access_time, accesses.response_time, accesses.content_id").
		Joins("left join contents on contents.id = accesses.content_id").
		Joins("left join bots on bots.id = accesses.bot_id").
		Joins("left join urls on urls.id = contents.url_id")
	return db
}

//GetALLRSSesWhithoutContent ...
func GetALLRSSesWhithoutContent() *[]ResultWhithoutContent {
	result := new([]ResultWhithoutContent)
	AccessesJoinByID().
		Order("access_time desc").
		Scan(result)
	return result
}

//GetContentByID ...
func GetContentByID(id string) (*Content, error) {
	result := new(Content)
	err := DB.Table("contents").Select("contents.id, contents.content").
		Where("contents.id = ?", id).
		Scan(result).
		Error
	return result, err
}

//Search ...
func Search(rssQuery *RssQuery) (*[]ResultWhithoutContent, error) {
	result := new([]ResultWhithoutContent)
	// Where ... AND ...
	// Adds condition if rssQuery.Parametr != ""
	db := AccessesJoinByID()
	addQueryPart(&db, "access_time >= ?", rssQuery.AccessTimeFrom)
	addQueryPart(&db, "access_time <= ?", rssQuery.AccessTimeTo)
	addQueryPart(&db, "response_time >= ?", rssQuery.ResponseTimeFrom)
	addQueryPart(&db, "response_time <= ?", rssQuery.ResponseTimeTo)
	addQueryPart(&db, "user_agent = ?", rssQuery.UserAgent)
	addQueryPart(&db, "url = ?", rssQuery.URL)
	//Find
	err := db.Find(result).Error
	return result, err
}

// addQueryPart ...
func addQueryPart(db **gorm.DB, queryPart string, data string) {
	if data != "" {
		*db = (*db).Where(queryPart, data)
	}
}
