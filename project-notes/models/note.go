package models

import "github.com/jinzhu/gorm"

type Note struct {
    gorm.Model
    Title  string `json:"title"`
    Body   string `json:"body"`
    UserID uint   `json:"user_id"`
}
