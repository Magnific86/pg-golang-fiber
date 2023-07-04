package models

import "gorm.io/gorm"

type Posts struct {
	ID      uint    `gorm:"primary key;autoIncrement" json:"id"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func MigratePosts(db *gorm.DB) error {
	err := db.AutoMigrate(&Posts{})
	return err
}
