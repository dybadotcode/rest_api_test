# RSS_REST


 REST_API для логирования rss

# RSS structure

type Rss struct {
	ID           int64     `gorm:"type:bigint; primaryKey"`
	AccessTime   time.Time `gorm:"type:timestamp with time zone"`
	ResponseTime time.Time `gorm:"type:timestamp with time zone"`
	UserAgent    string    `gorm:"type:varchar"`
	URL          string    `gorm:"type:text"`
	Content      string    `gorm:"type:text"`
}

# Query templates

# POST "/rsses"


{
    "accessTime":   "2009-01-02T15:04:05+03:00",
	"responseTime": "2011-01-02T13:24:05+03:00",
	"userAgent": "Витя",
	"url": "https://developer.mozilla3.org",
	"content": "Делимся рецептами отличных домашних настоек"
}



# GET "/rsses"
# GET "/rsses/:id"


- host/rsses/4


# DELETE "/rsses/:id"


- host/rsses/4


# PATCH "/rsses/:id"


- host/rsses/4


# GET "/search"


{
	"accessTimeFrom": "2006-01-02T15:04:05+03:00",
	"accessTimeTo": "2006-01-02T15:04:05+03:00",
	"responseTimeFrom": "2006-01-02T15:04:05+03:00",
	"responseTimeTo": "2006-01-02T15:04:05+03:00",
	"userAgent": "Витя",
	"url": "https://developer.mozilla3.org",
	"content": "Делимся рецептами отличных домашних настоек"
}