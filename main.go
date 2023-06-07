package main

import (
	"fmt"
	"os"

	"github.com/AdityaBkIIITA/gin-gorm-rest/config"
	"github.com/AdityaBkIIITA/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	config.LoadEnv()   //loading env
	config.ConnectDB() //new database connection

	if v, f := os.LookupEnv("USE_AUTH"); f && v == "true" {
		fmt.Println("Using Auth")
		routes.AuthRouter(router)
	} else {
		routes.UserRoute(router)
	}

	router.Run(":8080")
}
