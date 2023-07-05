package models

import (
	"klevt/go/pkg/mod/gorm.io/gorm@v1.25.2"
	"syscall/js"

	"gorm.io/gorm"
)

type Blob struct {
	object *js.Value
}

type FormFile struct {
	Blob
	cur        int64
	buffersize int64
	size       int64
}

type Post struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	File    FormFile `json:"file"`
}

type Posts struct {
	ID      uint      `gorm:"primary key;autoIncrement" json:"id"`
	Title   *string   `json:"title"`
	Content *string   `json:"content"`
	File    *FormFile `json:"file"`
}

func MigratePosts(db *gorm.DB) error {
	err := db.AutoMigrate(&Posts{})
	return err
}
