package handlers

import (
	"Server/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")

	var storedPassword string
	err := db.DB.QueryRow("SELECT userpassword FROM auth WHERE userid = $1", id).Scan(&storedPassword)
	isValid := CheckPasswordHash(password, storedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Invalid username or password"})
		return
	}
	if isValid {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Invalid username or password"})
		return
	}

}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil // 비밀번호가 일치하면 nil 반환
}
