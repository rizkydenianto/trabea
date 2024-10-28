package main

import (
	"trabea/app/login"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/login/roles", login.GetRoles)

	router.Run("localhost:8080")
}
