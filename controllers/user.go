package controllers

import (
	"golang-rest-api-template/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindUsers godoc
// @Summary find users
// @Schemes
// @Description fetch all users data
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} user data
// @Router /users [get]
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// CreateUser godoc
// @Summary create user
// @Schemes
// @Description create user entry with name and email
// @Tags users
// @Accept json
// @Produce json
// @Success 201 {string} user data
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var input models.CreateUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Email: input.Email}

	models.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// FindUser godoc
// @Summary find user
// @Schemes
// @Description find user entry by id
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} user data
// @Router /users/{id} [get]
func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser godoc
// @Summary update user
// @Schemes
// @Description update user entry by id
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} user data
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	var input models.UpdateUser

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(models.User{Name: input.Name, Email: input.Email})

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser godoc
// @Summary delete user
// @Schemes
// @Description delete user entry by id
// @Tags users
// @Accept json
// @Produce json
// @Success 204 {string} empty content
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
