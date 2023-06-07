package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/AdityaBkIIITA/gin-gorm-rest/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	USER := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	fmt.Println("USER:", USER)

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", USER, PASSWORD, HOST, PORT)
	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	fmt.Println("Unable to connect database : ", err)
	// 	panic(err)
	// }
	// defer db.Close()

	// _, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")
	// if err != nil {
	// 	fmt.Println("Unable to initialise database : ", err)
	// 	panic(err)
	// }

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASSWORD, HOST, PORT, DBNAME)
	fmt.Println(dsn)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Unable to connect database with DBNAME : ", err)
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
