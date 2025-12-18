package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Subscription struct {
	Sub string
}
type User struct {
	gorm.Model
	Name         string
	Age          int
	MemberNumber *string
	Subscription Subscription `gorm:"type:json;serializer:json"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Print("before create")
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Print("after create")
	return
}

type User2 struct {
	gorm.Model
	Name         string
	Age          int
	MemberNumber *string
	Subscription Subscription `gorm:"embedded"`
}

func main() {
	dsn := "gormuser:password123@tcp(127.0.0.1:3306)/gorm_example?charset=utf8mb4&parseTime=True&loc=Local"

	// Step 3: Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Connected to database successfully!")

	// Step 4: Auto-migrate the User model (creates table if not exists)
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate database!")
	}
	db.AutoMigrate(&User2{})

	fmt.Println("Table migrated successfully!")

	// user := User{
	// 	Name:         "Abbas",
	// 	Age:          23,
	// 	MemberNumber: nil,
	// 	Subscription: Subscription{
	// 		Sub: "Normal",
	// 	},
	// }

	// result := db.Create(&user)
	// if result.Error != nil {
	// 	fmt.Println("Error:", result.Error)
	// } else {
	// 	fmt.Println("Inserted user ID:", user.ID)
	// }

	// var temp User
	// db.First(&temp, 1)
	// fmt.Printf("%+v\n", temp)
	// fmt.Print(temp.ID)

	// user2 := User{
	// 	Name:         "feeehu",
	// 	Age:          22,
	// 	MemberNumber: nil,
	// 	Subscription: Subscription{
	// 		Sub: "normal",
	// 	},
	// }
	// db.Create(&user2)

	var users []User
	db.Find(&users)

	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Subscription: %v\n", u.ID, u.Name, u.Age, u.Subscription)
	}

	// var temp User2
	// db.First(&temp, 3)
	// fmt.Printf("%v\n", temp)
	var tempp []User
	db.Where("id=?", 4).Find(&tempp)
	fmt.Printf("%v\n", tempp)

}
