package controllers

import "github.com/gin-gonic/gin"
 
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func HashPassword(password string) (string) {
	// Implement password hashing logic here
	return ""
}

func VerifyPassword(userPassword, providePassword string) (bool,string) {
	// Implement password verification logic here
	return false,""
}