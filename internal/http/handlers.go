package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Credentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (d *Controller) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (d *Controller) Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Invalid request", "user": user})
		return
	}
	u, err := d.ucUser.Register(user.Username, user.Email, user.Password)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User registered successfully",
		"user":    u,
	})
}

func (d *Controller) Login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	err := d.ucUser.Login(creds.Email, creds.Password)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Logged in successfully",
	})
}

func (d *Controller) Logout(c *gin.Context) {
}

func (d *Controller) Welcome(c *gin.Context) {
}

func (d *Controller) Refresh(c *gin.Context) {
}
