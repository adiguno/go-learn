package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Story table
type Story struct {
	gorm.Model
	id       int
	title    string
	authorID int
}

// Product table
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	fmt.Println("asdf")

	dsn := "user=postgres password=postgres DB.name=test-go-db port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Story{})
	var story Story
	// db.First(&story, 1) // find product with integer primary key
	// fmt.Println(story.title)

	db.First(&story, "title = ?", "Cool story")
	fmt.Println(story.title)

	//todo migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Println(product.Price)

}
