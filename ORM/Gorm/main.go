package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
	Age   int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("DB connected successfully")
	fmt.Print(db)

	db.AutoMigrate(&User{})

	user := User{
		Name:  "Abbas",
		Email: "123@gmail.com",
		Age:   23,
	}
	db.Create(&user)
	var user2 User
	db.First(&user2, 1)
	fmt.Print(user2)
}
