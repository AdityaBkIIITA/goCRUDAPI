package config

import (
	"database/sql"
	"fmt"

	"github.com/AdityaBkIIITA/gin-gorm-rest/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:aditya1304@tcp(localhost:3306)/userdb?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),
	// )
	fmt.Println(dsn)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to initialize GORM")
	}

	DB.AutoMigrate(&models.User{})
}
