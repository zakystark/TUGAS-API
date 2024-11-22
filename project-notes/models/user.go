package models

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
    Notes    []Note `json:"notes" gorm:"foreignkey:UserID"`
}
