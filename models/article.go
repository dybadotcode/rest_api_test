package models

type Article struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Writer  string `json:"writer"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
