package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-ecommerce-v2/database"
	"go-ecommerce-v2/models"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	println(os.Getenv("DB_PORT"))
	db := database.OpenDbConnection()

	//drop(db)
	migrate(db)

	goGinicEngine := gin.Default()

	goGinicEngine.Run(":8080")
}

//func drop(db *gorm.DB) {
//	db.DropTableIfExists(
//		&models.FileUpload{},
//		&models.Comment{},
//		&models.OrderItem{}, &models.Order{}, &models.Address{},
//		&models.ProductCategory{}, &models.ProductTag{},
//		&models.Tag{}, &models.Category{},
//		&models.Product{},
//		&models.UserRole{}, &models.Role{}, &models.User{},
//	)
//}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Address{})

	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Comment{})

	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.OrderItem{})

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.ProductCategory{})

	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.UserRole{})

	db.AutoMigrate(&models.Tag{})
	db.AutoMigrate(&models.ProductTag{})

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.FileUpload{})

}
