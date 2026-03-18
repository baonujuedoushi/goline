package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func orm() {
	type User struct {
		gorm.Model
		Name  string `gorm:"type:varchar(100);not null"`
		Email string `gorm:"uniqueIndex"`
		Age   int    `gorm:"default:18"`
	}
	db := newDb()
	db.Create(&User{Name: "test", Email: "test@gamil.com", Age: 20})
	var result User
	db.Where("name = ?", "test").First(&result)
	var userList User
	db.Where("name IN ?", []string{"test1", "test2"}).Scan(&userList)
}

func newDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("dns to db"), &gorm.Config{})
	if err != nil {
		panic("database connection failed")
	}
	return db
}
