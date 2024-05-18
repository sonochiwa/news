package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sonochiwa/news/internal/models"
	"github.com/sonochiwa/news/internal/utils"
)

func (h *Handlers) signUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !utils.IsEmailValid(user.Login) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}

	if len(user.Login) == 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no email provided and len must be greater than 3"})
		return
	}

	if len(user.PasswordHash) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 8 characters"})
		return
	}

	r, err := h.service.Users.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": r})
}

func (h *Handlers) signIn(c *gin.Context) {
	var input models.SignInUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cred, err := h.service.Users.CheckUser(input.Login)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if input.Login != cred.Login || !utils.CheckPasswordHash(input.Password, cred.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	user, _ := h.service.Users.GetUserByEmail(input.Login)

	fmt.Println(user.Login)

	token, err := utils.GenerateJWT(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
