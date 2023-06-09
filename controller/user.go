package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AdityaBkIIITA/gin-gorm-rest/auth"
	"github.com/AdityaBkIIITA/gin-gorm-rest/config"
	"github.com/AdityaBkIIITA/gin-gorm-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Tokener(c *gin.Context) {

	fmt.Println("Tokener Working!!!")

	pswd := uuid.New().String()
	user := time.Now().String()

	token, err := auth.GenerateAccessToken(user, pswd)

	if err != nil {
		c.Writer.WriteHeader(500)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetAllUsers(c *gin.Context) {
	fmt.Println("Rest API!")
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, &users)
}

func GetUser(c *gin.Context) {
	user := models.User{}
	config.DB.Where("id=?", c.Param("id")).Find(&user)
	c.JSON(http.StatusOK, &user)
}

func CreateUser(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(http.StatusOK, &user)

}

func UpdateUser(c *gin.Context) {

	// Find the existing user by ID
	user := models.User{}
	result := config.DB.First(&user, c.Param("id"))
	if result.Error != nil {
		// Handle the case when the user is not found
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update the user fields with the new values
	var updatedUser models.User
	c.ShouldBindJSON(&updatedUser)

	// Update the user fields
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	// Update other fields as needed

	// Save the updated user
	result = config.DB.Save(&user)
	if result.Error != nil {
		// Handle the error while saving the user
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Return the updated user
	c.JSON(http.StatusOK, &user)
}

func DeleteUser(c *gin.Context) {
	user := models.User{}
	config.DB.Where("id= ?", c.Param("id")).Delete(&user)
	c.JSON(http.StatusOK, &user)
}

func DeleteAllUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Delete(&users)
	c.JSON(http.StatusOK, &users)
}
